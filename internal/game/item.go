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
}

func collectItem(state State, itemIds []Id) (State, *FlagSet, error) {
	for _, itemId := range itemIds {
		state.PlayerState.Inventory[itemId] = ItemRegistry[itemId]
	}

	return state, nil, nil
}
