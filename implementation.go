package main

import (
	"fmt"
	"regexp"
	"strings"
)

func stringIsValid(inputString string) bool {
	const regex = `[0-9/\*\^\+\-\(\)]`
	arrayFromString := strings.Split(inputString, "")

	var stringIsValid bool = true

	for _, val := range arrayFromString {
		matched, _ := regexp.MatchString(regex, string(val))
		if !matched {
			stringIsValid = false
		}
	}

	return stringIsValid
}

// TODO: document this function.
// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {

	const invalidError = "The string is invalid"

	stringIsValid := stringIsValid(input)

	if !stringIsValid {
		fmt.Println(invalidError)
		return "_", fmt.Errorf(invalidError)
	}

	return "TODO", fmt.Errorf("TODO")
}

func main() {
	PrefixToPostfix("a12312&^(**^%&")
}
