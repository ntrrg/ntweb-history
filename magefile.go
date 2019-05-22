// +build mage

package main

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	hc "github.com/gohugoio/hugo/config"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default = Build.Default

	hugoVersion = "latest"
	hugoPort    = "1313"
	hugoConfig  = "config.yaml"
	dockerImage = "ntrrg/ntweb"

	lintContainer = strings.Replace(dockerImage, "/", "-", -1) + "-lint"
)

type Build mg.Namespace

func (Build) Default() error {
	return runHugo()
}

func (Build) Docker() error {
	return sh.RunV("docker", "build", "-t", dockerImage, ".")
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

func Run() error {
	return runHugo(
		"server", "-D", "-E", "-F", "--noHTTPCache", "--bind", "0.0.0.0",
		"--port", hugoPort, "--baseUrl", "/", "--appendPort=false",
	)
}

type Gen mg.Namespace

func (Gen) Go() error {
	cfg, err := getHugoConfig(hugoConfig)
	if err != nil {
		return err
	}

	goDir := filepath.Clean("content/go")

	indexFileContent := `
---
title: Go packages
hidden: true
---
`[1:]

	for lang := range cfg.GetStringMap("languages") {
		err = ioutil.WriteFile(
			filepath.Join(goDir, "_index."+lang+".md"),
			[]byte(indexFileContent),
			0644,
		)

		if err != nil {
			return err
		}
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

	return hc.FromConfigString(string(cfgData), filepath.Ext(hugoConfig))
}

func getGoPkgs(pkg string) []string {
	return nil
}

func writePackagePage() error {
	_ = `
---
---
`[1:]

	return nil
}

func runHugo(args ...string) error {
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
