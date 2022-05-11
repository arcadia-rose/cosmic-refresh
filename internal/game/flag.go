package game

import "fmt"

type Flag struct {
	Description string
	Set         bool
}

var FlagRegistry = map[Id]*Flag{
	Id(2000): HaveLightSource(),
	Id(2001): ItemDiscoverable(),
	Id(2002): BookPuzzleSolved(),
	Id(2003): SecretRoomDiscovered(),
	Id(2004): ReadPage1(),
}

func HaveLightSource() *Flag {
	f := new(Flag)
	f.Description = "A light source is required to activate dark rooms."
	f.Set = false
	return f
}

func ItemDiscoverable() *Flag {
	f := new(Flag)
	f.Description = "Some items can only be revealed once the player has fulfilled some condition."
	f.Set = false
	return f
}

func BookPuzzleSolved() *Flag {
	f := new(Flag)
	f.Description = "The player has solved the book puzzle."
	f.Set = false
	return f
}

func SecretRoomDiscovered() *Flag {
	f := new(Flag)
	f.Description = "The player knows of the secret room behind the parlour"
	f.Set = false
	return f
}

func ReadPage1() *Flag {
	f := new(Flag)
	f.Description = "The player has read the page describing the ritual"
	f.Set = false
	return f
}

func ObtainLightSource(state State, _ids []Id) State {
	state.Flags[Id(2000)].Set = true
	return state
}

func ToggleItemDiscoverability(state State, ids []Id) State {
	fmt.Printf("Discovered %v\n", ids)
	if len(ids) != 1 {
		return state
	}

	state.Flags[ids[0]].Set = !state.Flags[ids[0]].Set
	return state
}
