package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)


func main(){
	program := fillArray()

	for a := 0; a < 100; a++ {
		for b := 0; b < 100; b++ {
			tempProgram := make([]int, len(program))
			copy(tempProgram, program)
			tempProgram[1] = a
			tempProgram[2] = b
			res := executeProgram(tempProgram)
			if res == 19690720{
				fmt.Printf("%d", (100 * a + b))
				os.Exit(0)
			}
		}
	}
}

func fillArray() []int {
	file, _ := os.Open("input")
	defer file.Close()

	var returnSlice []int

	result, _ := ioutil.ReadAll(file)
	resultString := fmt.Sprintf("%s", result)
	stringSplice := strings.Split(resultString, ",")

	for _, item := range stringSplice {
		tempInt, _ := strconv.Atoi(item) 
		returnSlice = append(returnSlice, tempInt)
	} 

	return returnSlice
}

func executeProgram(program []int) int {
	for i := 0; i < len(program); {
		if program[i] == 1 {
			tempRes := program[program[i+1]] + program[program[i+2]]
			program[program[i+3]] = tempRes
		} else if program[i] == 2 {
			tempRes := program[program[i+1]] * program[program[i+2]]
			program[program[i+3]] = tempRes
		} else if program[i] == 99 {
			break
		} else {
			fmt.Println("Something went wrong")
			break
		}
		i += 4
	}

	return program[0]
}