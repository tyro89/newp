package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type Template struct {
	Name string
	Data string
}

type BuildFile struct {
	Date      string
	Lang      string
	Templates []Template
}

func (b BuildFile) drop(filename string) error {
	payload := "// This file was generated on " + b.Date + ".\n"
	payload += "// Do not modify this file directly. To generate\n"
	payload += "// an updated version run \"go generate\".\n"
	payload += "\n"
	payload += "package main\n"
	payload += "\n"
	payload += "const (\n"

	for _, t := range b.Templates {
		payload += "	" + t.Name + " = `" + t.Data + "`\n"
	}

	payload += ")\n"
	payload += "\n"
	payload += "func init() {\n"
	payload += "	register(\"" + b.Lang + "\", build" + b.Lang + ")\n"
	payload += "}\n"
	payload += "\n"
	payload += "func build" + b.Lang + "(b Build) error {\n"

	for _, t := range b.Templates {
		payload += "	drop(b, \"" + path.Join(b.Lang, t.Name) + "\", " + t.Name + ")\n"
	}

	payload += "}\n"

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	_, err = f.WriteString(payload)
	return err
}

func langBuildFile(root string) (BuildFile, error) {

}

func main() {
	langFiles := make([]BuildFile, 0)
	sharedFile := BuildFile{
		Date:      "",
		Lang:      "",
		Templates: make([]Template, 0),
	}

	fileInfs, err := osutil.ReadDir("templates")
	if err != nil {
		log.Fatal(err)
	}

	for _, fs := range fileInfs {
		filepath = path.Join("templates", fs.Name())
		if fs.IsDir() {
			bf, err := langBuildFile(filepath)
			if err != nil {
				return err
			}
			langFiles = append(langFiles, bf)
			continue
		}

		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		t := Template{
			Name: fs.Name(),
			Data: data,
		}

		sharedFile.Templates = append(sharedFile.Templates, t)
	}

	// Read stuff and set up build files

	if err := sharedFile.drop("build_shared.go"); err != nil {
		log.Fatal(err)
	}

	for _, f := range langFiles {
		if err := f.drop(fmt.Sprintf("build_%s.go", f.Lang)); err != nil {
			log.Fatal(err)
		}
	}
}
