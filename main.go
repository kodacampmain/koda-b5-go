package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/kodacampmain/koda-b5-go/internals/counter"
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
	// defer fmt.Println("Bye Bye")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Printf("Recovery from Panic: %v\n", r)
	// 	}
	// }()
	// MyFunc()

	// str := "Hello World"
	// fmt.Printf("Value: %s\nReference: %v\n", str, &str)
	// secondStr := str
	// var strPointer *string = &str
	// str = "World Hello"
	// fmt.Printf("Pointer: %v\nValue: %s\n", strPointer, *strPointer)

	// bulbasaur := pokemons.NewPokemon("bulbasaur", "bulbasaur.jpg", []string{"grass", "poison"}, []pokemons.Abilities{
	// 	{
	// 		Name:     "bulbasaur",
	// 		IsHidden: false,
	// 	},
	// 	{
	// 		Name:     "chlorophyll",
	// 		IsHidden: true,
	// 	},
	// })
	// fmt.Println(bulbasaur.GetPokemonNameWithType())
	// bulbasaur.UpdatePokemonImage("bulbasaur.png")
	// fmt.Println(bulbasaur.GetPokemonImage())

	// dog := animals.Dog{}
	// cat := animals.Cat{}
	// animals.GetAnimalColor(dog)
	// animals.GetAnimalColor(cat)
	// var HP = 100
	// fmt.Println("Get Attacked by the DOG")
	// fmt.Printf("Hp calculation:\n%d%d = %d\n", HP, dog.Attack(), HP+dog.Attack())
	// animals.GetAnimalSound(cat)
	// animals.GetAnimalSound(dog)
	// Panicable()
	// os.Exit(0)

	// var wg sync.WaitGroup

	// helloChan := make(chan string, 4)

	// wg.Add(1)
	// go GoHello(&wg, helloChan)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for range 5 {
	// 		time.Sleep(time.Millisecond * 100)
	// 		fmt.Println("Halo")
	// 	}
	// }()
	// wg.Go(GoHalo)
	// wg.Go(func() {
	// 	defer func() {
	// 		if r := recover(); r != nil {
	// 			fmt.Println(r)
	// 		}
	// 	}()
	// 	for i := range 5 {
	// 		fmt.Println(i)
	// 		if i == 3 {
	// 			panic("waduh")
	// 		}
	// 	}
	// })

	// for {
	// 	time.Sleep(time.Millisecond * 500)
	// 	strHello, ok := <-helloChan
	// 	if !ok {
	// 		break
	// 	}
	// 	// Simulasi ambil data dari db
	// 	fmt.Print("Dari Channel: ")
	// 	fmt.Println(strHello)

	// 	strHello2, ok := <-helloChan
	// 	if !ok {
	// 		break
	// 	}
	// 	// Simulasi ambil data dari db
	// 	// time.Sleep(time.Millisecond * 500)
	// 	fmt.Print("Dari Channel: ")
	// 	fmt.Println(strHello2)
	// }
	// wg.Wait()

	counter.RunCounter()
}

func MyFunc() {
	// FILO => FIRST in LAST out
	defer fmt.Println("Good Bye")
	defer fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	// createDatabaseConnection
	// defer closeDatabaseConnection
	// jalankan query untuk ambil data
}

func Panicable() {
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

func GoHello(wg *sync.WaitGroup, c chan string) {
	defer func() {
		fmt.Println("Pengiriman Selesai")
		close(c)
		// wg.Done()
	}()
	str := []string{"Virgil", "Rohman", "Bian", "Ghifar", "Ari"}
	for _, v := range str {
		// time.Sleep(time.Millisecond * 100)
		// fmt.Println("Hello")
		// randNumber := rand.Intn(4)
		c <- v
	}
}

func GoHalo() {
	for range 5 {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Halo")
	}
}

type Message struct{}

func NewMessage() Message {
	return Message{}
}
func Blackboard(mc chan Message)             {}
func MsgSender(mc chan Message, msg Message) {}
