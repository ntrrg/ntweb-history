// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gohugoio/hugo/commands"
	"github.com/gohugoio/hugo/config"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default = Build

	hugoVersion = "0.69.0"
	hugoPort    = "1313"
	hugoConfig  = "config.yaml"

	cfg config.Provider

	buildFuncs []interface{}
	cleanFuncs []interface{}
	lintFuncs  []interface{}
)

func Build() error {
	resp := commands.Execute([]string{})

	return resp.Err
}

func BuildAll() error {
	mg.Deps(buildFuncs...)

	return Build()
}

type BumpVersion mg.Namespace

func (BumpVersion) Hugo() error {
	fn := func(path string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err

		case findString(path,
			".git",
			".mage/output", ".mage/vendor",
			"archetypes", "assets", "content", "data", "i18n", "layouts",
			"public", "resources",
		):

			return filepath.SkipDir

		case
			info.IsDir(),
			strings.HasSuffix(path, ".swp"),
			findString(path, ".mage/go.sum"):

			return nil
		}

		fd, err := os.Open(path)
		if err != nil {
			return err
		}

		defer fd.Close()

		s := bufio.NewScanner(fd)

		for s.Scan() {
			if bytes.Contains(s.Bytes(), []byte(hugoVersion)) {
				fmt.Println(path)
				break
			}
		}

		if err := s.Err(); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		return nil
	}

	return filepath.Walk(".", fn)
}

func Clean() error {
	mg.Deps(cleanFuncs...)

	targets := []string{".mage/output", "public", "resources"}

	for _, target := range targets {
		if err := sh.Rm(filepath.Clean(target)); err != nil {
			return err
		}
	}

	return nil
}

func Lint() error {
	mg.SerialDeps(lintFuncs...)

	files, err := filepath.Glob(".mage/*.go")
	if err != nil {
		return err
	}

	args := []string{"-d", "-e", "-s"}
	args = append(args, files...)
	return sh.RunV("gofmt", args...)
}

func Run() error {
	resp := commands.Execute([]string{
		"server", "-D", "-E", "-F", "--noHTTPCache", "--i18n-warnings",
		"--disableFastRender", "--bind", "0.0.0.0", "--port", hugoPort,
		"--baseUrl", "/", "--appendPort=false",
	})

	return resp.Err
}

func init() {
	var err error

	data, err := ioutil.ReadFile(hugoConfig)
	cfg, err = config.FromConfigString(string(data), filepath.Ext(hugoConfig))
	if err != nil {
		panic(err)
	}
}

func findString(s string, list ...string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}

	return false
}
