
package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	//"math"
)


func main(){
	totalResult := 0

	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		result := int(input / 3) - 2
		totalResult += result
	}
	fmt.Println(totalResult)
}