package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var res int64 = 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineText := scanner.Text()
		combinedNumber, err := combineFirstAndLastNumberInString(lineText)
		if err != nil {
			fmt.Printf("Error combining first and last number in string %s: %v\n", lineText, err)
			return
		}
		res += int64(combinedNumber)
	}
	fmt.Printf("Result: %d\n", res)
}

func combineFirstAndLastNumberInString(s string) (int32, error) {
	var firstNumber, lastNumber rune
	var foundFirst, foundLast bool

	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			firstNumber = runes[i]
			foundFirst = true
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			lastNumber = runes[i]
			foundLast = true
			break
		}
	}

	if !foundFirst && !foundLast {
		return 0, fmt.Errorf("could not find first and last number in string")
	} else if foundFirst && !foundLast {
		resString := string(firstNumber) + string(firstNumber)
		res, err := strconv.Atoi(resString)
		if err != nil {
			return 0, err
		}
		return int32(res), nil
	}

	combinedNumberString := string(firstNumber) + string(lastNumber)
	combinedNumber, err := strconv.Atoi(combinedNumberString)
	if err != nil {
		return 0, err
	}

	return int32(combinedNumber), nil

}
