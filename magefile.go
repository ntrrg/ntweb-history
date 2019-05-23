// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
)

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

	for _, target := range []string{"public", "resources"} {
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
	cfg, err := getHugoConfig(hugoConfig)
	if err != nil {
		return err
	}

	goDir := filepath.Clean("content/go")

	indexFileContent := []byte(`
---
title: Go packages
hidden: true
---
`[1:])

	for lang := range cfg.GetStringMap("languages") {
		err = ioutil.WriteFile(
			filepath.Join(goDir, "_index."+lang+".md"),
			indexFileContent,
			0644,
		)

		if err != nil {
			return err
		}
	}

	re, err := regexp.Compile(`(?:(\w+) )?(https?:/.+)`)
	if err != nil {
		return err
	}

	for _, pkg := range cfg.GetStringSlice("params.goPackages") {
		pkgEls := re.FindStringSubmatch(pkg)
		name, src := pkgEls[1], pkgEls[2]

		if name == "" {
			name = path.Base(src)
		}

		err = ioutil.WriteFile(
			filepath.Join(goDir, name+".md"),
			[]byte(genPackagePage(name, src)),
			0644,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func genPackagePage(name, src string) string {
	content := `
---
title: %s
source-code: %s
url: /go/%s
---
`[1:]

	return fmt.Sprintf(content, name, src, name)
}

func getGoPkgs(pkg string) []string {
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

	return hc.FromConfigString(string(cfgData), filepath.Ext(hugoConfig))
}

func runHugoDocker(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

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
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

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
