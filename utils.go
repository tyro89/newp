package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func parseFlags() (string, string) {
	langPtr := flag.String("l", "", "The language of the project.")
	namePtr := flag.String("n", "", "The name of the project.")

	flag.Parse()

	lang := *langPtr
	name := *namePtr

	if len(lang) == 0 {
		fmt.Println("The '-l' parameter used to set the project language is required and must be at least one character long.")
		os.Exit(1)
	}

	if len(name) == 0 {
		fmt.Println("The '-n' parameter used to set the project name is required and must be at least one character long.")
		os.Exit(1)
	}

	return lang, name
}

func projectPath(name string) string {
	root, err := os.Getwd()

	if err != nil {
		fmt.Print("Unable to get working directory.")
		os.Exit(1)
	}

	return strings.TrimSuffix(root, "/") + "/" + name
}

func mkdir(path string) {
	os.Mkdir(path, os.FileMode(dirPerm))
}

func mkfile(path string) *os.File {
	f, err := os.Create(path)

	if err != nil {
		fmt.Println("Failed creating file.")
		os.Exit(1)
	}

	return f
}

func writeLine(f *os.File, line string) {
	f.WriteString(line + "\n")
}

func snakeCase(name string) string {
	parts := nameParts(name)

	snaked := ""
	for i, part := range parts {
		snaked += strings.ToLower(part)
		if i+1 < len(parts) {
			snaked += "_"
		}
	}

	return snaked
}

func camelCase(name string) string {
	parts := nameParts(name)

	camelized := ""
	for _, part := range parts {
		if len(part) > 0 {
			camelized += strings.ToUpper(string(part[0]))
			if len(part) > 1 {
				camelized += strings.ToLower(part[1:])
			}
		}
	}

	return camelized
}

func nameParts(name string) []string {
	parts := []string{}
	for _, char := range name {
		if len(parts) == 0 {
			parts = append(parts, string(char))
		} else if string(char) == "_" {
			parts = append(parts, "")
		} else if unicode.IsUpper(char) {
			parts = append(parts, strings.ToLower(string(char)))
		} else {
			parts[len(parts)-1] += string(char)
		}
	}

	return parts
}
