package main

import (
	"syscall/js"

	"github.com/arcadia-rose/shopify-hackdays-march-2022/internal/game"
)

var State game.State

func NewState() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		State = game.NewState()
		return State.ToMap()
	})
}

func GameLoop() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0].String()
		argc := args[1].Int()
		ids := []game.Id{}
		for i := 0; i < argc; i++ {
			index := i + 2 // starting index is 2
			ids = append(ids, game.Id(args[index].Int()))
		}

		State = game.GameLoop(State, game.Event(event), ids)

		return State.ToMap()
	})
}

func main() {
	js.Global().Set("NewState", NewState())
	js.Global().Set("GameLoop", GameLoop())
	<-make(chan bool)
}
