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
	for _, itemId := range itemIds {
		item := ItemRegistry[itemId]
		state.PlayerState.Inventory[itemId] = item

		state.CurrentRoom.RemoveItem(item, itemId)
	}

	return state, nil, nil
}
