package conditionals

import "fmt"

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
