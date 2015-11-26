// Package label is an example using go-frp modeled after the example found in
// the book "Functional Reactive Programming" by Stephen Blackheath and Anthony Jones:
// http://www.manning.com/books/functional-reactive-programming
package label

import (
	"log"

	h "github.com/gmlewis/go-frp/html"
	"honnef.co/go/js/dom"
)

// MODEL

type Model string

func (m Model) String() string { return string(m) }

// UPDATE

type Action func(Model, dom.Event) Model

func Updater(model Model) func(Action, dom.Event) Model {
	return func(action Action, event dom.Event) Model { return model.Update(action, event) }
}
func (m Model) Update(action Action, event dom.Event) Model { return action(m, event) }

func Keypress(model Model, event dom.Event) Model {
	log.Printf("Keypress: model=%v, event=%#v", model, event)
	return model
}

// VIEW

func (m Model) View(rootUpdateFunc, wrapFunc interface{}) h.HTML {
	return h.Div(
		h.Input(string(m)).OnKeypress(rootUpdateFunc, Updater(m), Keypress),
		h.Label(string(m)),
	)
}
