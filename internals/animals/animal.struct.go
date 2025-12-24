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

type IAnimal interface {
	Sound() string
	GetColor() string
}

func GetAnimalSound(animal IAnimal) {
	fmt.Println(animal.Sound())
}

func GetDogSound() {
	dog := Dog{}
	fmt.Println(dog.Sound())
}
func GetCatSound() {
	cat := Cat{}
	fmt.Println(cat.Sound())
}

func GetAnimalColor(animal IAnimal) {
	fmt.Println(animal.GetColor())
}
