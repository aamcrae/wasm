package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

var input = flag.String("input", "tags", "Input YAML file")
var output = flag.String("output", "tags.go", "Output file")
var pkg = flag.String("package", "wasm", "Package to use")

type Name struct {
	Name  string
	Cname string
}

const tagTempl = `func (h *HTML) {{.Cname}}(elems ...any) *frag {
	return tag("{{.Name}}", elems)
}

`
const emptyTagTempl = `func (h *HTML) {{.Cname}}(elems ...any) *frag {
	return emptyTag("{{.Name}}", elems)
}

`
const attrTempl = `func (h *HTML) {{.Cname}}(elems ...any) attr {
	return attribute("{{.Name}}", elems)
}

`
const attrNoArgTempl = `func (h *HTML) {{.Cname}}(elems ...any) attr {
	return attrNoArg("{{.Name}}", elems)
}

`

// Fix order of types
var order []string = []string{"tags", "emptytags", "attributes", "attr-noarg"}
var templates []string = []string{tagTempl, emptyTagTempl, attrTempl, attrNoArgTempl}

func main() {
	flag.Parse()

	data, err := os.ReadFile(*input)
	if err != nil {
		log.Fatalf("%s: %v", *input, err)
	}
	m := make(map[string][]string)
	err = yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("%s: %v", *input, err)
	}
	of, err := os.Create(*output)
	if err != nil {
		log.Fatalf("%s: %v", *output, err)
	}
	defer of.Close()
	fmt.Fprintln(of, "// DO NOT EDIT - generated file")
	fmt.Fprintf(of, "package %s\n\n", *pkg)
	c := cases.Title(language.English)

	for i, k := range order {
		t := template.Must(template.New(k).Parse(templates[i]))
		for _, n := range m[k] {
			//n := e.(string)
			err := t.Execute(of, Name{Name: n, Cname: c.String(n)})
			if err != nil {
				log.Fatalf("%s: %v", *output, err)
			}
		}
	}
}
