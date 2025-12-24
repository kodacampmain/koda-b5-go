package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// log.Println("Hello World")
	// fmt.Println([]byte("Hello World"))
	// fmt.Println([]rune("ᄈ"))
	// var myName string // Manifest
	// age := -1         // Inference
	// myName = "Koda"   // Assignment
	// fmt.Println(myName)
	// fmt.Println(age)
	// fmt.Println(reflect.TypeOf(age))
	// cdn.DetermineOldOrYoung(age)
	// greeting, error := CreateGreetingWithNameAndAge(myName, age)
	// var greeting string
	// if greeting, error := cdn.CreateGreetingWithNameAndAge(myName, age); error != nil {
	// 	fmt.Println(error.Error())
	// } else {
	// 	fmt.Printf("My Greeting\n%s\n", greeting)
	// }
	// fmt.Println(greeting)
	// konversi tipe data = tipe(variabel)
	// fmt.Println(float32(age), reflect.TypeOf(float32(age)))
	// var ppn float32 = 0.11
	// ppnStr := strconv.Itoa(int(ppn))
	// fmt.Println(ppn, ppnStr)
	// fmt.Println(string(ppn))
	// l.PrintEvenNumber(10)
	// arrayslice.ArrayAndSlice()
	// pokemons.GetPokemon()
	// userinput.Init()
	defer fmt.Println("Bye Bye")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovery from Panic: %v\n", r)
		}
	}()
	myFunc()
	panicable()
	os.Exit(0)
}

func myFunc() {
	// FILO => FIRST in LAST out
	defer fmt.Println("Good Bye")
	defer fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	// createDatabaseConnection
	// defer closeDatabaseConnection
	// jalankan query untuk ambil data
}

func panicable() {
	defer fmt.Println("Selamat Tinggal")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovery from Panic: %v\n", r)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if input == "panic" {
		panic("terjadi panic")
		// fmt.Println("Waduh")
	}
}
