package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Build struct {
	NameSnakeCase string
	NameCamelCase string
	Root          string
}

// Parses a snake case and/or a camel case string and extracts its parts.
func parseParts(s string) []string {
	parts := []string{""}

	for _, char := range s {

		// Deal with underscore char
		if string(char) == "_" {
			if parts[len(parts)-1] != "" {
				parts = append(parts, "")
			}
			continue
		}

		// Deal with uppercase char
		if unicode.IsUpper(char) {
			if parts[len(parts)-1] != "" {
				parts = append(parts, strings.ToLower(string(char)))
			} else {
				parts[len(parts)-1] += strings.ToLower(string(char))
			}
			continue
		}

		// Deal with lower case char
		parts[len(parts)-1] += string(char)
	}

	if parts[len(parts)-1] == "" {
		return parts[0 : len(parts)-1]
	}

	return parts
}

func snakeCase(name string) string {
	parts := parseParts(name)
	return strings.Join(parts, "_")
}

func camelCase(name string) string {
	parts := parseParts(name)
	for i, part := range parts {
		p := strings.ToUpper(string(part[0]))
		if len(part) > 1 {
			p += part[1:]
		}
		parts[i] = p
	}
	return strings.Join(parts, "")
}

func getRoot(name string) string {
	root, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	parts := parseParts(name)
	project := strings.Join(parts, "-")

	return fmt.Sprintf("%s/%s", strings.TrimSuffix(root, "/"), project)
}

func newBuild(name string) Build {
	return Build{
		NameSnakeCase: snakeCase(name),
		NameCamelCase: camelCase(name),
		Root:          getRoot(name),
	}
}
