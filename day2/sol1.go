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

	program[1] = 12
	program[2] = 2

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

	resultString := ""

	for index, item := range program {
		if index != 0 {
			resultString += ","
		}
		resultString += fmt.Sprintf("%d", item)
	}

	fmt.Println(resultString)
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
