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

	var res int64 = 0
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

	var timeString string = ""
	var distanceString string = ""

	times := strings.Split(timesString, " ")
	distances := strings.Split(distancesString, " ")

	for _, time := range times {
		if time == "" {
			continue
		}
		timeString += time
	}

	for _, distance := range distances {
		if distance == "" {
			continue
		}
		distanceString += distance
	}

	time, err := strconv.Atoi(timeString)
	if err != nil {
		fmt.Println("Error converting time: ", err)
		return
	}
	distance, err := strconv.Atoi(distanceString)
	if err != nil {
		fmt.Println("Error converting distance: ", err)
		return
	}

	fmt.Printf("Time: %d\n", int64(time))
	fmt.Printf("Distance: %d\n", int64(distance))

	res, err = calculatePossibleWays(int64(time), int64(distance))
	if err != nil {
		fmt.Println("Error calculating possible ways: ", err)
		return
	}
	fmt.Printf("Result: %d\n", res)
}

func calculatePossibleWays(time int64, distance int64) (int64, error) {
	var res int64 = 0

	for i := int64(0); i < time; i++ {
		distanceTravelled := i * (time - i)
		if distanceTravelled >= distance {
			res++
		}
	}

	return res, nil
}
