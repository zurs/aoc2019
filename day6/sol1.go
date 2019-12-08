
package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

type Orbiter struct {
	orbits string
}

func main() {
	orbitsMap := getOrbits()

	totalSum := 0
	totalSpawns := 0

	c := make(chan int)

	for index, _ := range orbitsMap {
		go countToCOM(orbitsMap[index], orbitsMap, c)
		totalSpawns++
	}

	for totalSpawns > 0 {
		totalSum += <-c
		totalSpawns--
	}

	fmt.Printf("%d", totalSum)
}

func countToCOM(orbiter Orbiter, orbitsMap map[string]Orbiter, c chan int) {
	orbitsCount := 0

	for true {
		orbitsCount++
		if orbiter.orbits == "COM" {
			break;
		}
		orbiter = orbitsMap[orbiter.orbits]
	}

	c <- orbitsCount
}

func getOrbits() map[string]Orbiter {
	file, _ := os.Open("input")
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	inputs := strings.Split(string(content), "\r\n")

	orbiters := make(map[string]Orbiter)

	for _, item := range inputs {
		names := strings.Split(item, ")")
		orbiters[names[1]] = Orbiter{names[0]}
	}

	return orbiters
}