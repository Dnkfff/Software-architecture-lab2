package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/ddynikov/Software-architecture-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to evaluate")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	var input io.Reader = nil
	var output = os.Stdout

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	}

	if *inputFile != "" {
		f, err := os.Open(*inputFile)
		if err != nil {
			fmt.Println("Error")
		}
		input = f
	}

	if *outputFile != "" {
		f, err := os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error")
		}
		output = f
	}

	if input == nil {
		fmt.Println("No stdIn defined")
		return
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	err := handler.Compute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
