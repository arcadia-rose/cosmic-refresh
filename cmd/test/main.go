package main

import (
  "fmt"

  "github.com/arcadia-rose/shopify-hackdays-march-2022/internal/game"
)

func main() {
  state := game.NewState()
  evt := game.Event("collectItem")
  key := []game.Id{game.Id(0)}

  fmt.Printf("Number of items in inventory: %d\n", len(state.PlayerState.Inventory))

  state = game.GameLoop(state, evt, key)
  
  fmt.Printf("Number of items in inventory: %d\n", len(state.PlayerState.Inventory))
}
