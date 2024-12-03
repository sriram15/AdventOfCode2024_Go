package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func MullItOver() {

	data, err := os.ReadFile("./day3/input_test.txt")
	if err != nil {
		fmt.Println("Error while reading the files")
		return
	}

	dataStr := string(data)

	count := MulPart1(dataStr)

	fmt.Println(count)

}

func MulPart1(dataStr string) int64 {

	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(dataStr, -1)

	var count int64
	for _, match := range matches {

		leftItem, err1 := strconv.ParseInt(match[1], 10, 64)
		rightItem, err2 := strconv.ParseInt(match[2], 10, 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Error while converting string to int")
			return 0

		}

		count += leftItem * rightItem
	}
	return count
}
