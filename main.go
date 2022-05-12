package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	inputString, err := ReadInput()
	if err != nil {
		fmt.Println("failed reading input")
		panic(err)
	}

	input, err := ParseInput[Input](inputString)
	if err != nil {
		fmt.Println("failed parsing input JSON to struct")
		panic(err)
	}

	lang := input.Lang
	fmt.Println("Woof in", lang)
}

type Input struct {
	Lang string `json:"lang"`
}

///////////////////////////////////////////////////////////////////////////////
// The stuff below should go in the extensions library
///////////////////////////////////////////////////////////////////////////////

func GetInputArgs() *Input {
	var args Input
	args.Lang = "en"
	return &args
}

func ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	consecutiveNewlinesCount := 0
	chars := []rune{}

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			return "", nil
		}
		if char == '\n' {
			consecutiveNewlinesCount++
		} else {
			consecutiveNewlinesCount = 0
			chars = append(chars, char)
		}

		if consecutiveNewlinesCount == 2 {
			break
		}
	}

	inputString := string(chars)
	return inputString, nil
}

func ParseInput[T any](inputString string) (*T, error) {
	rawBytes := []byte(inputString)

	var input T
	err := json.Unmarshal(rawBytes, &input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}
