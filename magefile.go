// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
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
)

func init() {
	var err error
	cfg, err = getHugoConfig(hugoConfig)
	if err != nil {
		panic(err)
	}

	wd, err = os.Getwd()
	if err != nil {
		return err
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

	for _, target := range []string{"go-pkgs", "public", "resources"} {
		if err := os.RemoveAll(target); err != nil {
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

func (Gen) Go() error {
	goDir := filepath.Clean("content/go")

	err := writeMultiLangFile(
		filepath.Join(goDir, "_index.md"),
		[]byte(`
---
title: Go packages
hidden: true
---
`[1:]),
		0644,
	)

	if err != nil {
		return err
	}

	re, err := regexp.Compile(`(?:([\w/\-]+) )?(https?:/.+)`)
	if err != nil {
		return err
	}

	for _, src := range cfg.GetStringSlice("params.goPackages") {
		_ = re.FindStringSubmatch(src)
		pkgs, err := getGoPkgs(src)
		if err != nil {
			return nil
		}

		_ = pkgs

		// for _, pkg := range pkgs {
		// 	err := writeMultiLangFile(
		// 		filepath.Join(goDir, strings.ReplaceAll(pkg, "/", "-")+".md"),
		// 		[]byte(pkg), 0644,
		// 	)

		// 	if err != nil {
		// 		return err
		// 	}
		// }
	}

	return nil
}

func genPackagePage(name, src, desc string) string {
	content := `
---
title: %s
source-code: %s
descripton: %s
url: /go/%s
---
`[1:]

	return fmt.Sprintf(content, name, src, desc, name)
}

func getGoPkgs(src string) ([]string, error) {
	goPkgsDir := "go-pkgs"
	if err := os.MkdirAll(goPkgsDir, 0755); err != nil {
		return nil, err
	}

	_ = sh.RunV("git", "-C", goPkgsDir, "clone", src)

	err := sh.RunV(
		"go", "list", "-f", "{{ .ImportComment }}: {{ .Doc }}",
		filepath.Join("./", goPkgsDir),
	)

	if err != nil {
		return nil, err
	}

	// genPackagePage(prefix, src)

	return []string{}, nil
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
		return nil
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
	err = sh.RunV(
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
