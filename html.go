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

// H1 builds a H1 element
func (h *HTML) H1(elems ...any) *frag {
	return tag("h1", elems)
}

// H2 builds a H2 element
func (h *HTML) H2(elems ...any) *frag {
	return tag("h2", elems)
}

// H3 builds a H3 element
func (h *HTML) H3(elems ...any) *frag {
	return tag("h3", elems)
}

// H4 builds a H4 element
func (h *HTML) H4(elems ...any) *frag {
	return tag("h4", elems)
}

// H5 builds a H5 element
func (h *HTML) H5(elems ...any) *frag {
	return tag("h5", elems)
}

// H6 builds a H6 element
func (h *HTML) H6(elems ...any) *frag {
	return tag("h6", elems)
}

// Img builds a Img element
func (h *HTML) Img(elems ...any) *frag {
	return tag("img", elems)
}

// Div builds a Div element
func (h *HTML) Div(elems ...any) *frag {
	return tag("div", elems)
}

// A builds an anchor element
func (h *HTML) A(elems ...any) *frag {
	return tag("a", elems)
}

// Span builds a span element
func (h *HTML) Span(elems ...any) *frag {
	return tag("span", elems)
}

// Ol builds a ordered list
func (h *HTML) Ol(elems ...any) *frag {
	return tag("ol", elems)
}

// Ul builds an unordered list
func (h *HTML) Ul(elems ...any) *frag {
	return tag("ul", elems)
}

// Li builds a list item element
func (h *HTML) Li(elems ...any) *frag {
	return tag("li", elems)
}

// Table builds a Table element
func (h *HTML) Table(elems ...any) *frag {
	return tag("table", elems)
}

// Tr builds a table row element
func (h *HTML) Tr(elems ...any) *frag {
	return tag("tr", elems)
}

// Td builds a table data element
func (h *HTML) Td(elems ...any) *frag {
	return tag("td", elems)
}

// P builds a paragraph element
func (h *HTML) P(elems ...any) *frag {
	return tag("p", elems)
}

// Empty elements

// Br builds a break element
func (h *HTML) Br(elems ...any) *frag {
	return emptyTag("br", elems)
}

// Hr builds a hr element
func (h *HTML) Hr(elems ...any) *frag {
	return emptyTag("br", elems)
}

// Link builds a link element
func (h *HTML) Link(elems ...any) *frag {
	return emptyTag("link", elems)
}

// Attributes

func (h *HTML) Alt(s ...any) attr {
	return attribute("alt", s)
}

func (h *HTML) Title(s ...any) attr {
	return attribute("title", s)
}

func (h *HTML) Src(s ...any) attr {
	return attribute("src", s)
}

func (h *HTML) Onclick(s ...any) attr {
	return attribute("onclick", s)
}

func (h *HTML) Href(s ...any) attr {
	return attribute("href", s)
}

func (h *HTML) Rel(s ...any) attr {
	return attribute("rel", s)
}

func (h *HTML) Type(s ...any) attr {
	return attribute("type", s)
}

func (h *HTML) Border(s ...any) attr {
	return attribute("border", s)
}

func (h *HTML) Summary(s ...any) attr {
	return attribute("summary", s)
}

func (h *HTML) Class(s ...any) attr {
	return attribute("class", s)
}

func (h *HTML) Id(s ...any) attr {
	return attribute("id", s)
}

func (h *HTML) Style(s ...any) attr {
	return attribute("style", s)
}

// If no arguments, skip setting the value.
func (h *HTML) Download(s ...any) attr {
	if len(s) == 0 {
		s = []any{flag(f_no_arg)}
	}
	return attribute("download", s)
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
	f.WriteRune(' ')
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
