package game

type Id uint

type Event string

type Room struct {
}

type Item struct {
  Name   string `json:"name"`
  Amount uint   `json:"amount"`
}

type Player struct {
  Insight   uint        `json:"insight"`
  Position  Id          `json:"position"`
  Inventory map[Id]Item `json:"inventory"`
}

// Internal representation of the game state only usably on the Go side.
type State struct {
	PlayerState   Player
	EventHandlers map[Event]func(State, []Id) State
}

func (s State) View() StateView {
  return StateView {
    PlayerState: s.PlayerState,
  }
}

// Shared representation of the game state meant to be communicated between JS & Go.
// The non-serializable components of `State` must be reconstructed for internal use.
type StateView struct {
  PlayerState Player `json:"player"`
}

func (s StateView) State() State {
  newState := NewState()
  newState.PlayerState = s.PlayerState
  return newState
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
