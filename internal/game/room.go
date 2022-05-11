package game

import (
	"errors"
	"fmt"
)

type Room struct {
	Description string          `json:"description"`
	Items       []Item          `json:"items"`
	Actions     []Action        `json:"actions"`
	Properties  map[string]bool `json:"properties"`
}

func RoomRegistry(state State) map[Id]Room {
	return map[Id]Room{
		Id(1000): MainEntrance(),
		Id(1001): ShoeRoom(),
		Id(1002): DarkRoom(state),
		Id(1003): LockedRoom(),
		Id(1004): UnlockedRoom(),
	}
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
		Properties:  map[string]bool{},
	}
}

func ShoeRoom() Room {
	description := `The entrance is surprisingly spaceous.  To your right is a large rack for
  leaving one's shoes.  It's unlikely that a single living soul has been through here in some time
  and the shoes left on the rack are tattered and falling to pieces.
  
  You can see there is a dark room to the right as you look down the hallway.`

	actions := []Action{
		{
			Do: CollectItemEvt,
			It: "Take shoe",
			To: Id(1),
			Is: ItemE,
		},
		{
			Do: CollectItemEvt,
			It: "Take candlestick",
			To: Id(2),
			Is: ItemE,
		},
		{
			Do: EnterRoomEvt,
			It: "Enter dark room",
			To: Id(1002),
			Is: RoomE,
		},
	}

	return Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(1)], ItemRegistry[Id(2)]},
		Actions:     actions,
		Properties:  map[string]bool{},
	}
}

func DarkRoom(state State) Room {
	description := `A dark room. Good thing you brought that candlestick.`

	room := Room{
		Description: description,
		Items:       []Item{},
		Actions: []Action{
			{
				Do: EnterRoomEvt,
				It: "Enter left door",
				To: Id(1003),
				Is: RoomE,
			},
			{
				Do: EnterRoomEvt,
				It: "Enter right door",
				To: Id(1004),
				Is: RoomE,
			},
		},
		Properties: map[string]bool{
			"dark": true,
		},
	}

	if !state.PlayerState.HasKey() {
		room.Actions = append(room.Actions, Action{
			Do: SearchEvt,
			It: "Search desk",
			To: Id(2001),
			Is: FlagE,
		})
	}

	return room
}

func LockedRoom() Room {
	description := `Strangely ordinary-looking for a place that was hard to get into.`

	return Room{
		Description: description,
		Items:       []Item{},
		Actions: []Action{
			{
				Do: EnterRoomEvt,
				It: "Exit",
				To: Id(1002),
				Is: RoomE,
			},
		},
		Properties: map[string]bool{
			"locked": true,
		},
	}
}

func UnlockedRoom() Room {
	description := `Just your ordinary empty room.
	On the far side of the wall is a bookshelf, mostly empty but featuring a scant two rows of books.
	The titles seem strangely prosaic: a gardening book, a dry-looking biology textbook (at least a century out of date).
	You can't imagine reading any of them.
	Strangely, the spines are all facing away from you; you can't seem to make sense of why.
	Beneath the books is a small wooden box, nondescript, with a tightly-fitted lid. There's no obvious lock, but it won't open.`

	return Room{
		Description: description,
		Items:       []Item{},
		Actions: []Action{
			{
				Do: EnterRoomEvt,
				It: "Exit",
				To: Id(1002),
				Is: RoomE,
			},
		},
		Properties: map[string]bool{
			"checkboxes": true,
		},
	}
}

func (r *Room) AppendDescription(desc string) {
	r.Description += "\n" + desc
}

func (r *Room) AddAction(action Action) {
	r.Actions = append(r.Actions, action)
}

func (r *Room) RemoveAction(event Event, target Id) {
	filteredActions := []Action{}
	for _, action := range r.Actions {
		if action.Do == event && action.To == target {
			fmt.Printf("Removing %s targetting %d\n", event, target)
			continue
		} else {
			filteredActions = append(filteredActions, action)
		}
	}

	r.Actions = filteredActions
	fmt.Printf("%d actions left", len(r.Actions))
}

func (r *Room) RemoveItem(item Item, id Id) {
	// Produce a filtered copy of the inventory with the
	// fetched item removed.
	invCopy := []Item{}
	for _, roomItem := range r.Items {
		if item == roomItem {
			continue
		} else {
			invCopy = append(invCopy, roomItem)
		}
	}

	r.Items = invCopy

	r.RemoveAction(CollectItemEvt, id)
}

func enterRoom(state State, roomIds []Id) (State, *FlagSet, error) {
	if len(roomIds) != 1 {
		return state, nil, errors.New("No room ID specified.")
	}

	targetRoom := RoomRegistry(state)[roomIds[0]]
	if targetRoom.Properties["dark"] && !FlagRegistry[Id(2000)].Set {
		state.Notifications = append(state.Notifications, "The room is too dark.  What would you even do when you got in there?")
		return state, nil, nil
	}
	if targetRoom.Properties["locked"] && !state.PlayerState.HasKey() {
		state.Notifications = append(state.Notifications, "The room is locked.  You need a key to get in.")
		return state, nil, nil
	}

	state.CurrentRoom = RoomRegistry(state)[roomIds[0]]
	return state, nil, nil
}

func search(state State, ids []Id) (State, *FlagSet, error) {
	fmt.Printf("Searching %v\n", ids)

	state = ToggleItemDiscoverability(state, ids)

	for _, id := range ids {
		state.CurrentRoom.RemoveAction(SearchEvt, id)
	}

	state.CurrentRoom.AppendDescription(`You search the desk and find a key.`)
	state.CurrentRoom.AddAction(Action{
		Do: CollectItemEvt,
		It: "Take key",
		To: Id(0),
		Is: ItemE,
	})

	return state, nil, nil
}
