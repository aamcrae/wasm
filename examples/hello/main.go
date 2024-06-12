package main

import (
	"github.com/aamcrae/wasm"
)

func main() {
	w := wasm.GetWindow()
	w.SetTitle("Hello to Go wasm!")
	w.Display(new(wasm.HTML).H1("Hello, world!").String())
}
