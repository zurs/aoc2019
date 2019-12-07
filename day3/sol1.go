
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
	line2Points := calcLinePoints(input2)

	_, distance := findClosestIntersection(line1Points, line2Points)

	fmt.Printf("%d", distance)
}

func findClosestIntersection(line1Points, line2Points []Point) (Point,int) {
	closest := 0
	var closestPoint Point 
	for _, point1 := range line1Points {
		for _, point2 := range line2Points {
			if point1.x == point2.x && point1.y == point2.y && point1.x != 0 && point1.y != 0 {
				distance := abs(point1.x) + abs(point1.y)
				if distance < closest || closest == 0 {
					closest = distance
					closestPoint = point1
				}
			}
		}
	}

	return closestPoint, closest
}

func splitInputs(input []byte) ([]string, []string) {
	s := string(input)
	inputs := strings.Split(s, "\r\n")
	input1 := strings.Split(inputs[0], ",")
	input2 := strings.Split(inputs[1], ",")
	return input1, input2
}

func calcLinePoints(inputs []string) []Point {
	var resultPoints []Point

	resultPoints = append(resultPoints, Point{0, 0})

	for _, item := range inputs {
		distance, _ := strconv.Atoi(item[1:])
		for i := 0; i < distance; i++ {
			newPoint := resultPoints[len(resultPoints)-1]
			if item[0] == 'U' {
				newPoint.y++
			} else if item[0] == 'D' {
				newPoint.y--
			} else if item[0] == 'L' {
				newPoint.x--
			} else if item[0] == 'R' {
				newPoint.x++
			}

			resultPoints = append(resultPoints, newPoint)
		}
	}

	return resultPoints
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
