// +build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/magefile/mage/sh"
)

func Build() error {
	name := "olcolabs"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	gopath, err := sh.Output("go", "env", "GOPATH")
	if err != nil {
		return fmt.Errorf("can't determine GOPATH: %v", err)
	}
	paths := strings.Split(gopath, string([]rune{os.PathListSeparator}))
	bin := filepath.Join(paths[0], "bin")
	// specifically don't mkdirall, if you have an invalid gopath in the first
	// place, that's not on us to fix.
	if err := os.Mkdir(bin, 0700); err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create %q: %v", bin, err)
	}
	path := filepath.Join(bin, name)

	// we use go build here because if someone built with go get, then `go
	// install` turns into a no-op, and `go install -a` fails on people's
	// machines that have go installed in a non-writeable directory (such as
	// normal OS installs in /usr/bin)
	return sh.RunV("go", "build", "-o", path, "github.com/bigodines/olcolabs")
}

// Generates a new release.  Expects the TAG environment variable to be set,
// which will create a new tag with that name.
func Release() (err error) {
	if os.Getenv("TAG") == "" {
		return errors.New("MSG and TAG environment variables are required")
	}
	if err := sh.RunV("git", "tag", "-a", "$TAG"); err != nil {
		return err
	}
	if err := sh.RunV("git", "push", "origin", "$TAG"); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			sh.RunV("git", "tag", "--delete", "$TAG")
			sh.RunV("git", "push", "--delete", "origin", "$TAG")
		}
	}()
	return sh.RunV("goreleaser")
}

// tag returns the git tag for the current branch or "" if none.
func tag() string {
	s, _ := sh.Output("git", "describe", "--tags")
	return s
}

// hash returns the git hash for the current repo or "" if none.
func hash() string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return hash
}
