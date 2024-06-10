package main

import (
	h "github.com/aamcrae/wasm"
)

func main() {
	w := h.GetWindow()
	w.SetTitle("Hello to Go wasm!")
	w.Display(h.H1("Hello, world!"))
}
