package main

import (
	"syscall/js"

	"github.com/arcadia-rose/shopify-hackdays-march-2022/internal/game"
)

func NewState() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		s := game.NewState()
		return map[string]interface{}{
			"PlayerState": s.PlayerState.ToMap(),
		}
	})
}

func main() {
	js.Global().Set("NewState", NewState())
	<-make(chan bool)
}
