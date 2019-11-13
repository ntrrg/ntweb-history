// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	hc "github.com/gohugoio/hugo/config"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	ntos "nt.web.ve/go/ntgo/os"
)

var (
	Default = Build

	hugoVersion = "0.59.1"
	hugoPort    = "1313"
	hugoConfig  = "config.yaml"

	cfg hc.Provider
)

func Build() error {
	return sh.RunV("hugo")
}

type BumpVersion mg.Namespace

func (BumpVersion) Hugo() error {
	fn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == ".git" || path == "content" {
			return filepath.SkipDir
		}

		if info.IsDir() || strings.HasSuffix(path, ".swp") || path == "go.sum" {
			return nil
		}

		fd, err := os.Open(path)
		if err != nil {
			return err
		}

		s := bufio.NewScanner(fd)

		for s.Scan() {
			if bytes.Contains(s.Bytes(), []byte(hugoVersion)) {
				fmt.Println(path)
			}
		}

		if err := s.Err(); err != nil {
			return err
		}

		return nil
	}

	return filepath.Walk(".", fn)
}

func Clean() error {
	for _, dst := range []string{prjsRepos, goRepos, "public", "resources"} {
		if err := sh.Rm(dst); err != nil {
			return err
		}
	}

	return nil
}

type Lint mg.Namespace

func (Lint) Go() error {
	return sh.RunV("gofmt", "-d", "-e", "-s", "magefile.go")
}

func Run() error {
	return sh.RunV(
		"hugo", "server", "-D", "-E", "-F", "--noHTTPCache", "--i18n-warnings",
		"--bind", "0.0.0.0", "--port", hugoPort,
		"--baseUrl", "/", "--appendPort=false",
	)
}

// Helpers

func cloneRepo(dst, src string) error {
	if dst == "" {
		dst = path.Base(src)
	}

	if _, err := os.Stat(dst); err == nil {
		return sh.Run("git", "-C", dst, "pull", "origin", "master")
	}

	if err := sh.Run("git", "clone", src, dst); err != nil {
		if err := sh.Rm(dst); err != nil {
			return err
		}

		return err
	}

	return nil
}

func getHugoConfig(cfgFile string) (hc.Provider, error) {
	f, err := os.Open(cfgFile)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	cfgData, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return hc.FromConfigString(string(cfgData), filepath.Ext(cfgFile))
}

func init() {
	var err error

	cfg, err = getHugoConfig(hugoConfig)
	if err != nil {
		panic(err)
	}
}

////////////
// Docker //
////////////

type Docker mg.Namespace

func (Docker) Build() error {
	return runHugoDocker()
}

func (Docker) Run() error {
	return runHugoDocker(
		"server", "-D", "-E", "-F", "--noHTTPCache", "--i18n-warnings",
		"--bind", "0.0.0.0", "--port", hugoPort,
		"--baseUrl", "/", "--appendPort=false",
	)
}

func runHugoDocker(args ...string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	args = append([]string{
		"run", "--rm", "-i", "-t",
		"-u", u.Uid,
		"-v", wd + ":/site/",
		"-p", hugoPort + ":" + hugoPort,
		"ntrrg/hugo:" + hugoVersion,
	}, args...)

	return sh.RunV("docker", args...)
}

////////////////////////
// Content generation //
////////////////////////

type Gen mg.Namespace

func (Gen) Default() {
	mg.Deps(Gen.GoPkgs, Gen.Projects)
}

// Go packages

var (
	goDir   = filepath.Clean("content/go")
	goRepos = filepath.Join(os.TempDir(), "ntweb-go-pkgs")
)

func (Gen) GoPkgs() error {
	if err := os.MkdirAll(goRepos, 0755); err != nil {
		return err
	}

	srcRE, err := regexp.Compile(`(?:([\w/\-]+) )?(https?:/.+)`)
	if err != nil {
		return err
	}

	pkgRE, err := regexp.Compile(`([\w/\-.]+): (.+)?`)
	if err != nil {
		return err
	}

	for _, repo := range cfg.GetStringSlice("params.goPackages") {
		repoElems := srcRE.FindStringSubmatch(repo)
		name, src := repoElems[1], repoElems[2]

		if name == "" {
			name = path.Base(src)
		}

		dst := filepath.Join(goRepos, name)

		if err := cloneRepo(dst, src); err != nil {
			return err
		}

		c := exec.Command("go", "list", "-m")
		c.Dir = dst
		output, err := c.Output()
		if err != nil {
			return err
		}

		mod := string(bytes.TrimSpace(output))
		if err := writePackagePage(src, mod, mod, ""); err != nil {
			return err
		}

		c = exec.Command(
			"go", "list", "-f", "{{ .ImportComment }}: {{ .Doc }}", "./"+name+"/...",
		)

		c.Dir = goRepos
		output, err = c.Output()
		if err != nil {
			return err
		}

		for _, entry := range bytes.Split(bytes.TrimSpace(output), []byte("\n")) {
			elms := pkgRE.FindSubmatch(entry)

			if len(elms) < 3 {
				return fmt.Errorf("bad package format in module %s: %q", mod, entry)
			}

			pkg, doc := string(elms[1]), string(elms[2])

			if err := writePackagePage(src, mod, pkg, doc); err != nil {
				return err
			}
		}
	}

	return nil
}

func genPackagePage(src, mod, pkg, doc string) string {
	prefix := path.Dir(mod) + "/"
	pkg = strings.TrimPrefix(string(pkg), prefix)

	content := `
---
title: %s
source-code: %s
module: %s
description: %s
url: /go/%s
---
`[1:]

	return fmt.Sprintf(content, pkg, src, mod, doc, pkg)
}

func writeMultiLangFile(path string, content []byte, mode os.FileMode) error {
	i := strings.LastIndex(path, ".")
	if i < 0 {
		i = len(path) - 1
	}

	ext := filepath.Ext(path)

	for lang := range cfg.GetStringMap("languages") {
		path := path[:i] + "." + lang + ext

		err := ioutil.WriteFile(path, content, mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func writePackagePage(src, mod, pkg, doc string) error {
	return writeMultiLangFile(
		filepath.Join(goDir, strings.ReplaceAll(pkg, "/", "-")+".md"),
		[]byte(genPackagePage(src, mod, pkg, doc)), 0644,
	)
}

// Projects

var (
	prjsDir   = filepath.Clean("content/projects")
	prjsRepos = filepath.Join(os.TempDir(), "ntweb-projects")
)

func (Gen) Projects() error {
	if err := os.MkdirAll(prjsRepos, 0755); err != nil {
		return err
	}

	for _, repoUrl := range cfg.GetStringSlice("params.projects") {
		name := path.Base(repoUrl)
		prjRepo := filepath.Join(prjsRepos, name)
		if err := cloneRepo(prjRepo, repoUrl); err != nil {
			return err
		}

		src := filepath.Join(prjRepo, ".ntweb")
		dst := filepath.Join(prjsDir, name)

		if err := sh.Rm(dst); err != nil {
			return err
		}

		if err := ntos.Cp(dst, src); err != nil {
			return err
		}

		output, err := sh.Output("git", "-C", prjRepo, "log", "--format=%cI")
		if err != nil {
			return err
		}

		dates := strings.Split(output, "\n")
		publishedAt := dates[len(dates)-1]
		modifiedAt := dates[0]

		license, err := getLincense(filepath.Join(prjRepo, "LICENSE"))
		if err != nil {
			return err
		}

		metadata := []byte(`---
publishdate: ` + publishedAt + `
date: ` + modifiedAt + `
metadata:
  source-code: ` + repoUrl + `
  license: ` + license + `
`)

		indexes, err := filepath.Glob(dst + "/index.*.md")
		if err != nil {
			return err
		}

		for _, index := range indexes {
			var (
				f       *os.File
				content []byte
				err     error
			)

			f, err = os.Open(index)
			if err != nil {
				goto finish
			}

			content, err = ioutil.ReadAll(f)
			if err != nil {
				goto finish
			}

			f.Close()
			f, err = os.Create(index)
			if err != nil {
				goto finish
			}

			content = append(
				metadata,
				content[bytes.IndexByte(content, '\n')+1:]...,
			)
			_, err = f.Write(content)

		finish:
			f.Close()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getLincense(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		switch {
		case bytes.Contains(s.Bytes(), []byte("The MIT License")):
			return "MIT", nil
		case bytes.Contains(s.Bytes(), []byte("DO WHAT THE FUCK YOU WANT TO")):
			return "WTFPL", nil
		}
	}

	return "", s.Err()
}
