package wasm

import (
	"time"

	"syscall/js"
)

// Minimum movement to consider a touch event to be a swipe.
const minSwipe = 30

// Maximum time before a touch event stops being a swipe.
const maxSwipeTime = time.Millisecond * 300

// Direction is an enum indicating the swipe direction.
type Direction int

const (
	Right Direction = iota
	Left
	Up
	Down
)

// Window is the main structure for interfacing to the browser
type Window struct {
	Width, Height                int // Width and Height of window
	Window, Document, Head, Body js.Value
	// touch values
	startTime      time.Time
	startX, startY int
	endX, endY     int
	multiTouch     bool
}

// GetWindow creates a Window ready to interface to the browser.
func GetWindow() *Window {
	w := &Window{}
	w.Window = js.Global()
	w.Document = w.Window.Get("document")
	w.Head = w.Document.Get("head")
	w.Body = w.Document.Get("body")
	w.refreshSize()
	return w
}

// GetById returns the named JS element.
func (w *Window) GetById(id string) js.Value {
	return w.Document.Call("getElementById", id)
}

// Display sets the HTML onto the window.
func (w *Window) Display(s string) {
	w.Body.Set("innerHTML", s)
}

// LoadStyle adds a link to read the CSS file indicated
func (w *Window) LoadStyle(s string) {
	link := w.Document.Call("createElement", "link")
	link.Set("type", "text/css")
	link.Set("rel", "stylesheet")
	link.Set("href", s)
	w.Head.Call("appendChild", link)
}

// AddStyle adds the CSS string to the window directly.
func (w *Window) AddStyle(s string) {
	style := w.Document.Call("createElement", "style")
	style.Set("type", "text/css")
	ss := style.Get("styleSheet")
	if ss.Truthy() {
		ss.Set("cssText", s)
	} else {
		style.Call("appendChild", w.Document.Call("createTextNode", s))
	}
	w.Head.Call("appendChild", style)
}

// SetTitle sets the window title
func (w *Window) SetTitle(title string) *Window {
	w.Document.Set("title", title)
	return w
}

// Alert displays an alert box.
func (w *Window) Alert(a string) {
	w.Window.Call("alert", js.ValueOf(a))
}

// Goto navigates the browser to the URL.
func (w *Window) Goto(url string) {
	w.Window.Get("location").Set("href", js.ValueOf(url))
}

// OnSwipe registers a callback to be called for swipe events.
// If the callback handles the event, it returns true.
func (w *Window) OnSwipe(f func(Direction) bool) {
	touchStartJS := js.FuncOf(func(this js.Value, args []js.Value) any {
		t := args[0].Get("touches")
		if t.IsUndefined() {
			return nil
		}
		w.startTime = time.Now()
		w.startX = t.Index(0).Get("clientX").Int()
		w.startY = t.Index(0).Get("clientY").Int()
		w.endX = w.startX
		w.endY = w.startY
		// Ignore multi-touch gestures
		w.multiTouch = t.Length() > 1
		return nil
	})
	touchMoveJS := js.FuncOf(func(this js.Value, args []js.Value) any {
		e := args[0].Get("targetTouches")
		if e.IsUndefined() {
			return nil
		}
		if e.Length() == 1 {
			w.endX = e.Index(0).Get("clientX").Int()
			w.endY = e.Index(0).Get("clientY").Int()
		} else {
			w.multiTouch = true
		}
		return nil
	})
	touchEndJS := js.FuncOf(func(this js.Value, args []js.Value) any {
		if w.multiTouch {
			return nil
		}
		// If the swipe event lasted more than the preset max, do not
		// consider this a swipe event.
		if time.Now().Sub(w.startTime) > maxSwipeTime {
			return nil
		}
		e := args[0]
		x := w.startX - w.endX
		y := w.startY - w.endY
		var d Direction
		ax := abs(x)
		ay := abs(y)
		// Figure out up/down or left/right
		if ax > ay {
			if ax < minSwipe {
				return nil
			}
			if x > 0 {
				d = Left
			} else {
				d = Right
			}
		} else {
			if ay < minSwipe {
				return nil
			}
			if y > 0 {
				d = Up
			} else {
				d = Down
			}
		}
		if f(d) {
			// Don't process the default action if the
			// swipe was handled by the callback
			e.Call("preventDefault")
		}
		return nil
	})
	touchCancelJS := js.FuncOf(func(this js.Value, args []js.Value) any {
		return nil
	})
	w.Window.Call("addEventListener", "touchstart", touchStartJS)
	w.Window.Call("addEventListener", "touchend", touchEndJS)
	w.Window.Call("addEventListener", "touchmove", touchMoveJS)
	w.Window.Call("addEventListener", "touchcancel", touchCancelJS)
}

// OnResize registers a callback to be invoked when the window changes size.
func (w *Window) OnResize(f func()) {
	resizeJS := js.FuncOf(func(this js.Value, args []js.Value) any {
		w.refreshSize()
		f()
		return nil
	})
	w.Window.Call("addEventListener", "resize", resizeJS)
}

// OnKey registers a callback to be invoked when a key is pressed on this element.
func (w *Window) OnKey(elem js.Value, f func(key string)) {
	elem.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) any {
		f(js.ValueOf(args[0].Get("key")).String())
		return nil
	}))
}

// AddJSFunction registers a javascript function that invokes the Go function passed.
func (w *Window) AddJSFunction(name string, f func(js.Value, []js.Value) any) {
	w.Window.Set(name, js.FuncOf(f))
}

// refreshSize updates the width/height of the window.
func (w *Window) refreshSize() {
	w.Width = js.ValueOf(w.Body.Get("clientWidth")).Int()
	w.Height = js.ValueOf(w.Window.Get("innerHeight")).Int()
}

// Wait forces this thread to wait.
func (w *Window) Wait() {
	select {}
}

// Fetch retrieves the file from the server.
func (w *Window) Fetcher(file string) *fetcher {
	return newFetcher(w, file)
}

// abs returns the absolute value of a value
func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
