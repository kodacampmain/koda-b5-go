package pokemons

type pokemon struct {
	name      string
	img       string
	types     []string
	abilities []abilities
}

type abilities struct {
	name     string
	isHidden bool
}
