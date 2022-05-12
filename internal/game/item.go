package game

import (
	"fmt"
	"reflect"
)

var CheckboxSolution = []Id{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}

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
	Id(8): {
		Name:        "Page 3 of caduceus book",
		Description: `Describes medicines, elixirs and illnesses with strange symptoms...`,
	},
	Id(9): {
		Name:        "Page 1 of caduceus book",
		Description: `Filled with complex diagrams. Next to a basic biological study of human anatomy, you see a strangely occult symbol you don't recognize.`,
	},
	Id(10): {
		Name:        "Page 2 of caduceus book",
		Description: `An anatomical study of a squid. At least, you think it's a squid, but some of these body parts look off to you...`,
	},
	Id(11): {
		Name:        "Bodies",
		Description: `No sign of foul play, as far as you can tell. These look like natural deaths.`,
	},
	Id(12): {
		Name:        "Notebook",
		Description: `You don't recognize it? But after all, you're the one who wrote it.`,
	},
	Id(13): {
		Name: "Wall",
		Description: `Something feels familiar about this symbol. Familiar, and wrong. Your head sears, a sharp pain that feels like it could split your skull in two. You have a memory of having been in this place, once.
		
		You need to get out of here. Something horrible happened, and you weren't meant to come back here. You can't let it be repeated, no matter what.`,
	},
}

var ItemInteractions = []func([]Id) (string, *FlagSet){
	MagnifyYellowBook,
	MagnifyCaduceusBook,
	MagnifyCaduceusPage3,
	MagnifyCaduceusPage1,
	MagnifyCaduceusPage2,
}

func MagnifyYellowBook(items []Id) (string, *FlagSet) {
	foundMag := false
	foundBook := false

	for _, id := range items {
		if id == Id(7) {
			foundMag = true
		}
		if id == Id(3) {
			foundBook = true
		}
	}

	if foundMag && foundBook {
		return `That's funny. You notice a note scrawled in the margins that you missed before. It reads, "Please reshelve the books when you're finished. Remember, the yellow sign is the shelving guide..."`, nil
	}

	return "", nil
}

func MagnifyCaduceusBook(items []Id) (string, *FlagSet) {
	foundMag := false
	foundBook := false

	for _, id := range items {
		if id == Id(7) {
			foundMag = true
		}
		if id == Id(6) {
			foundBook = true
		}
	}

	if foundMag && foundBook {
		return `The words in the book appear larger, clearer, sharper... more profound.`, nil
	}

	return "", nil
}

func MagnifyCaduceusPage3(items []Id) (string, *FlagSet) {
	foundMag := false
	foundPage := false

	for _, id := range items {
		if id == Id(7) {
			foundMag = true
		}
		if id == Id(8) {
			foundPage = true
		}
	}

	if foundMag && foundPage {
		flagSet := &FlagSet{
			FlagId:   Id(2003),
			NewValue: true,
		}
		return `New words appear on the page that you can't perceive without it, hastily scribbled by someone who seemed to be in a great hurry. It says, "What I sealed up behind that parlour needs to stay forgotten. Whatever you do, don't let me go back there."`, flagSet
	}

	return "", nil
}

func MagnifyCaduceusPage1(items []Id) (string, *FlagSet) {
	foundMag := false
	foundPage := false

	for _, id := range items {
		if id == Id(7) {
			foundMag = true
		}
		if id == Id(9) {
			foundPage = true
		}
	}

	if foundMag && foundPage {
		flagSet := &FlagSet{
			FlagId:   Id(2004),
			NewValue: true,
		}
		return `The entire page fills with tightly-packed words documenting the process of finding and communicating with - something. You're not sure what. It describes moonlit rituals, strange contortions that you struggle to picture in your mind's eye. The notes are detailed, as though the writer had already tried many of them.`, flagSet
	}

	return "", nil
}

func MagnifyCaduceusPage2(items []Id) (string, *FlagSet) {
	foundMag := false
	foundPage := false

	for _, id := range items {
		if id == Id(7) {
			foundMag = true
		}
		if id == Id(10) {
			foundPage = true
		}
	}

	if foundMag && foundPage {
		flagSet := &FlagSet{
			FlagId:   Id(2005),
			NewValue: true,
		}
		return `What looked before like an anatomical diagram now seems more like the template for a contract, with the diagram itself forming one party's signature. The rest of the contract is already filled out, and a neat hand has signed the other side of the form with your name.`, flagSet
	}

	return "", nil
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
	foundUse := false

	var flagSet *FlagSet
	for _, interaction := range ItemInteractions {
		result, flagFound := interaction(itemIds)
		if flagFound != nil {
			flagSet = flagFound
		}

		if result != "" {
			state.Notifications = append(state.Notifications, result)
			foundUse = true
		}
	}

	if len(itemIds) == 0 {
		state.Notifications = append(state.Notifications, "How are you planning to use nothing?")
	} else if len(itemIds) == 1 && !foundUse {
		state.Notifications = append(state.Notifications, "You're not really sure how to use that.")
	} else if !foundUse {
		state.Notifications = append(state.Notifications, "You're not really sure how to use those together.")
	}

	return state, flagSet, nil
}

func examineItem(state State, itemIds []Id) (State, *FlagSet, error) {
	for _, itemId := range itemIds {
		item, ok := ItemRegistry[itemId]
		if !ok {
			continue
		}

		state.Notifications = append(state.Notifications, item.Description)

		// End the game here
		if itemId == Id(13) {
			state.CurrentRoom = GameOver(state)
			state.PlayerState.Inventory = make(map[Id]Item)
		}
	}
	return state, nil, nil
}

func countChecks(checkboxes []Id) int {
	count := 0
	for _, checkbox := range checkboxes {
		if checkbox == Id(1) {
			count++
		}
	}

	return count
}

func openBox(state State, flags []Id) (State, *FlagSet, error) {
	fmt.Printf("Passed in values: %v\n", flags)

	if countChecks(flags) != 4 {
		state.Notifications = append(
			state.Notifications,
			`You somehow get the feeling you need to place all four books...
			Maybe one of them has a clue about how to arrange them.`,
		)
		return state, nil, nil
	}

	var flagSet *FlagSet = nil

	solutionCorrect := reflect.DeepEqual(flags, CheckboxSolution)

	_, haveBook1 := state.PlayerState.Inventory[Id(3)]
	_, haveBook2 := state.PlayerState.Inventory[Id(4)]
	_, haveBook3 := state.PlayerState.Inventory[Id(5)]
	_, haveBook4 := state.PlayerState.Inventory[Id(6)]

	if haveBook1 && haveBook2 && haveBook3 && haveBook4 && solutionCorrect {
		state.Notifications = append(
			state.Notifications,
			`The lid of the box pops open with a satisfying clink. Inside, you find a page that appears to have been torn from a book.
			Upon further inspection, you find the page describes medicines, elixirs and illnesses with strange symptoms...`,
		)

		flagSet = &FlagSet{
			FlagId:   Id(2002),
			NewValue: true,
		}

		state.PlayerState.Inventory[Id(8)] = ItemRegistry[Id(8)]
	} else {
		state.Notifications = append(state.Notifications, "You fiddle with the box but the lid stays tightly attached.")
	}

	return state, flagSet, nil
}
