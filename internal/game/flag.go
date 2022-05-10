package game

type Flag struct {
	Description string
	Set         bool
}

var FlagRegistry = map[Id]*Flag{
	Id(2000): HaveLightSource(),
}

func HaveLightSource() *Flag {
	f := new(Flag)
	f.Description = "A light source is required to activate dark rooms."
	f.Set = false
	return f
}

func ObtainLightSource(state State, _ids []Id) State {
	state.Flags[Id(2000)].Set = true
	return state
}
