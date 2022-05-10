package game

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
			It: "Take",
			To: Id(1),
			Is: ItemE,
		},
	}

	return Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(1)]},
		Actions:     actions,
	}
}

func enterRoom(state State, roomIds []Id) State {
	if len(roomIds) != 1 {
		return state
	}

	state.CurrentRoom = RoomRegistry[roomIds[0]]
	return state
}
