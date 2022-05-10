package game

import (
	"errors"
	"fmt"
)

type Room struct {
	Description string   `json:"description"`
	Items       []Item   `json:"items"`
	Actions     []Action `json:"actions"`
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
	}

	return Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(1)], ItemRegistry[Id(2)]},
		Actions:     actions,
	}
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
	fmt.Printf("Items: %v\n", r.Items)

	filteredActions := []Action{}
	for _, action := range r.Actions {
		if action.Do == CollectItemEvt && action.To == id {
			continue
		} else {
			filteredActions = append(filteredActions, action)
		}
	}

	r.Actions = filteredActions
}

func enterRoom(state State, roomIds []Id) (State, *FlagSet, error) {
	if len(roomIds) != 1 {
		return state, nil, errors.New("No room ID specified.")
	}

	state.CurrentRoom = RoomRegistry[roomIds[0]]
	return state, nil, nil
}
