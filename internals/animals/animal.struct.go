package animals

import "fmt"

type Dog struct{}

func (d Dog) Sound() string {
	return "Woof"
}
func (d Dog) GetColor() string {
	return "Brown"
}
func (d Dog) Attack() int {
	return -20
}

type Cat struct{}

func (c Cat) Sound() string {
	return "Nyaa"
}
func (c Cat) GetColor() string {
	return "Yellow"
}

type Animal interface {
	Sound() string
	GetColor() string
}

func GetAnimalSound(animal Animal) {
	fmt.Println(animal.Sound())
}

func GetAnimalColor(animal Animal) {
	fmt.Println(animal.GetColor())
}
