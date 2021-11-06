package main

import (
	"fmt"
	"regexp"
	"strings"
)

func stringIsValid(inputString string) bool {
	const regex = `[0-9/\*\^\+\-\(\)]`
	trimString := strings.TrimSpace(inputString)
	arrayFromString := strings.Split(trimString, "")

	var stringIsValid bool = true

	for _, val := range arrayFromString {
		matched, _ := regexp.MatchString(regex, string(val))
		if !matched {
			stringIsValid = false
		}
	}

	return stringIsValid
}

// check if string is number
func stringIsNumber(input string) bool {
	const regex = `[0-9]`
	matched, _ := regexp.MatchString(regex, string(input))

	return matched
}

// check if string is operator
func stringIsOperator(input string) bool {
	const regex = `[\+\-\/\*\^]`

	matched, _ := regexp.MatchString(regex, string(input))

	return matched
}

// translate string to special array
func arrayFromString(input string) []string {
	trimString := strings.TrimSpace(input)
	sampleArray := strings.Split(trimString, "")

	var result []string

	var temporary string = ""
	for _, val := range sampleArray {
		if stringIsNumber(val) {
			temporary = temporary + string(val)
		} else if temporary != "" {
			result = append(result, temporary)
			result = append(result, val)
			temporary = ""
		} else {
			result = append(result, val)
		}
	}

	if temporary != "" {
		result = append(result, temporary)
	}

	return result
}

func ExpressionToPostfix(input string) (string, error) {

	// check if string is correct
	stringIsValid := stringIsValid(input)
	if !stringIsValid {
		return "_", fmt.Errorf("The string is invalid")
	}

	// algoritm start
	var stack []string
	var queue []string

	array := arrayFromString(input)

	fmt.Println(array)

	for _, val := range array {
		if stringIsNumber(val) {
			queue = append(queue, val)
			// if element is number
		} else if stringIsOperator(val) && (len(stack) == 0 || stack[len(stack)-1] == "(") {
			stack = append(stack, val)
			// if element equals operator +-/*
			// & stack is empty || stack have the last element - "("
		}

		// if input operator is higher than last in stack
		// else if isHigherOperator()
	}

	fmt.Println(queue)
	fmt.Println(stack)

	return "TODO", fmt.Errorf("TODO")
}

func main() {
	ExpressionToPostfix("123-123+(322/122)/5")
}
