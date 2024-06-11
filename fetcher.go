package html

import (
	"errors"
	"sync"
	"sync/atomic"

	"syscall/js"
)

// Fetcher is an interface to the JS fetch API.
// It is suitable for use with tinygo.
type Fetcher struct {
	w      *Window
	wg     sync.WaitGroup
	data   []byte
	err    error
	errF   js.Func
	respF  js.Func
	readyF js.Func
	ready  atomic.Bool
}

// NewFetcher returns a new instance of a fetcher.
// The JS fetch api is called to start the reading of the file.
func NewFetcher(w *Window, url string) *Fetcher {
	f := &Fetcher{w: w}
	fPromise := w.window.Call("fetch", js.ValueOf(url))
	f.wg.Add(1)
	f.respF = js.FuncOf(f.response)
	f.errF = js.FuncOf(f.reject)
	f.readyF = js.FuncOf(f.dataReady)
	fPromise.Call("then", f.respF, f.errF)
	return f
}

// Get returns the data retrieved, or an error if the file
// could not be accessed.
func (f *Fetcher) Get() ([]byte, error) {
	f.wg.Wait()
	f.errF.Release()
	f.respF.Release()
	f.readyF.Release()
	return f.data, f.err
}

// Ready returns true if Get will succeed without block
func (f *Fetcher) Ready() bool {
	return f.ready.Load()
}

func (f *Fetcher) reject(this js.Value, args []js.Value) any {
	f.err = errors.New("Rejected")
	f.wg.Done()
	f.ready.Store(true)
	return nil
}

func (f *Fetcher) response(this js.Value, args []js.Value) any {
	ok := args[0].Get("ok")
	if js.ValueOf(ok).Bool() == true {
		// Accessing the data is done through a second promise.
		p2 := args[0].Call("text")
		p2.Call("then", f.readyF, f.errF)
		return p2
	}
	// The file could not be accessed.
	f.err = errors.New(js.ValueOf(args[0].Get("status")).String())
	f.wg.Done()
	f.ready.Store(true)
	return nil
}

func (f *Fetcher) dataReady(this js.Value, args []js.Value) any {
	f.data = []byte(js.ValueOf(args[0]).String())
	f.wg.Done()
	f.ready.Store(true)
	return nil
}
