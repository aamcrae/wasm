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
)

// Text composes the list of elements to a single string
func Text(s ...any) string {
	_, other, flags := unpack(s)
	if (flags & f_drop) != 0 {
		return ""
	}
	var b strings.Builder
	wrAll(&b, other, false)
	return b.String()
}

// H1 builds a H1 element
func H1(elems ...any) string {
	return tag("h1", elems)
}

// H2 builds a H2 element
func H2(elems ...any) string {
	return tag("h2", elems)
}

// H3 builds a H3 element
func H3(elems ...any) string {
	return tag("h3", elems)
}

// H4 builds a H4 element
func H4(elems ...any) string {
	return tag("h4", elems)
}

// H5 builds a H5 element
func H5(elems ...any) string {
	return tag("h5", elems)
}

// H6 builds a H6 element
func H6(elems ...any) string {
	return tag("h6", elems)
}

// Img builds a Img element
func Img(elems ...any) string {
	return tag("img", elems)
}

// Div builds a Div element
func Div(elems ...any) string {
	return tag("div", elems)
}

// A builds an anchor element
func A(elems ...any) string {
	return tag("a", elems)
}

// Span builds a span element
func Span(elems ...any) string {
	return tag("span", elems)
}

// Table builds a Table element
func Table(elems ...any) string {
	return tag("table", elems)
}

// Tr builds a table row element
func Tr(elems ...any) string {
	return tag("tr", elems)
}

// Td builds a table data element
func Td(elems ...any) string {
	return tag("td", elems)
}

// P builds a paragraph element
func P(elems ...any) string {
	return tag("p", elems)
}

// Empty elements

// Br builds a break element
func Br(elems ...any) string {
	return emptyTag("br", elems)
}

// Hr builds a hr element
func Hr(elems ...any) string {
	return emptyTag("br", elems)
}

// Link builds a link element
func Link(elems ...any) string {
	return emptyTag("link", elems)
}

// Attributes

func Alt(s ...any) attr {
	return attribute("alt", s)
}

func Title(s ...any) attr {
	return attribute("title", s)
}

func Src(s ...any) attr {
	return attribute("src", s)
}

func Onclick(s ...any) attr {
	return attribute("onclick", s)
}

func Href(s ...any) attr {
	return attribute("href", s)
}

func Rel(s ...any) attr {
	return attribute("rel", s)
}

func Type(s ...any) attr {
	return attribute("type", s)
}

func Border(s ...any) attr {
	return attribute("border", s)
}

func Summary(s ...any) attr {
	return attribute("summary", s)
}

func Class(s ...any) attr {
	return attribute("class", s)
}

func Id(s ...any) attr {
	return attribute("id", s)
}

func Style(s ...any) attr {
	return attribute("style", s)
}

// If no arguments, skip setting the value.
func Download(s ...any) attr {
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
func If(c bool) flag {
	if !c {
		return f_drop
	} else {
		return 0
	}
}

// For non-empty tags, do not generate the closing tag
func Open() flag {
	return f_no_close
}

// For non-empty tags, generate the closing tag.
func Close() flag {
	return f_no_open
}

func tag(nm string, elems []any) string {
	return wrTag(nm, elems, false)
}

func emptyTag(nm string, elems []any) string {
	return wrTag(nm, elems, true)
}

func wrTag(nm string, elems []any, empty bool) string {
	attrs, other, flags := unpack(elems)
	if (flags & f_drop) != 0 {
		return ""
	}
	var sb strings.Builder
	if (flags & f_no_open) == 0 {
		sb.WriteRune('<')
		sb.WriteString(nm)
		wrAll(&sb, attrs, true)
		sb.WriteRune('>')
	}
	wrAll(&sb, other, false)
	if !empty && (flags & f_no_close)==0 {
		sb.WriteString("</")
		sb.WriteString(nm)
		sb.WriteRune('>')
	}
	return sb.String()
}

func attribute(nm string, elems []any) attr {
	attrs, other, flags := unpack(elems)
	if (flags & f_drop) != 0 || len(attrs) > 0 {
		return ""
	}
	var sb strings.Builder
	// Leave a space before each attribute.
	sb.WriteRune(' ')
	sb.WriteString(nm)
	if (flags & f_no_arg) == 0 {
		sb.WriteString("=\"")
		wrAll(&sb, other, false)
		sb.WriteString("\"")
	}
	return attr(sb.String())
}

func unpack(s []any) ([]any, []any, flag) {
	var other []any
	var attrs []any
	var flags flag
	for _, ele := range s {
		switch v := ele.(type) {
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

func wrAll(sb *strings.Builder, s []any, space bool) {
	for _, ele := range s {
		if space {
			sb.WriteRune(' ')
		}
		wr(sb, ele)
	}
}

func wr(sb *strings.Builder, s any) {
	switch v := s.(type) {
	case string:
		sb.WriteString(v)
	case attr:
		sb.WriteString(string(v))
	case []byte:
		sb.Write(v)
	case rune:
		sb.WriteRune(v)
	case int:
		sb.WriteString(strconv.FormatInt(int64(v), 10))
	default:
		panic("wr: Unknown type")
	}
}
