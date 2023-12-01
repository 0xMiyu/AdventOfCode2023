package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var numberMap = map[string]string{
	"zero":  "0",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	// var res int64 = 0

	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Println("Error opening file: ", err)
	// 	return
	// }

	// defer file.Close()

	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	lineText := scanner.Text()
	// 	combinedNumber, err := getShitDone(lineText)
	// 	if err != nil {
	// 		fmt.Printf("Error combining first and last number in string %s: %v\n", lineText, err)
	// 		return
	// 	}
	// 	fmt.Printf("Result: %d\n", combinedNumber)
	// 	res += int64(combinedNumber)
	// }
	// fmt.Printf("Result: %d\n", res)

	// testString := "qxxkgqqcggbjc85"
	// res, err := getShitDone(testString)
	// if err != nil {
	// 	fmt.Printf("Error combining first and last number in string %s: %v\n", testString, err)
	// 	return
	// }
	// fmt.Printf("Result: %d\n", res)
	input, _ := os.ReadFile("input.txt")
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	calc := func(nums []string) (result int) {
		first := regexp.MustCompile(`(` + strings.Join(nums, "|") + `)`)
		last := regexp.MustCompile(`.*` + first.String())

		for _, s := range strings.Fields(string(input)) {
			result += 10 * (slices.Index(nums, first.FindStringSubmatch(s)[1])%9 + 1)
			result += slices.Index(nums, last.FindStringSubmatch(s)[1])%9 + 1
		}

		return
	}

	fmt.Println(calc(nums[:9]))
	fmt.Println(calc(nums))

}
func findSubstringIndex(s, substring string) int {
	index := strings.Index(s, substring)
	return index
}

func getShitDone(s string) (int, error) {
	runes := []rune(s)
	var firstDigitIndex int = -1
	var firstSpeltIndex int = len(s)
	var firstDigit int = -1
	var firstSpelt int = -1
	var firstRes int
	var err error

	// Scan for first digit
	for i := 0; i < len(runes); i++ {
		if runes[i] >= '0' && runes[i] <= '9' {
			firstDigitIndex = i
			firstDigit, err = strconv.Atoi(string(runes[i]))
			if err != nil {
				return 0, err
			}
			break
		}
	}

	// Scan for first spelled out number
	for word, number := range numberMap {
		index := findSubstringIndex(s, word)
		if index != -1 && (firstSpeltIndex == len(s) || index < firstSpeltIndex) {
			firstSpeltIndex = index
			firstSpelt, err = strconv.Atoi(number)
			if err != nil {
				return 0, err
			}
		}
	}

	// Determine the earliest found index
	if firstDigitIndex == -1 && firstSpeltIndex == len(s) {
		return 0, fmt.Errorf("could not find first number in string")
	}
	if firstDigitIndex != -1 && (firstDigitIndex < firstSpeltIndex || firstSpeltIndex == len(s)) {
		firstRes = firstDigit
	} else {
		firstRes = firstSpelt
	}

	if firstDigitIndex > firstSpeltIndex && firstSpeltIndex != len(s) {
		runes = runes[firstSpeltIndex+1:]
	} else {
		runes = runes[firstDigitIndex+1:]
	}

	// Scan for second digit
	var secondDigitIndex int = -1
	var secondDigit int = -1
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= '0' && runes[i] <= '9' {
			secondDigitIndex = i
			secondDigit, err = strconv.Atoi(string(runes[i]))
			if err != nil {
				return 0, err
			}
			break
		}
	}

	// Scan for second spelled out number
	var secondSpeltIndex int = -1
	var secondSpelt int = -1
	for word, number := range numberMap {
		index := findSubstringIndex(string(runes), word)
		if index != -1 && index > secondSpeltIndex {
			secondSpeltIndex = index
			secondSpelt, err = strconv.Atoi(number)
			if err != nil {
				return 0, err
			}
		}
	}

	// Determine the latest found index
	if secondDigitIndex == -1 && secondSpeltIndex == -1 {
		res := 10*firstRes + firstRes
		return res, nil
	} else if secondDigitIndex != -1 && (secondDigitIndex > secondSpeltIndex || secondSpeltIndex == -1) {
		res := 10*firstRes + secondDigit
		return res, nil
	} else {
		res := 10*firstRes + secondSpelt
		return res, nil
	}

}
