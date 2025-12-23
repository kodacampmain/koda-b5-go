package main

import (
	"fmt"

	arrayslice "github.com/kodacampmain/koda-b5-go/internals/arraySlice"
	cdn "github.com/kodacampmain/koda-b5-go/internals/conditionals"
	l "github.com/kodacampmain/koda-b5-go/internals/loop"
	"github.com/kodacampmain/koda-b5-go/internals/pokemons"
)

func main() {
	// log.Println("Hello World")
	// fmt.Println([]byte("Hello World"))
	// fmt.Println([]rune("ᄈ"))
	var myName string // Manifest
	age := -1         // Inference
	myName = "Koda"   // Assignment
	// fmt.Println(myName)
	// fmt.Println(age)
	// fmt.Println(reflect.TypeOf(age))
	cdn.DetermineOldOrYoung(age)
	// greeting, error := CreateGreetingWithNameAndAge(myName, age)
	// var greeting string
	if greeting, error := cdn.CreateGreetingWithNameAndAge(myName, age); error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("My Greeting\n%s\n", greeting)
	}
	// fmt.Println(greeting)
	// konversi tipe data = tipe(variabel)
	// fmt.Println(float32(age), reflect.TypeOf(float32(age)))
	// var ppn float32 = 0.11
	// ppnStr := strconv.Itoa(int(ppn))
	// fmt.Println(ppn, ppnStr)
	// fmt.Println(string(ppn))
	l.PrintEvenNumber(10)
	arrayslice.ArrayAndSlice()

	pokemons.GetPokemon()
}
