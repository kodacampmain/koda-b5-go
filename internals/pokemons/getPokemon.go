package pokemons

import "fmt"

func GetPokemon() {
	// var pokemonTypes map[string]string
	pokemonTypes := make(map[string]string)
	fmt.Println(pokemonTypes)
	pokemonTypes["Electric"] = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/types/generation-viii/sword-shield/13.png"
	pokemonTypes["Grass"] = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/types/generation-viii/sword-shield/12.png"
	pokemonTypes["Psychic"] = "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/types/generation-viii/sword-shield/14.png"
	// fmt.Println(pokemonTypes)
	types := []string{"Psychic", "Psychic", "Electric"}

	for _, v := range types {
		fmt.Println(pokemonTypes[v])
	}

	electrode := pokemon{
		name:  "Electrode",
		img:   "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/shiny/101.png",
		types: []string{"Electric"},
	}

	// exeggcute := pokemon{
	// 	name: "Exeggcute",
	// 	img:  "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/shiny/102.png",
	// 	types: []string{
	// 		"Grass",
	// 		"Psychic",
	// 	},
	// }

	fmt.Printf("Name: %s\nTypes: %s\n", electrode.name, electrode.types[0])
	// fmt.Printf("Name: %s\nTypes: %v\n", electrode.name, electrode.types)
}
