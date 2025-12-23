package conditionals

import (
	"errors"
	"fmt"
)

func CreateGreetingWithNameAndAge(name string, age int) (string, error) {
	if age < 0 {
		return "", errors.New("age has to be equal or more than zero")
	}

	str := fmt.Sprintf("Hello, My Name is %s\nI am %d years old", name, age)
	return str, nil
}
