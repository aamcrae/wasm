# wasm Go library support

This library contains support for running Go programs under Web Assembly.

It consists of a window/DOM interface object, and simple compositor functions
used for building dynamic HTML. One of the goals is to be usable with [tinygo](https://tinygo.org/),
which does not support the full range of Go's standard library.

[Examples](examples/server) are available to show how the library can be used.

# Window

Window is a simple interface to the browser DOM, allowing callbacks
to be added for keyboard shortcuts, window resizing, swipe gestures:

```
	w := wasm.Window()
	w.OnSwipe(func (d wasm.Direction) bool {
		if d == wasm.Right {
			...
			return true // Right swipe is handled
		}
		return false // Not interested in other swipes, use default action
	})
	w.OnResize(resized)
	w.SetTitle("My window")
	w.Wait()
```

# Compositor

The compositor allows HTML pages or fragments to be generated in a
easy and consistent flow of functions. Tags are automatically closed
where required, and common attributes are supported. For example:

```
import (
	"github.com/aamcrae/wasm"
)
...
	w := wasm.Window()
	h := wasm.NewHTML()
	h.Wr(h.H1("Title Page"))
	h.Wr(h.A(h.Href("page/index.html"), h.Id("myid"), h.Img(h.Class("image"), h.Src("flower.jpg"), h.Alt("Flower"))))
	h.Wr(h.Span(h.Class("myspan"), h.Open()) // Don't close tag
	h.Wr(h.P("This is a paragraph", h.Br("with a break in it)))
	h.Wr(h.Text("Combining numbers ", 12345, ", runes ", ' ', rune(0x21A7), " and strings"))
	h.Wr(h.Span(h.Close())) // Now add the closing tag for span
	w.Display(h.String())
```

Modifiers and conditionals are allowed so that the flow can be maintained when
deeply embedded functions are used (without resorting to if/else control flow):

```
	// Only display title if it is not empty
	h.Wr(h.H1(h.If(len(title) > 0), title))
	h.Wr(h.A(h.Open(), h.Id("id", i), h.Href("#")))
	// complicated code to generate anchor
	h.Wr(h.A(h.Close()))
```

Be aware that complete syntax checking is **not** performed by the compositor e.g there
is nothing stopping incorrect attributes being used in tags, or tags not being closed etc.

# Fetcher

Fetcher is an interface to the Javascript [fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)
API. This interface is usable with tinygo (which does not support net/http).
It allows concurrent fetching of resources:

```
	// Start fetching all the files required
	f1 := w.Fetcher("data/file1")
	f2 := w.Fetcher("data/file2")
	f3 := w.Fetcher("data/file3")
	...
	go func() {
		val, err := f3.Get() // Blocks until file3 is read.
		...
	}
	// Retrieve data only when available
	if f1.Ready() {
		data1, err := f1.Get()
		...
	}
	data2, err := f2.Get() // blocks until ready
```
