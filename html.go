//go:generate go run gen/main.go --input=tags --output=tags.go

package wasm

import (
	"strconv"
	"strings"

	_ "syscall/js"
)

type attr string
type flag int

const (
	f_drop flag = 1 << iota
	f_no_open
	f_no_close
	f_no_arg
	f_tags
)

type frag struct {
	strings.Builder
}

type HTML struct {
	frag
}

func NewHTML() *HTML {
	return new(HTML)
}

// Short hand write
func (h *HTML) Wr(s ...any) {
	h.frag.wrAll(s, false)
}

// Text composes the list of elements to a single string
// No tags are allowed, and a string is returned.
func (h *HTML) Text(s ...any) string {
	attrs, other, flags := unpack(s)
	if (flags & f_drop) != 0 {
		return ""
	}
	if (flags&f_tags) != 0 || len(attrs) > 0 {
		panic("Tags in text")
	}
	f := new(frag)
	f.wrAll(other, false)
	return f.String()
}

/*
 * Modifiers, which set flags to control
 * the behaviour.
 */

// If will drop this element if the condition is false
func (h *HTML) If(c bool) flag {
	if !c {
		return f_drop
	} else {
		return 0
	}
}

// For non-empty tags, do not generate the closing tag
func (h *HTML) Open() flag {
	return f_no_close
}

// For non-empty tags, generate the closing tag.
func (h *HTML) Close() flag {
	return f_no_open
}

func tag(nm string, elems []any) *frag {
	return wrTag(nm, elems, false)
}

func emptyTag(nm string, elems []any) *frag {
	return wrTag(nm, elems, true)
}

func wrTag(nm string, elems []any, empty bool) *frag {
	f := new(frag)
	attrs, other, flags := unpack(elems)
	if (flags & f_drop) != 0 {
		return f
	}
	if (flags & f_no_open) == 0 {
		f.WriteRune('<')
		f.WriteString(nm)
		f.wrAll(attrs, true)
		f.WriteRune('>')
	}
	f.wrAll(other, false)
	if !empty && (flags&f_no_close) == 0 {
		f.WriteString("</")
		f.WriteString(nm)
		f.WriteRune('>')
	}
	return f
}

func attrNoArg(nm string, elems []any) attr {
	if len(elems) == 0 {
		elems = []any{flag(f_no_arg)}
	}
	return attribute(nm, elems)
}

func attribute(nm string, elems []any) attr {
	attrs, other, flags := unpack(elems)
	if (flags&f_tags) != 0 || len(attrs) > 0 {
		panic("Illegal attribute")
	}
	if (flags & f_drop) != 0 {
		return ""
	}
	f := new(frag)
	// Leave a space before each attribute.
	f.WriteString(nm)
	if (flags & f_no_arg) == 0 {
		f.WriteString("=\"")
		f.wrAll(other, false)
		f.WriteString("\"")
	}
	return attr(f.String())
}

func unpack(s []any) ([]any, []any, flag) {
	var other []any
	var attrs []any
	var flags flag
	for _, ele := range s {
		switch v := ele.(type) {
		case *frag:
			flags |= f_tags
			other = append(other, ele)
		case attr:
			attrs = append(attrs, ele)
		case flag:
			flags |= v
		default:
			other = append(other, ele)
		}
	}
	return attrs, other, flags
}

func (f *frag) wrAll(s []any, space bool) {
	for _, ele := range s {
		if space {
			f.WriteRune(' ')
		}
		f.wr(ele)
	}
}

func (f *frag) wr(s any) {
	switch v := s.(type) {
	case string:
		f.WriteString(v)
	case attr:
		f.WriteString(string(v))
	case []byte:
		f.Write(v)
	case rune:
		f.WriteRune(v)
	case int:
		f.WriteString(strconv.FormatInt(int64(v), 10))
	case *frag:
		f.WriteString(v.String())
	default:
		panic("wr: Unknown type")
	}
}
