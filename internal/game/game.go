package game

import "strconv"

type Id uint

func (i Id) ToString() string {
	return strconv.FormatInt(int64(i), 10)
}

type Event string

type Room struct {
}

type Item struct {
	Name   string
	Amount uint
}

func (i Item) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Name":   i.Name,
		"Amount": i.Amount,
	}
}

type Player struct {
	Insight   uint
	Position  Id
	Inventory map[Id]Item
}

type State struct {
	PlayerState   Player
	EventHandlers map[Event]func(State, []Id) State
}

func (s State) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"PlayerState": s.PlayerState.ToMap(),
	}
}

var ItemRegistry = map[Id]Item{
	Id(0): Item{
		Name:   "Key",
		Amount: 1,
	},
}

func GameLoop(state State, event Event, ids []Id) State {
	handler, found := state.EventHandlers[event]
	if found {
		return handler(state, ids)
	}
	return state
}

func NewPlayer() Player {
	return Player{
		Insight:   0,
		Position:  Id(0),
		Inventory: map[Id]Item{},
	}
}

func (p Player) ToMap() map[string]interface{} {
	items := map[string]interface{}{}
	for id, item := range p.Inventory {
		items[id.ToString()] = item.ToMap()
	}

	return map[string]interface{}{
		"Insight":   p.Insight,
		"Position":  int(p.Position),
		"Inventory": items,
	}
}

func NewState() State {
	return State{
		PlayerState: NewPlayer(),
		EventHandlers: map[Event]func(State, []Id) State{
			Event("collectItem"): collectItem,
		},
	}
}

func collectItem(state State, itemIds []Id) State {
	for _, itemId := range itemIds {
		state.PlayerState.Inventory[itemId] = ItemRegistry[itemId]
	}

	return state
}
