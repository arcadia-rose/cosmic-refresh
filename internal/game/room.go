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
	Id          Id              `json:"id"`
}

func RoomRegistry(state State) map[Id]Room {
	return map[Id]Room{
		Id(1000): MainEntrance(),
		Id(1001): ShoeRoom(state),
		Id(1002): DarkRoom(state),
		Id(1003): LockedRoom(state),
		Id(1004): UnlockedRoom(state),
		Id(1005): Parlour(state),
		Id(1006): Study(state),
		Id(1007): SecretRoom(state),
		Id(1008): Workshop(state),
		Id(1009): GameOver(state),
	}
}

func MainEntrance() Room {
	description := `The main entrance. Nowhere to go but in.`

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
		Id:          Id(1000),
	}
}

func ShoeRoom(state State) Room {
	description := `The entrance is surprisingly spaceous.  To your right is a large rack for
  leaving one's shoes.  It's unlikely that a single living soul has been through here in some time
  and the shoes left on the rack are tattered and falling to pieces.

	There's a small door to the left, leading to some sort of small parlour.
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
			It: "Enter small door",
			To: Id(1005),
			Is: RoomE,
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
		Id:          Id(1001),
	}.Prepare(state)
}

func Parlour(state State) Room {
	description := `Some kind of parlour. Altogether a little too homey for this place.
	A chair sits against a wall, with a small end table next to it. The wall behind it looks unstable - you wouldn't like to sit there too long.
	A few books are strewn on the end table, left there by someone who isn't terribly careful with their possessions when they're done with them.
	
	To your right is a beautiful, large wooden door that has been left completely ajar.  You can see that it leads to a study.`

	room := Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(5)]},
		Actions: []Action{
			{
				Do: EnterRoomEvt,
				It: "Exit",
				To: Id(1001),
				Is: RoomE,
			},
			{
				Do: CollectItemEvt,
				It: "Take book",
				To: Id(5),
				Is: ItemE,
			},
			{
				Do: EnterRoomEvt,
				It: "Enter study",
				To: Id(1006),
				Is: RoomE,
			},
		},
		Properties: map[string]bool{},
		Id:         Id(1005),
	}

	// Flag set via reading the page that suggests
	// this room had one more exit.
	if state.Flags[Id(2003)].Set {
		room.Actions = append(room.Actions, Action{
			Do: EnterRoomEvt,
			It: "Squeeze through the crumbling wall",
			To: Id(1007),
			Is: RoomE,
		})
	}

	return room.Prepare(state)
}

func Study(state State) Room {
	description := `A study with a surprising air of grandiosity in spite of its relatively small size.
	It has an eclectic quality about it with various trinkets and knick-knacks on the shelves and desk.
	
	Approaching the desk, you realize that the room must have been in use fairly recently.
	There's a strange, worn-out book with a detailed leather binding on the desk and what you can only
	describe as a pattern resembling a snake in the shape of the letter Z on the cover.
	
	Next to the book is a beautiful magnifying glass.  With a wooden handle that must have been hand-carved,
	there isn't anything about it that would immediately make you think it any more than a standard magnifying
	glass.  Somehow, though, you find it difficult to avert your gaze from the glimmer of the lens.`

	return Room{
		Description: description,
		Items: []Item{
			ItemRegistry[Id(4)],
			ItemRegistry[Id(7)],
		},
		Actions: []Action{
			{
				Do: CollectItemEvt,
				It: "Take book",
				To: Id(4),
				Is: ItemE,
			},
			{
				Do: CollectItemEvt,
				It: "Take magnifying glass",
				To: Id(7),
				Is: ItemE,
			},
			{
				Do: EnterRoomEvt,
				It: "Exit",
				To: Id(1005),
				Is: RoomE,
			},
		},
		Properties: map[string]bool{},
		Id:         Id(1006),
	}.Prepare(state)
}

func DarkRoom(state State) Room {
	description := `A dark room. Good thing you brought that candlestick.`

	room := Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(5)]},
		Actions: []Action{
			{
				Do: EnterRoomEvt,
				It: "Enter left door",
				To: Id(1003),
				Is: RoomE,
			},
			{
				Do: EnterRoomEvt,
				It: "Enter door ahead",
				To: Id(1004),
				Is: RoomE,
			},
			{
				Do: EnterRoomEvt,
				It: "Exit",
				To: Id(1001),
				Is: RoomE,
			},
		},
		Properties: map[string]bool{
			"dark": true,
		},
		Id: Id(1002),
	}

	if !state.PlayerState.HasItem(Id(0)) {
		room.Actions = append(room.Actions, Action{
			Do: SearchEvt,
			It: "Search desk",
			To: Id(2001),
			Is: FlagE,
		})
	}

	return room.Prepare(state)
}

func LockedRoom(state State) Room {
	description := `Strangely ordinary-looking for a place that was hard to get into.
	A pile of books lies on the ground. Seems someone was in a hurry and didn't reshelve their books when they were done.
	Your inner librarian groans.`

	return Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(3)], ItemRegistry[Id(6)]},
		Actions: []Action{
			{
				Do: EnterRoomEvt,
				It: "Exit",
				To: Id(1002),
				Is: RoomE,
			},
			{
				Do: CollectItemEvt,
				It: "Take book with Y-shaped symbol",
				To: Id(3),
				Is: ItemE,
			},
			{
				Do: CollectItemEvt,
				It: "Take book with caduceus",
				To: Id(6),
				Is: ItemE,
			},
		},
		Properties: map[string]bool{
			"locked": true,
		},
		Id: Id(1003),
	}.Prepare(state)
}

func UnlockedRoom(state State) Room {
	description := `Just your ordinary empty room.
	On the far side of the wall is a bookshelf, mostly empty but featuring a scant two rows of books.
	The titles seem strangely prosaic: a gardening book, a dry-looking biology textbook (at least a century out of date).
	You can't imagine reading any of them.
	There are a few gaps between books; you don't understand the cataloguing system here, and it's not obvious to you what belongs in the empty spaces.
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
			"checkboxes": !state.Flags[Id(2002)].Set,
		},
		Id: Id(1004),
	}.Prepare(state)
}

func SecretRoom(state State) Room {
	description := `Surprisingly clean and tidy for a room that was locked behind a hastily-built wall. Seems like it wasn't actually sealed away all that long ago.
	One wall is lined with bookshelves, but the books themselves are missing.
	The front part of the room has a table covered with some equipment you don't understand terribly well, a pair of small vials, and a few sheets of paper. Looks like someone's been hard at work in here.
	The chair at the table looks a lot less comfortable than the one in the room you just came from. Seems like this was a room for work, not for rest.`

	actions := []Action{
		{
			Do: CollectItemEvt,
			It: "Take page",
			To: Id(9),
			Is: ItemE,
		},
	}

	if state.Flags[Id(2004)].Set {
		description += "\nYou notice a door at the far side of the room. Somehow, you failed to notice it the first time you arrived."

		actions = append(actions, Action{
			Do: EnterRoomEvt,
			It: "Enter far door",
			To: Id(1008),
			Is: RoomE,
		})
	}

	return Room{
		Description: description,
		Items:       []Item{ItemRegistry[Id(9)]},
		Actions:     actions,
		Properties:  map[string]bool{},
		Id:          Id(1007),
	}.Prepare(state)
}

func Workshop(state State) Room {
	// You feel unsettled by them in a way you can't explain.
	description := `A large, open clinical room. This feels like a huge area to have been hidden away like this.
	Several tables, each with a human body lying on them. It all seems orderly, methodical. Someone did this very deliberately.
	A shelf in the back is lined with jars containing some kind of biological sample, but it doesn't seem to be human in origin.
	On the only table without a body lies a thick notebook, with every single page covered in densely written notes. Seems to have been written by whoever was studying these bodies. There are notes on their condition, descriptions of some kind of affliction you don't recognize. It seems their suffering began long before whoever was studying them began their investigation.`

	actions := []Action{
		{
			Do: ExamineItemEvt,
			It: "Examine bodies",
			To: Id(11),
			Is: ItemE,
		},
		{
			Do: ExamineItemEvt,
			It: "Examine notebook",
			To: Id(12),
			Is: ItemE,
		},
		{
			Do: CollectItemEvt,
			It: "Take page",
			To: Id(10),
			Is: ItemE,
		},
	}

	if state.Flags[Id(2005)].Set { // flag indicating the player read the page in this room
		description += "\nOne wall in its entirety is filled with an enormous symbol, intricate and enormously complex. It must have taken days to draw."

		actions = append(actions, Action{
			Do: ExamineItemEvt,
			It: "Examine wall",
			To: Id(13),
			Is: ItemE,
		})
	}

	return Room{
		Description: description,
		Items:       []Item{},
		Actions:     actions,
		Properties:  map[string]bool{},
		Id:          Id(1008),
	}.Prepare(state)
}

// The ending sequence triggered after seeing the symbol
func GameOver(state State) Room {
	description := `You have to leave, now. No one should ever return to this place where they could make the mistake that took place in this room.
	
	You hurry out of the room, as quickly as your legs can take you, hands scrambling for support along the wall. You don't feel like you're moving very well. You're not sure if you care.
	
	Using whatever you have with you, you prop up the unstable wall again, leaving that place closed up behind you. You hope there's no return to this place.
	
	You put back everything you took, leaving these things like a little memento of this cursed place. You're not really sure who it's for.
	
	Finally, you stagger out the front door. You're gone. Your head is pounding, worse every passing moment, and what you remembered is already leaving you. You're not sure you can get any further on your own. You wait a moment, hoping to gather your strength and focus your mind again. You have to refresh.
	
	The main entrance.`

	return Room{
		Description: description,
		Items:       []Item{},
		Actions:     []Action{},
		Properties:  map[string]bool{},
		Id:          Id(1009),
	}.Prepare(state)
}

func (r *Room) AppendDescription(desc string) {
	r.Description += "\n" + desc
}

func (r *Room) AddAction(action Action) {
	r.Actions = append(r.Actions, action)
}

func (r Room) Prepare(state State) Room {
	for id, item := range state.PlayerState.Inventory {
		r.RemoveItem(item, id)
	}

	return r
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
	if targetRoom.Properties["locked"] && !state.PlayerState.HasItem(Id(0)) {
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
