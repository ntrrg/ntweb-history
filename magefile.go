// +build mage

package main

import (
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
)

var (
	Default = Build.Default

	hugoVersion = "0.55.6"
	hugoPort    = "1313"
	hugoConfig  = "config.yaml"
	dockerImage = "ntrrg/ntweb"

	lintContainer = strings.Replace(dockerImage, "/", "-", -1) + "-lint"

	cfg hc.Provider
	wd  string

	goDir     = filepath.Clean("content/go")
	goPkgsDir = filepath.Join(os.TempDir(), "ntweb-go-pkgs")
)

func init() {
	var err error

	cfg, err = getHugoConfig(hugoConfig)
	if err != nil {
		panic(err)
	}

	wd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

type Build mg.Namespace

func (Build) Default() error {
	return sh.RunV("hugo")
}

func (Build) Docker() error {
	return runHugoDocker()
}

func Clean() error {
	_ = sh.RunV("docker", "rm", "-f", lintContainer)
	_ = sh.RunV("docker", "rm", "-f", lintContainer+"-watch")

	for _, dst := range []string{goPkgsDir, "public", "resources"} {
		if err := sh.Rm(dst); err != nil {
			return err
		}
	}

	return nil
}

type Lint mg.Namespace

func (Lint) Default() error {
	return runLinter(lintContainer, "latest")
}

func (Lint) Watch() error {
	return runLinter(lintContainer+"-watch", "watch")
}

type Run mg.Namespace

func (Run) Default() error {
	return sh.RunV(
		"hugo", "server", "-D", "-E", "-F",
		"--noHTTPCache", "--port", hugoPort,
	)
}

func (Run) Docker() error {
	return runHugoDocker(
		"server", "-D", "-E", "-F", "--noHTTPCache",
		"--bind", "0.0.0.0", "--port", hugoPort,
		"--baseUrl", "/", "--appendPort=false",
	)
}

type Gen mg.Namespace

func (Gen) GoPkgs() error {
	if err := os.MkdirAll(goPkgsDir, 0755); err != nil {
		return err
	}

	srcRE, err := regexp.Compile(`(?:([\w/\-]+) )?(https?:/.+)`)
	if err != nil {
		return err
	}

	pkgRE, err := regexp.Compile(`(?:([\w/\-.]+): )(.+)`)
	if err != nil {
		return err
	}

	for _, repo := range cfg.GetStringSlice("params.goPackages") {
		repoElems := srcRE.FindStringSubmatch(repo)
		name, src := repoElems[1], repoElems[2]

		if name == "" {
			name = path.Base(src)
		}

		dst := filepath.Join(goPkgsDir, name)

		if _, err := os.Stat(dst); err != nil {
			gerr := sh.Run("git", "clone", "--depth", "1", src, dst)
			if gerr != nil {
				if gerr := sh.Rm(dst); gerr != nil {
					return gerr
				}

				return gerr
			}
		}

		c := exec.Command("go", "list", "-m")
		c.Dir = dst
		output, err := c.Output()
		if err != nil {
			return err
		}

		mod := string(bytes.TrimSpace(output))
		prefix := path.Dir(mod) + "/"

		c = exec.Command(
			"go", "list", "-f", "{{ .ImportComment }}: {{ .Doc }}", "./"+name+"/...",
		)

		c.Dir = goPkgsDir
		output, err = c.Output()
		if err != nil {
			return err
		}

		for _, entry := range bytes.Split(bytes.TrimSpace(output), []byte("\n")) {
			els := pkgRE.FindSubmatch(entry)

			if len(els) < 3 {
				return fmt.Errorf("Bad package format in module %s: %q", mod,  entry)
			}

			pkg, doc := strings.TrimPrefix(string(els[1]), prefix), string(els[2])

			err := writeMultiLangFile(
				filepath.Join(goDir, strings.ReplaceAll(pkg, "/", "-")+".md"),
				[]byte(genPackagePage(pkg, mod, src, doc)), 0644,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func genPackagePage(pkg, mod, src, doc string) string {
	content := `
---
title: %s
module: %s
source-code: %s
description: %s
url: /go/%s
---
`[1:]

	return fmt.Sprintf(content, pkg, mod, src, doc, pkg)
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

	return hc.FromConfigString(string(cfgData), filepath.Ext(hugoConfig))
}

func runHugoDocker(args ...string) error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	args = append([]string{
		"run", "--rm", "-i", "-t",
		"-u", u.Uid,
		"-v", wd+":/srv/",
		"-p", hugoPort+":"+hugoPort,
		"ntrrg/hugo:"+hugoVersion,
	}, args...)

	return sh.RunV("docker", args...)
}

func runLinter(name, tag string) error {
	err := sh.RunV(
		"docker", "run", "--name", name, "-i", "-t",
		"-v", wd+":/files/",
		"ntrrg/md-linter:"+tag,
	)

	if err != nil {
		return sh.RunV("docker", "start", "-a", "-i", name)
	}

	return nil
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
