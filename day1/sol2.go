
package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)


func main(){
	totalResult := 0

	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input, _ := strconv.Atoi(scanner.Text())
		totalResult += calculateModuleFuel(input)
	}
	fmt.Println(totalResult)
}

func calculateModuleFuel(fuel int) int {
	moduleFuel := int(fuel / 3) - 2
	newFuel := moduleFuel

	for newFuel > 0 {
		newFuel = int(newFuel / 3) - 2
		if newFuel < 0 {
			break
		}
		moduleFuel += newFuel
	}

	return moduleFuel
}