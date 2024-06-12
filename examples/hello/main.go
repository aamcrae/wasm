package main

import (
	"github.com/aamcrae/wasm"
)

func main() {
	w := html.GetWindow()
	w.SetTitle("Hello to Go wasm!")
	w.Display(w.HTML().H1("Hello, world!").String())
}
