package main

import (
	"bufio"
	"fmt"
	"os"
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

	var res int32 = 1
	var line1text string
	var line2text string
	if scanner.Scan() {
		line1text = scanner.Text()
	}
	if scanner.Scan() {
		line2text = scanner.Text()
	}
	timesString := strings.TrimSpace(strings.Split(line1text, ":")[1])
	distancesString := strings.TrimSpace(strings.Split(line2text, ":")[1])

	var timeArr []int32
	var distanceArr []int32

	times := strings.Split(timesString, " ")
	distances := strings.Split(distancesString, " ")

	for _, time := range times {
		if time == "" {
			continue
		}
		timeInt, err := strconv.Atoi(time)
		if err != nil {
			fmt.Println("Error converting time: ", err)
			return
		}
		timeArr = append(timeArr, int32(timeInt))
	}

	for _, distance := range distances {
		if distance == "" {
			continue
		}
		distanceInt, err := strconv.Atoi(distance)
		if err != nil {
			fmt.Println("Error converting distance: ", err)
			return
		}
		distanceArr = append(distanceArr, int32(distanceInt))
	}
	fmt.Printf("Times: %v\n", timeArr)
	fmt.Printf("Distances: %v\n", distanceArr)

	for i := 0; i < len(timeArr); i++ {
		possibleWays, err := calculatePossibleWays(timeArr[i], distanceArr[i])
		if err != nil {
			fmt.Println("Error calculating possible ways: ", err)
			return
		}
		res *= possibleWays
	}

	fmt.Printf("Result: %d\n", res)
}

func calculatePossibleWays(time int32, distance int32) (int32, error) {
	var res int32 = 0

	for i := int32(0); i < time; i++ {
		distanceTravelled := i * (time - i)
		if distanceTravelled >= distance {
			res++
		}
	}

	return res, nil
}
