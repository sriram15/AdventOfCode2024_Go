package day2

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isValidReport(reports []string) bool {

	isAsc, isDesc := false, false
	for index := range len(reports) - 1 {

		// Parse the string input to integer
		currentItem, err1 := strconv.ParseInt(reports[index], 10, 64)
		nextItem, err2 := strconv.ParseInt(reports[index+1], 10, 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Error while converting the string to int")
			return false
		}

		// Set if the direction is ascending or decending for first iteration
		if index == 0 {
			if currentItem > nextItem {
				isDesc = true
			}
			if currentItem < nextItem {
				isAsc = true
			}
		}

		// Check if there is a change in the order and then break if so
		if (currentItem > nextItem && isAsc) || (currentItem < nextItem && isDesc) {
			return false
		}

		diff := math.Abs(float64(currentItem - nextItem))

		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func RedNosedReports() {

	safeReports := 0

	data, err := os.ReadFile("./day2/input_main.txt")
	if err != nil {
		fmt.Println("Error reading the file")
		return
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		reports := strings.Split(line, " ")
		if len(reports) == 0 || len(reports) == 1 {
			break
		}

		isSafe := isValidReport(reports)

		if isSafe {
			safeReports++
		}
	}

	fmt.Println("Safe report ", safeReports)
}
