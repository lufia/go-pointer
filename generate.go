//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"unicode"
)

var (
	flagType = flag.String("type", "", "`name` of the builtin type")
)

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		name := filepath.Base(os.Args[0])
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "usage: %s [options] [data ...]\n\n", name)
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	if *flagType == "" {
		flag.Usage()
	}
	err := generate(&Target{
		Type:     *flagType,
		Name:     title(*flagType),
		TestData: flag.Args(),
	})
	if err != nil {
		log.Fatalf("cannot generate: %v", err)
	}
}

func title(s string) string {
	if s == "" {
		return ""
	}
	u := []rune(s)
	return string(unicode.ToUpper(u[0])) + string(u[1:])
}

//go:embed *.tmpl
var content embed.FS

var (
	funcTmpl = template.Must(template.ParseFS(content, "func.tmpl"))
	testTmpl = template.Must(template.ParseFS(content, "test.tmpl"))
)

type Target struct {
	Type     string
	Name     string
	TestData []string
}

func generate(t *Target) error {
	if err := execute(t.Type+".go", funcTmpl, t); err != nil {
		return err
	}
	if err := execute(t.Type+"_test.go", testTmpl, t); err != nil {
		return err
	}
	return nil
}

func execute(name string, tmpl *template.Template, t *Target) error {
	var w bytes.Buffer
	if err := tmpl.Execute(&w, t); err != nil {
		return err
	}
	data, err := format.Source(w.Bytes())
	if err != nil {
		return err
	}
	return os.WriteFile(name, data, 0664)
}
