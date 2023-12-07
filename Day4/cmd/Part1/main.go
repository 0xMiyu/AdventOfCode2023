package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

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
		stringWithoutCardId := strings.Split(lineText, ":")[1]
		winningNumbersString := strings.Split(stringWithoutCardId, "|")[0]
		numbersYouHaveString := strings.Split(stringWithoutCardId, "|")[1]

		winningNumbersArr, err := getIntArrayFromString(winningNumbersString)
		if err != nil {
			fmt.Println("Error getting int array from string: ", err)
			return
		}

		numbersYouHaveArr, err := getIntArrayFromString(numbersYouHaveString)
		if err != nil {
			fmt.Println("Error getting int array from string: ", err)
			return
		}

		res += calculateScore(winningNumbersArr, numbersYouHaveArr)
	}
	fmt.Printf("Score: %d\n", res)
}

func calculateScore(winningNumbers []int, numbersYouHave []int) int32 {
	var res int32 = 0
	for _, number := range winningNumbers {
		if slices.Contains(numbersYouHave, number) {
			res++
		}
	}
	return int32(math.Pow(2, float64(res-1)))
}

func getIntArrayFromString(line string) ([]int, error) {
	var res []int
	numbers := strings.Split(line, " ")
	for _, number := range numbers {
		if strings.TrimSpace(number) == "" {
			continue
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		res = append(res, num)

	}
	return res, nil
}

func getId(line string) (int, error) {
	cardString := strings.Split(line, ":")[0]
	cardNumberString := strings.Split(cardString, " ")[1]
	res, err := strconv.Atoi(cardNumberString)
	if err != nil {
		return 0, err
	}
	return res, nil
}
