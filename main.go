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
	age := -1         // Inference
	myName = "Koda"   // Assignment
	// fmt.Println(myName)
	// fmt.Println(age)
	// fmt.Println(reflect.TypeOf(age))
	// DetermineOldOrYoung(age)
	// greeting, error := CreateGreetingWithNameAndAge(myName, age)
	// var greeting string
	if greeting, error := CreateGreetingWithNameAndAge(myName, age); error != nil {
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
	PrintEvenNumber(10)
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

func PrintEvenNumber(limit int) {
	// for i := 0; i <= limit; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// i := 0
	// for i <= limit {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	i++
	// }

	for i := range limit {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
