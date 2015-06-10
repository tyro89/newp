package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

//go:generate go run scripts/generate_build.go

var builds = map[string]func(Build) error{
	GOPROJECT:   buildGo,
	RUBYPROJECT: buildRuby,
}

func parseFlags() (string, string) {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("newp needs to be invoked like: newp [lang] [name].")
	}
	return args[0], args[1]
}

func build(lang, name string) error {
	if f, ok := builds[strings.ToLower(lang)]; ok {
		return f(newBuild(name))
	}
	return fmt.Errorf("newp does not know how to build a project for: %s", lang)
}

func main() {
	lang, name := parseFlags()
	if err := build(lang, name); err != nil {
		log.Fatal(err)
	}
	log.Print("Project built successfully.")
}
