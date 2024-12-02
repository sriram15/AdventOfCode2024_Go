package day1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func HistorianHysteria() {
	// Read the file input
	data, err := os.ReadFile("./day1/input_main.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create two slices to store left and right items
	var left, right []int64
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {

		// If there is some content, split using the 3 spaces provided in input
		if len(line) > 0 {

			fields := strings.Split(line, "   ")

			leftItem, lerr := strconv.ParseInt(fields[0], 10, 64)
			rightItem, rerr := strconv.ParseInt(fields[1], 10, 64)

			if lerr != nil || rerr != nil {
				fmt.Println(lerr, rerr)
				return
			}
			left = append(left, leftItem)
			right = append(right, rightItem)

		}
	}

	// Sort the left and right array to easily find the diff
	slices.Sort(left)
	slices.Sort(right)

	var diff int64 = 0
	var similarityScore int64 = 0
	for index, l := range left {

		// PART 1 - Calc diff
		if left[index] >= right[index] {
			diff += (left[index] - right[index])
		} else {
			diff += (right[index] - left[index])
		}

		// PART 2 - Calc similarity score
		similarityCount := 0
		for _, r := range right {
			if r == l {
				similarityCount++
			}
		}

		similarityScore += (l * int64(similarityCount))

	}

	fmt.Println("Diff is ", diff)
	fmt.Println("Similarity score is", similarityScore)
}
