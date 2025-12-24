package pokemons

import (
	"fmt"
	"strings"
)

type Abilities struct {
	Name     string
	IsHidden bool
}

type gender string

const (
	female gender = "female"
	male   gender = "male"
)

type pokemon struct {
	name      string
	img       string
	types     []string
	abilities []Abilities
	gender    gender
	// timestamp time.Time
}

// Constructor Function
func NewPokemon(name, img string, types []string, abilities []Abilities) *pokemon {
	return &pokemon{
		name:      name,
		img:       img,
		types:     types,
		abilities: abilities,
		gender:    female,
		// timestamp: time.Now(),
	}
}

func (p *pokemon) GetPokemonNameWithType() string {
	return fmt.Sprintf("%s\n%s", p.name, strings.Join(p.types, ", "))
}

func (p *pokemon) GetPokemonImage() string {
	return p.img
}

func (p *pokemon) UpdatePokemonImage(NewImg string) {
	p.img = NewImg
}
