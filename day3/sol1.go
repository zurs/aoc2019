
package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	content, _ := ioutil.ReadAll(file)
	input1, input2 := splitInputs(content)

	line1Points := calcLinePoints(input1)
	_ = calcLinePoints(input2)


	fmt.Printf("%v", line1Points)
}

func splitInputs(input []byte) ([]string, []string) {
	s := string(input)
	inputs := strings.Split(s, "\n")
	input1 := strings.Split(inputs[0], ",")
	input2 := strings.Split(inputs[1], ",")
	return input1, input2
}

func calcLinePoints(inputs []string) []Point {
	var resultPoints []Point

	resultPoints = append(resultPoints, Point{0, 0})

	for _, item := range inputs {
		newPoint := resultPoints[len(resultPoints)-1] 
		distance, _ := strconv.Atoi(item[1:])
		if item[0] == 'U' {
			newPoint.x += distance
		} else if item[0] == 'D' {
			newPoint.x -= distance
		} else if item[0] == 'L' {
			newPoint.y -= distance
		} else if item[0] == 'R' {
			newPoint.y += distance
		}

		resultPoints = append(resultPoints, newPoint)
	}

	return resultPoints
}
