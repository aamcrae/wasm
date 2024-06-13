package main

import (
	"fmt"

	"syscall/js"

	"github.com/aamcrae/wasm"
)

func main() {
	w := wasm.GetWindow()
	w.SetTitle("Compositor examples!")
	w.AddStyle("body { background-color: #CCCCCC} .cell { text-align: right; width: 2em;}")
	w.AddJSFunction("runTables", func(js.Value, []js.Value) any {
		TimesTable(w)
		return nil
	})
	h := new(wasm.HTML)
	h.Wr(h.H1("Compositor examples"))
	h.Wr(h.P("Text and numbers (e.g: ", 1234, " and ", 5678, ") can be intermingled", h.Br(),
		"as well as other elements like br"))
	h.Wr(h.P("and runes (", rune(0x21A7), ")"))
	h.Wr(h.P(h.Style("font-weight:bold; font-size: large"), "Here is a large, bold link to ", h.A(h.Href("../hello/index.html"), "Hello World")))
	h.Wr(h.H2("Inputs and tables are supported"))
	h.Wr(h.Span(h.Input(h.Id("max"), h.Type("number"), h.Value("10")),
		h.Button(h.Type("button"), h.Onclick("runTables()"), "Run")))
	h.Wr(h.P(h.Table(h.Summary("times table"), h.Border(2), h.Tbody(h.Id("data"))).String()))
	h.Wr(h.P("Lists are supported as well"))
	h.Wr(h.Ol(h.Li("item number one"), h.Li("Item number two"), h.Li("Item number three")))
	w.Display(h.String())
	w.OnKey(w.GetById("max"), func(key string) {
		if key == "Enter" {
			TimesTable(w)
		}
	})
	TimesTable(w)
	w.Wait()
}

func TimesTable(w *wasm.Window) {
	d := w.GetById("data")
	mValue := w.GetById("max")
	var max int
	fmt.Sscanf(mValue.Get("value").String(), "%d", &max)
	if max < 1 {
		max = 1
		mValue.Set("value", js.ValueOf("1"))
		w.Alert("Minimum value is 1")
	}
	if max > 30 {
		max = 30
		mValue.Set("value", js.ValueOf("30"))
		w.Alert("Maximum value is 30")
	}
	h := new(wasm.HTML)
	for i := 1; i <= max; i++ {
		h.Wr(h.Tr(h.Open()))
		for j := 1; j <= max; j++ {
			h.Wr(h.Td(h.Class("cell"), i*j))
		}
		h.Wr(h.Tr(h.Close()))
	}
	d.Set("innerHTML", h.String())
}
