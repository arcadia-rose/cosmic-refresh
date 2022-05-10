package game

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
