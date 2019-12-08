package main

import (
	"strconv"
	"fmt"
)

func main() {
	num := 130254
	max := 678275

	numValids := 0

	for ; num < max; num++ {
		if checkValid(num) {
			numValids++
		}
	}

	fmt.Printf("%d", numValids)
}

func checkValid(num int) bool {
	number := strconv.Itoa(num)
	return isIncreasingDigits(number) && hasDoubleDigits(number)
}

func isIncreasingDigits(num string) bool {
	highestNum := '0'
	for _, item := range num {
		if item < highestNum {
			return false
		} else {
			highestNum = item
		}
	}
	return true
}

func hasDoubleDigits(num string) bool {
	var prev2Num rune
	var prevNum rune
	for index, item := range num {
		lastIndex := index == len(num) - 1
		if lastIndex && item == prevNum && item != prev2Num {
			return true
		} else if !lastIndex && item == prevNum && item != prev2Num && item != rune(num[index+1]) {
			return true
		}

		prev2Num = prevNum
		prevNum = item
	}
	return false
}