# wasm Go library support

This library contains support for running Go programs under Web Assembly.

It consists of a window/DOM interface object, and simple compositor functions
used for building dynamic HTML. One of the goals is to be usable with [tinygo](https://tinygo.org/),
which does bnot support the full range of Go's standard library.

[Examples](example) are available to show how the library can be used.

# Window

Window is a simple interface to the browser DOM, allowing callbacks
to be added for keyboard shortcuts, window resizing, swipe gestures:

```
	w := h.Window()
	w.OnSwipe(func (d h.Direction) bool {
		if d == h.Right {
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
	h "github.com/aamcrae/wasm"
)
...
	w := h.Window()
	vars b string.Builder
	b.WriteString(h.H1("Title Page"))
	b.WriteString(A(Href("page/index.html"), Id("myid"), Img(Class("image"), Src("my_image.jpg"), Alt("Flower"))))
	b/WriteString(Span(Class("myspan"), Open()) // Don't close tag
	b.WriteString(P("This is a paragraph", Br("with a break in it)))
	b.WriteString(Text("Combining numbers ", 12345, ", runes ", ' ', rune(0x21A7), " and strings"))
	b.WriteString(Span(Close())) // Now add the closing tag for span
	w.Display(b.String())
```

Modifiers and conditionals are allowed so that the functional flow can be maintained:

```
	// Only display title if it is not empty
	b.WriteString(H1(If(len(title) > 0), title))
	b.WriteString(A(Open(), Id("id", i), Href("#")))
    // complicated code to generate anchor
	b.WritesString(A(Close()))
```

# Fetcher

Fetcher is an interface to the Javascript [fetch](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)
API which is usable with tinygo (which does not support net/http).
It allows concurrent fetching of multiple files:

```
	// Start fetching all the files required
	f1 := NewFetcher(w, "data/file1")
	f2 := NewFetcher(w, "data/file2")
	f3 := NewFetcher(w, "data/file3")
	...
	// Retrieve data only when available
	if f1.Ready() {
		data1, err := f1.Get()
		...
	}
	data2, err := f2.Get() // blocks until ready
```
