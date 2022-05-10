package main

import (
  "encoding/json"
	"syscall/js"

	"github.com/arcadia-rose/shopify-hackdays-march-2022/internal/game"
)

var State game.State

func stateAsJSON() string {
  view := State.View()
  encoded, _ := json.Marshal(view)
  return string(encoded)
}

func NewState() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		State = game.NewState()
		
    return stateAsJSON()
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

		return stateAsJSON()
	})
}

func main() {
	js.Global().Set("NewState", NewState())
	js.Global().Set("GameLoop", GameLoop())
	<-make(chan bool)
}
