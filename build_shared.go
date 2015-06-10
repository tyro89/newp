package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func buildShared(b Build) error {
	if err := buildRootDir(b); err != nil {
		return err
	}

	if err := buildReadme(b); err != nil {
		return err
	}

	return nil
}

func buildRootDir(b Build) error {
	return os.Mkdir(b.Root, os.FileMode(0755))
}

func buildReadme(b Build) error {
	template := `# %s

Welcome! Proper README is on the way.
`
	data := fmt.Sprintf(template, b.NameCamelCase)
	return ioutil.WriteFile(path.Join(b.Root, "README"), []byte(data), os.FileMode(0644))
}
