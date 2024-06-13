package main

import (
	"fmt"
	"strings"

	"syscall/js"

	html "github.com/aamcrae/wasm"
)

func main() {
	w := html.GetWindow()
	w.SetTitle("Fetcher examples")
	h := new(html.HTML)
	c := make(chan string)
	h.Wr(h.H1("Fetcher examples."))
	for i := 1; i < 6; i++ {
		h.Wr(h.H2("Nursery rhyme ", i))
		h.Wr(h.Div(h.Id("nr", i)))
		go get(w, c, fmt.Sprintf("file%d.dat", i))
	}
	w.Display(h.String())
	i := 1
	for {
		w.GetById(fmt.Sprintf("nr%d", i)).Set("innerHTML", js.ValueOf(<- c))
		i++
	}
}

func get(w *html.Window, c chan string, file string) {
	f := w.Fetcher(file)
	if b, err := f.Get(); err != nil {
		c <- fmt.Sprintf("%s error %v", file, err)
	} else {
		s := strings.TrimRight(string(b), "\n") // Remove trailing return
		h := new(html.HTML)
		h.Wr(h.P(h.Open()))
		for _, l := range strings.Split(s, "\n") {
			h.Wr(l, h.Br())
		}
		h.Wr(h.P(h.Close()))
		c <- h.String()
	}
}
