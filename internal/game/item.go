package game

import "fmt"

type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var ItemRegistry = map[Id]Item{
	Id(0): {
		Name:        "Key",
		Description: "A simple iron key. You're not sure what door this opens.",
	},
	Id(1): {
		Name:        "Shoe",
		Description: "Disgusting. Don't put this on.",
	},
	Id(2): {
		Name:        "Candlestick",
		Description: `A candlestick. The candle has burned most of the way down, but it's still lit for now. You wonder, why was it lit in the first place?`,
	},
	Id(3): {
		Name:        "Book with Y-shaped symbol",
		Description: `An old-looking book with a strange, yellow Y-shaped symbol on the cover.`,
	},
	Id(4): {
		Name:        "Book with Z-shaped symbol",
		Description: `An old-looking book with a snakelike Z-shaped symbol on the cover.`,
	},
	Id(5): {
		Name:        "Book with squid-shaped symbol",
		Description: `An old-looking book with some kind of squid on the cover. You're not sure this is a biology textbook, though.`,
	},
	Id(6): {
		Name:        "Book with a caduceus",
		Description: `Some sort of book of folk remedies. Looks like several pages have been torn out.`,
	},
	Id(7): {
		Name:        "Magnifying glass",
		Description: `A normal-looking magnifying glass with a hand-carved wooden handle.`,
	},
}

func collectItem(state State, itemIds []Id) (State, *FlagSet, error) {
	var flagSet *FlagSet = nil

	for _, itemId := range itemIds {
		item := ItemRegistry[itemId]
		state.PlayerState.Inventory[itemId] = item

		if itemId == Id(2) { // lol
			flagSet = &FlagSet{
				FlagId:   Id(2000),
				NewValue: true,
			}
		}

		state.CurrentRoom.RemoveItem(item, itemId)
	}

	return state, flagSet, nil
}

func useItems(state State, itemIds []Id) (State, *FlagSet, error) {
	if len(itemIds) == 0 {
		state.Notifications = append(state.Notifications, "How are you planning to use nothing?")
	} else if len(itemIds) == 1 {
		state.Notifications = append(state.Notifications, "You're not really sure how to use that.")
	} else {
		state.Notifications = append(state.Notifications, "You're not really sure how to use those together.")
	}

	return state, nil, nil
}

func openBox(state State, itemIds []Id) (State, *FlagSet, error) {
	fmt.Printf("Passed in values: %v\n", itemIds)
	state.Notifications = append(state.Notifications, "You fiddle with the box but the lid stays tightly attached.")

	return state, nil, nil
}
