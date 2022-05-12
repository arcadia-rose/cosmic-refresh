package game

import "fmt"

type Id uint

type Entity uint

const (
	CollectItemEvt = Event("collectItem")
	EnterRoomEvt   = Event("enterRoom")
	SearchEvt      = Event("search")
	UseItemsEvt    = Event("useItems")
	ExamineItemEvt = Event("examineItem")
	OpenBoxEvt     = Event("openBox")
)

const (
	ItemE Entity = 1
	RoomE Entity = iota
	FlagE Entity = iota
)

type Player struct {
	Insight   uint        `json:"insight"`
	Inventory map[Id]Item `json:"inventory"`
}

func (p Player) HasItem(id Id) bool {
	_, found := p.Inventory[id]
	return found
}

// Internal representation of the game state only usably on the Go side.
type State struct {
	PlayerState   Player
	CurrentRoom   Room
	Notifications []string
	Flags         map[Id]*Flag
	EventHandlers map[Event]Handler
}

func (s State) View() StateView {
	return StateView{
		PlayerState:   s.PlayerState,
		CurrentRoom:   s.CurrentRoom,
		Notifications: s.Notifications,
	}
}

// Shared representation of the game state meant to be communicated between JS & Go.
// The non-serializable components of `State` must be reconstructed for internal use.
type StateView struct {
	PlayerState   Player   `json:"player"`
	CurrentRoom   Room     `json:"currentRoom"`
	Notifications []string `json:"notifications"`
}

func GameLoop(state State, event Event, ids []Id) State {
	state.Notifications = []string{}

	handler, found := state.EventHandlers[event]
	if found {
		newState, flagToSet, _ := handler(state, ids)

		if flagToSet != nil {
			newState.Flags[flagToSet.FlagId].Set = flagToSet.NewValue
			// Rerendering the room ensures that its properties
			// can change based on the new flags.
			newState.CurrentRoom = RoomRegistry(state)[newState.CurrentRoom.Id]
		}

		state = newState
	}

	fmt.Printf("Room has %d actions available\n", len(state.CurrentRoom.Actions))
	return state
}

func NewPlayer() Player {
	return Player{
		Insight:   0,
		Inventory: map[Id]Item{},
	}
}

func NewState() State {
	return State{
		PlayerState:   NewPlayer(),
		CurrentRoom:   MainEntrance(),
		Flags:         FlagRegistry,
		EventHandlers: EventsRegistry,
	}
}
