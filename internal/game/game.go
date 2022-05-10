package game

type Id uint

type Event string

type Entity uint

const (
	CollectItemEvt = Event("collectItem")
	EnterRoomEvt   = Event("enterRoom")
)

const (
	ItemE Entity = 1
	RoomE Entity = iota
)

type Player struct {
	Insight   uint        `json:"insight"`
	Inventory map[Id]Item `json:"inventory"`
}

// Internal representation of the game state only usably on the Go side.
type State struct {
	PlayerState   Player
	CurrentRoom   Room
	EventHandlers map[Event]func(State, []Id) State
}

func (s State) View() StateView {
	return StateView{
		PlayerState: s.PlayerState,
		CurrentRoom: s.CurrentRoom,
	}
}

// Shared representation of the game state meant to be communicated between JS & Go.
// The non-serializable components of `State` must be reconstructed for internal use.
type StateView struct {
	PlayerState Player `json:"player"`
	CurrentRoom Room   `json:"currentRoom"`
}

func (s StateView) State() State {
	newState := NewState()
	newState.PlayerState = s.PlayerState
	newState.CurrentRoom = s.CurrentRoom
	return newState
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
		Inventory: map[Id]Item{},
	}
}

func NewState() State {
	return State{
		PlayerState: NewPlayer(),
		CurrentRoom: MainEntrance(),
		EventHandlers: map[Event]func(State, []Id) State{
			CollectItemEvt: collectItem,
			EnterRoomEvt:   enterRoom,
		},
	}
}
