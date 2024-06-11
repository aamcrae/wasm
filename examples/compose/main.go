package main

import (
	"strings"

	h "github.com/aamcrae/wasm"
)

func main() {
	w := h.GetWindow()
	w.SetTitle("Compositor examples!")
	var b strings.Builder
	b.WriteString(h.H1("Compositor examples"))
	b.WriteString(h.P("Text and numbers (e.g: ", 1234, " and ", 5678, ") can be intermingled",
		h.Br(), "as well as other elements like br", h.Br()))
	b.WriteString(h.Text("and runes (", rune(0x21A7), ")"))
	b.WriteString(h.H2("Tables are supported"))
	b.WriteString(h.Table(h.Open(), h.Summary("times table"), h.Border(2)))
	for i := 1; i <= 10; i++ {
		b.WriteString(h.Tr(h.Open()))
		for j := 1; j <= 10; j++ {
			b.WriteString(h.Td(i * j))
		}
		b.WriteString(h.Tr(h.Close()))
	}
	b.WriteString(h.Table(h.Close()))
	w.Display(b.String())
}
