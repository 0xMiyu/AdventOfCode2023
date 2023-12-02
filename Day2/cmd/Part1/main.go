package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MaxRed = 12
var MaxGreen = 13
var MaxBlue = 14

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res int32 = 0
	for scanner.Scan() {
		lineText := scanner.Text()
		toAdd, err := getNumberToAdd(lineText)
		if err != nil {
			fmt.Println("Error parsing line: ", err)
		}
		res += int32(toAdd)
	}
	fmt.Printf("Result: %d\n", res)
}

func checkIfPossible(input string) (bool, error) {
	gameStrings := strings.Split(input, ", ")
	for _, gameString := range gameStrings {
		element := strings.Split(gameString, " ")
		switch element[1] {
		case "red":
			numRed, err := strconv.Atoi(element[0])
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
			}
			if numRed > MaxRed {
				return false, nil
			}
		case "green":
			numGreen, err := strconv.Atoi(element[0])
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
			}
			if numGreen > MaxGreen {
				return false, nil
			}
		case "blue":
			numBlue, err := strconv.Atoi(element[0])
			if err != nil {
				fmt.Println("Error converting string to int: ", err)
			}
			if numBlue > MaxBlue {
				return false, nil
			}
		}
	}
	return true, nil
}

func getNumberToAdd(line string) (int, error) {
	stringWithGameNumber := strings.Split(line, ":")[0]
	gameId, err := strconv.Atoi(strings.Split(stringWithGameNumber, " ")[1])
	if err != nil {
		fmt.Println("Error converting string to int: ", err)
	}

	gamesString := strings.Split(line, ":")[1]
	games := strings.Split(gamesString, ";")

	res := gameId
	for _, game := range games {
		isValid, err := checkIfPossible(strings.TrimSpace(game))
		if err != nil {
			fmt.Println("Error checking if possible: ", err)
		}
		if !isValid {
			res = 0
			break
		}
	}

	return res, nil
}
