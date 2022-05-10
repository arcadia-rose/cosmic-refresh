package game

type Event string

type Action struct {
	Do Event  `json:"do"`
	It string `json:"it"`
	To Id     `json:"to"`
	Is Entity `json:"is"`
}

type FlagSet struct {
	FlagId   Id
	NewValue bool
}

type Handler func(State, []Id) (State, *FlagSet, error)

var EventsRegistry = map[Event]Handler{
	CollectItemEvt: collectItem,
	EnterRoomEvt:   enterRoom,
}
