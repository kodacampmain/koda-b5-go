package arrayslice

import "fmt"

func ArrayAndSlice() {
	var hobbies = [5]string{"Reading", "Playing", "Fishing", "Sports", "Clubbing"}
	fmt.Println(hobbies)
	ages := [3]int{20, 21, 22}
	fmt.Println(ages)
	var fruits = []string{"Durian", "Semangka", "Melon"}
	fmt.Println(fruits)
	isOpen := []bool{false, true, false}
	fmt.Println(isOpen)

	myHobbies := hobbies[1:3]
	myHobbies = append(myHobbies, "Eating", "Driving")
	yourHobbies := make([]string, len(myHobbies))
	copy(yourHobbies, myHobbies)

	myHobbies[0] = "Gaming"
	hobbies[2] = "Swimming"
	yourHobbies = append(yourHobbies, "Swimming", "Mining")

	// fmt.Println("hobbies", hobbies)
	// fmt.Println("myHobbies", myHobbies)
	// fmt.Println("yourHobbies", yourHobbies)
	// for i, v := range yourHobbies {
	// 	fmt.Printf("Index %d: %s\n", i, v)
	// }
}
