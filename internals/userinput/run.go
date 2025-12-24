package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func Init() {
	defer fmt.Println("Selamat Tinggal")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Masukkan Pilihan: ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		if input == "1" {
			fmt.Print("Masukkan Perintah: ")
			scanner.Scan()
			input := scanner.Text()
			fmt.Printf("Perintah diterima <=== %s\n", input)
		}
		fmt.Printf("Input => %s\n", input)
	}
}
