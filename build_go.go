package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"code.google.com/p/go/src/pkg/text/template"
)

const GOPROJECT = "go"

func buildGo(b Build) error {
	if err := buildShared(b); err != nil {
		return err
	}

	if err := buildGoBin(b); err != nil {
		return err
	}

	if err := buildGoGdb(b); err != nil {
		return err
	}

	if err := buildGoGdbTest(b); err != nil {
		return err
	}

	//	if err := buildGoMain(b); err != nil {
	//		return err
	//	}
	//
	//	if err := buildGoTest(b); err != nil {
	//		return err
	//	}
	return nil
}

func buildGoBin(b Build) error {
	return os.Mkdir(path.Join(b.Root, "bin"), os.FileMode(0755))
}

func buildGoGdb(b Build) error {
	template := `#!/bin/bash

set -x
set -e

ROOT=\"$(cd \"$( dirname \"${BASH_SOURCE[0]}\" )/..\" && pwd)\"
WORK=\"$(cd $ROOT && go build -work -gcflags '-N -l' 2>&1 | cut -c6-)\"

cd $ROOT && sudo gdb %s -d $WORK
cd $ROOT && rm %s

rm -rf $WORK
`
	data := fmt.Sprintf(template, b.NameSnakeCase, b.NameSnakeCase)
	return ioutil.WriteFile(path.Join(b.Root, "bin", "gdb.sh"), []byte(data), os.FileMode(0644))
}

func buildGoGdbTest(b Build) error {
	t, err := template.New("").Parse(`#!/bin/bash

set -x
set -e

ROOT=\"$(cd \"$( dirname \"${BASH_SOURCE[0]}\" )/..\" && pwd)\"
WORK=\"$(cd $ROOT && go test -work -c -gcflags '-N -l' 2>&1 | cut -c6-)\"

cd $ROOT && sudo gdb {{ .NameSnakeCase }}.test -d $WORK
cd $ROOT && rm {{ .NameSnakeCase }}.test

rm -rf $WORK
`)

	if err != nil {
		return err
	}

	buf = bytes.Buffer{}
	if err := t.Execute(buf, b); err != nil {
		return err
	}

	return ioutil.WriteFile(path.Join(b.Root, "bin", "gdb_test.sh"), buf.Bytes(), os.FileMode(0644))
}

//func goMain(path string) {
//	f := mkfile(path)
//	writeLine(f, "package main")
//	writeLine(f, "")
//	writeLine(f, "import \"fmt\"")
//	writeLine(f, "")
//	writeLine(f, "func message() string {")
//	writeLine(f, "	return \"Hello World!\"")
//	writeLine(f, "}")
//	writeLine(f, "")
//	writeLine(f, "func main() {")
//	writeLine(f, "	msg := message()")
//	writeLine(f, "	fmt.Println(msg)")
//	writeLine(f, "}")
//}
//
//func goTest(path string) {
//	f := mkfile(path)
//	writeLine(f, "package main")
//	writeLine(f, "")
//	writeLine(f, "import \"testing\"")
//	writeLine(f, "")
//	writeLine(f, "func TestMyFunction(t *testing.T) {")
//	writeLine(f, "	expected := \"Hello World!\"")
//	writeLine(f, "	actual := message()")
//	writeLine(f, "	if actual != expected {")
//	writeLine(f, "		t.Errorf(\"\\nACTUAL: %s\\nEXPECTED: %s\", actual, expected)")
//	writeLine(f, "	}")
//	writeLine(f, "}")
//}
