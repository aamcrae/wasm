// DO NOT EDIT - generated file
package wasm

func (h *HTML) H1(elems ...any) *frag {
	return tag("h1", elems)
}

func (h *HTML) H2(elems ...any) *frag {
	return tag("h2", elems)
}

func (h *HTML) H3(elems ...any) *frag {
	return tag("h3", elems)
}

func (h *HTML) H4(elems ...any) *frag {
	return tag("h4", elems)
}

func (h *HTML) H5(elems ...any) *frag {
	return tag("h5", elems)
}

func (h *HTML) H6(elems ...any) *frag {
	return tag("h6", elems)
}

func (h *HTML) Div(elems ...any) *frag {
	return tag("div", elems)
}

func (h *HTML) A(elems ...any) *frag {
	return tag("a", elems)
}

func (h *HTML) Span(elems ...any) *frag {
	return tag("span", elems)
}

func (h *HTML) Ol(elems ...any) *frag {
	return tag("ol", elems)
}

func (h *HTML) Form(elems ...any) *frag {
	return tag("form", elems)
}

func (h *HTML) Label(elems ...any) *frag {
	return tag("label", elems)
}

func (h *HTML) Input(elems ...any) *frag {
	return tag("input", elems)
}

func (h *HTML) Ul(elems ...any) *frag {
	return tag("ul", elems)
}

func (h *HTML) Li(elems ...any) *frag {
	return tag("li", elems)
}

func (h *HTML) Table(elems ...any) *frag {
	return tag("table", elems)
}

func (h *HTML) Tbody(elems ...any) *frag {
	return tag("tbody", elems)
}

func (h *HTML) Tr(elems ...any) *frag {
	return tag("tr", elems)
}

func (h *HTML) Td(elems ...any) *frag {
	return tag("td", elems)
}

func (h *HTML) P(elems ...any) *frag {
	return tag("p", elems)
}

func (h *HTML) Br(elems ...any) *frag {
	return emptytag("br", elems)
}

func (h *HTML) Hr(elems ...any) *frag {
	return emptytag("hr", elems)
}

func (h *HTML) Link(elems ...any) *frag {
	return emptytag("link", elems)
}

func (h *HTML) Img(elems ...any) *frag {
	return emptytag("img", elems)
}

func (h *HTML) Alt(elems ...any) attr {
	return attribute("alt", elems)
}

func (h *HTML) Title(elems ...any) attr {
	return attribute("title", elems)
}

func (h *HTML) Src(elems ...any) attr {
	return attribute("src", elems)
}

func (h *HTML) For(elems ...any) attr {
	return attribute("for", elems)
}

func (h *HTML) Name(elems ...any) attr {
	return attribute("name", elems)
}

func (h *HTML) Onclick(elems ...any) attr {
	return attribute("onclick", elems)
}

func (h *HTML) Onsubmit(elems ...any) attr {
	return attribute("onsubmit", elems)
}

func (h *HTML) Href(elems ...any) attr {
	return attribute("href", elems)
}

func (h *HTML) Rel(elems ...any) attr {
	return attribute("rel", elems)
}

func (h *HTML) Type(elems ...any) attr {
	return attribute("type", elems)
}

func (h *HTML) Border(elems ...any) attr {
	return attribute("border", elems)
}

func (h *HTML) Summary(elems ...any) attr {
	return attribute("summary", elems)
}

func (h *HTML) Class(elems ...any) attr {
	return attribute("class", elems)
}

func (h *HTML) Id(elems ...any) attr {
	return attribute("id", elems)
}

func (h *HTML) Size(elems ...any) attr {
	return attribute("size", elems)
}

func (h *HTML) Style(elems ...any) attr {
	return attribute("style", elems)
}

func (h *HTML) Value(elems ...any) attr {
	return attribute("value", elems)
}

func (h *HTML) Download(elems ...any) attr {
	if len(s) == 0 {
		s = []any{flag(f_no_arg)}
	}
	return attribute("download", s)
}
