package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

func main() {
	timeStart := time.Now()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	var mu sync.Mutex
	var totalSum int32 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		wg.Add(1)

		go func(line string) {
			defer wg.Done()
			toAdd, err := getNumberToAdd(line)
			if err != nil {
				fmt.Println("Error in goroutine: ", err)
				return
			}

			mu.Lock()
			totalSum += int32(toAdd)
			mu.Unlock()
		}(lineText)
	}

	wg.Wait()
	fmt.Printf("Total Sum: %d\n", totalSum)
	fmt.Printf("Time taken: %v\n", time.Since(timeStart))
}

func getNumberToAdd(line string) (int, error) {
	gamesString := strings.Split(line, ":")[1]
	games := strings.Split(gamesString, ";")

	var minRed, minGreen, minBlue int = 0, 0, 0
	for _, game := range games {
		gameStrings := strings.Split(game, ", ")
		for _, gameString := range gameStrings {
			element := strings.Split(strings.TrimSpace(gameString), " ")
			switch element[1] {
			case "red":
				numRed, err := strconv.Atoi(element[0])
				if err != nil {
					fmt.Println("Error converting string to int: ", err)
				}
				if numRed > minRed {
					minRed = numRed
				}
			case "green":
				numGreen, err := strconv.Atoi(element[0])
				if err != nil {
					fmt.Println("Error converting string to int: ", err)
				}
				if numGreen > minGreen {
					minGreen = numGreen
				}
			case "blue":
				numBlue, err := strconv.Atoi(element[0])
				if err != nil {
					fmt.Println("Error converting string to int: ", err)
				}
				if numBlue > minBlue {
					minBlue = numBlue
				}
			}
		}
	}

	fmt.Printf("Min red: %d, min green: %d, min blue: %d\n", minRed, minGreen, minBlue)

	res := minRed * minBlue * minGreen

	return res, nil
}
