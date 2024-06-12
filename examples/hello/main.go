package main

import (
	"github.com/aamcrae/wasm"
)

func main() {
	w := html.GetWindow()
	w.SetTitle("Hello to Go wasm!")
	w.Display(new(html.HTML).H1("Hello, world!").String())
}
