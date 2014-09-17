package main

func initializeProject() {
	mkdir(path)
	readme(path + "/README.md")
}

func readme(path string) {
	f := mkfile(path)
	writeLine(f, "# "+nameCamelCase)
	writeLine(f, "")
	writeLine(f, "Welcome! Proper README is on the way.")
}
