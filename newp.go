package main

import "fmt"

const (
	dirPerm  uint32 = 0755
	filePerm uint32 = 0644
)

var (
	lang          string
	path          string
	name          string
	nameSnakeCase string
	nameCamelCase string
)

func main() {
	lang, name = parseFlags()

	nameSnakeCase = snakeCase(name)
	nameCamelCase = camelCase(name)

	path = projectPath(nameSnakeCase)

	initializeProject()

	switch {
	case "go" == lang:
		buildGoProject()
	case "ruby" == lang:
		buildRubyProject()
	default:
		fmt.Printf("Unsupported language: %s. Initialized empty project.\n", lang)
	}
}
