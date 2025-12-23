package loop

import "fmt"

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
