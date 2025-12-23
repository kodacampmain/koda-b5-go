package main

import (
	"errors"
	"fmt"
)

func main() {
	// log.Println("Hello World")
	// fmt.Println([]byte("Hello World"))
	// fmt.Println([]rune("ᄈ"))
	var myName string // Manifest
	age := 1          // Inference
	myName = "Koda"   // Assignment
	// fmt.Println(myName)
	// fmt.Println(age)
	// fmt.Println(reflect.TypeOf(age))
	DetermineOldOrYoung(age)
	// greeting, error := CreateGreetingWithNameAndAge(myName, age)
	var greeting string
	if greeting, error := CreateGreetingWithNameAndAge(myName, age); error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Printf("My Greeting\n%s\n", greeting)
	}
	fmt.Println(greeting)
}

func DetermineOldOrYoung(age int) {
	switch true {
	case age < 25 && age > 0:
		fmt.Println("Masih Muda")
	case age >= 25:
		fmt.Println("Sudah Tuwir")
	default:
		fmt.Println("Umur Aneh")
	}
}

func CreateGreetingWithNameAndAge(name string, age int) (string, error) {
	if age < 0 {
		return "", errors.New("age has to be equal or more than zero")
	}

	str := fmt.Sprintf("Hello, My Name is %s\nI am %d years old", name, age)
	return str, nil
}
