package html

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

type html struct {
	frag
}

// Short hand write
func (h *html) Wr(s ...any) {
	h.frag.wrAll(s, false)
}

// Text composes the list of elements to a single string
// No tags are allowed
func (h *html) Text(s ...any) *frag {
	f := new(frag)
	attrs, other, flags := unpack(s)
	if (flags & f_drop) != 0 {
		return f
	}
	if (flags&f_tags) != 0 || len(attrs) > 0 {
		panic("Tags in text")
	}
	f.wrAll(other, false)
	return f
}

// H1 builds a H1 element
func (h *html) H1(elems ...any) *frag {
	return tag("h1", elems)
}

// H2 builds a H2 element
func (h *html) H2(elems ...any) *frag {
	return tag("h2", elems)
}

// H3 builds a H3 element
func (h *html) H3(elems ...any) *frag {
	return tag("h3", elems)
}

// H4 builds a H4 element
func (h *html) H4(elems ...any) *frag {
	return tag("h4", elems)
}

// H5 builds a H5 element
func (h *html) H5(elems ...any) *frag {
	return tag("h5", elems)
}

// H6 builds a H6 element
func (h *html) H6(elems ...any) *frag {
	return tag("h6", elems)
}

// Img builds a Img element
func (h *html) Img(elems ...any) *frag {
	return tag("img", elems)
}

// Div builds a Div element
func (h *html) Div(elems ...any) *frag {
	return tag("div", elems)
}

// A builds an anchor element
func (h *html) A(elems ...any) *frag {
	return tag("a", elems)
}

// Span builds a span element
func (h *html) Span(elems ...any) *frag {
	return tag("span", elems)
}

// Table builds a Table element
func (h *html) Table(elems ...any) *frag {
	return tag("table", elems)
}

// Tr builds a table row element
func (h *html) Tr(elems ...any) *frag {
	return tag("tr", elems)
}

// Td builds a table data element
func (h *html) Td(elems ...any) *frag {
	return tag("td", elems)
}

// P builds a paragraph element
func (h *html) P(elems ...any) *frag {
	return tag("p", elems)
}

// Empty elements

// Br builds a break element
func (h *html) Br(elems ...any) *frag {
	return emptyTag("br", elems)
}

// Hr builds a hr element
func (h *html) Hr(elems ...any) *frag {
	return emptyTag("br", elems)
}

// Link builds a link element
func (h *html) Link(elems ...any) *frag {
	return emptyTag("link", elems)
}

// Attributes

func (h *html) Alt(s ...any) attr {
	return attribute("alt", s)
}

func (h *html) Title(s ...any) attr {
	return attribute("title", s)
}

func (h *html) Src(s ...any) attr {
	return attribute("src", s)
}

func (h *html) Onclick(s ...any) attr {
	return attribute("onclick", s)
}

func (h *html) Href(s ...any) attr {
	return attribute("href", s)
}

func (h *html) Rel(s ...any) attr {
	return attribute("rel", s)
}

func (h *html) Type(s ...any) attr {
	return attribute("type", s)
}

func (h *html) Border(s ...any) attr {
	return attribute("border", s)
}

func (h *html) Summary(s ...any) attr {
	return attribute("summary", s)
}

func (h *html) Class(s ...any) attr {
	return attribute("class", s)
}

func (h *html) Id(s ...any) attr {
	return attribute("id", s)
}

func (h *html) Style(s ...any) attr {
	return attribute("style", s)
}

// If no arguments, skip setting the value.
func (h *html) Download(s ...any) attr {
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
func (h *html) If(c bool) flag {
	if !c {
		return f_drop
	} else {
		return 0
	}
}

// For non-empty tags, do not generate the closing tag
func (h *html) Open() flag {
	return f_no_close
}

// For non-empty tags, generate the closing tag.
func (h *html) Close() flag {
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
