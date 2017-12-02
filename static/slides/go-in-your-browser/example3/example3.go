package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func main() {
	vecty.RenderBody(&Page{})
}

type Page struct {
	vecty.Core

	text string
}

func (p *Page) Render() *vecty.HTML {
	return elem.Body(
		vecty.Markup(vecty.Style("font-size", "20pt")),
		elem.Input(
			vecty.Markup(
				event.Input(func(e *vecty.Event) {
					p.text = e.Target.Get("value").String()
					p.text += " " + p.text
					vecty.Rerender(p)
				}),
			),
		),
		elem.Break(),
		vecty.Text(p.text),
	)
}
