
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
	
	jumps := getAllJumps("SAN", "COM", orbitsMap)

	for index, jump := range jumps {
		jumps2 := getAllJumps("YOU", jump, orbitsMap)
		if len(jumps2) > 0 {
			// -3 to remove the initial orbits and the intersecting orbit
			fmt.Printf("Jumps: %d", index + len(jumps2) - 3)
			os.Exit(0)
		} 
	}
}

func getAllJumps(from, to string, orbitsMap map[string]Orbiter) []string {
	var jumps []string

	jumps = append(jumps, from)
	orbiter := orbitsMap[from]

	for true {
		jumps = append(jumps, orbiter.orbits)
		if orbiter.orbits == to {
			break;
		}
		if orbiter.orbits == "" {
			return make([]string, 0)
		}
		orbiter = orbitsMap[orbiter.orbits]
	}

	return jumps
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