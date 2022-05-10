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

type Action struct {
	Do Event  `json:"do"`
	It string `json:"it"`
	To Id     `json:"to"`
	Is Entity `json:"is"`
}

type Room struct {
	Description string   `json:"description"`
	Items       []Item   `json:"items"`
	Actions     []Action `json:"actions"`
}

type Item struct {
	Name   string `json:"name"`
	Amount uint   `json:"amount"`
}

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

var ItemRegistry = map[Id]Item{
	Id(0): Item{
		Name:   "Key",
		Amount: 1,
	},
}

var RoomRegistry = map[Id]Room{
	Id(1000): MainEntrance(),
	Id(1001): ShoeRoom(),
}

func MainEntrance() Room {
	description := `You approach a big spooky door!`

	actions := []Action{
		{
			Do: EnterRoomEvt,
			It: "Enter",
			To: Id(1001),
			Is: RoomE,
		},
	}

	return Room{
		Description: description,
		Items:       []Item{},
		Actions:     actions,
	}
}

func ShoeRoom() Room {
	description := `The entrance is surprisingly spaceous.  To your right is a large rack for
  leaving one's shoes.  It's unlikely that a single living soul has been through here in some time
  and the shoes left on the rack are tattered and falling to pieces.`

	return Room{
		Description: description,
		Items:       []Item{},
		Actions:     []Action{},
	}
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

func collectItem(state State, itemIds []Id) State {
	for _, itemId := range itemIds {
		state.PlayerState.Inventory[itemId] = ItemRegistry[itemId]
	}

	return state
}

func enterRoom(state State, roomIds []Id) State {
	if len(roomIds) != 1 {
		return state
	}

	state.CurrentRoom = RoomRegistry[roomIds[0]]
	return state
}
