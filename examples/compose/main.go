package main

import (
	"github.com/aamcrae/wasm"
)

func main() {
	w := html.GetWindow()
	w.SetTitle("Compositor examples!")
	h := new(html.HTML)
	h.Wr(h.H1("Compositor examples"))
	h.Wr(h.P("Text and numbers (e.g: ", 1234, " and ", 5678, ") can be intermingled", h.Br(),
		"as well as other elements like br"))
	h.Wr(h.P("and runes (", rune(0x21A7), ")"))
	h.Wr(h.P(h.Style("font-weight:bold; font-size: large"), "Here is a large, bold link to ", h.A(h.Href("../hello/index.html"), "Hello World")))
	h.Wr(h.H2("Tables are supported"))
	h.Wr(h.Table(h.Open(), h.Summary("times table"), h.Border(2)))
	for i := 1; i <= 10; i++ {
		h.Wr(h.Tr(h.Open()))
		for j := 1; j <= 10; j++ {
			h.Wr(h.Td(i * j))
		}
		h.Wr(h.Tr(h.Close()))
	}
	h.Wr(h.Table(h.Close()))
	w.Display(h.String())
}
