package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

type Opcode struct {
	code int
	parameters []Parameter
}

type Parameter struct {
	value int
	mode int
}


func main(){
	program := fillProgram()

	for i := 0; i < len(program); {
		opcode := parseOpcode(program, i)
		//fmt.Printf("Executing: %v\n", opcode)
		jump := opcode.executeOpcode(&program)
		if jump != 0 {
			i = jump
		} else {
			i += opcode.noOfParameters() + 1
		}
		
	}
}

func getInput() int {
	return 5
}

func (opcode *Opcode) executeOpcode(program *[]string) int {
	jump := 0

	if opcode.noOfParameters() >= 2 {
		param1 := opcode.parameters[0].getValue(*program)
		param2 := opcode.parameters[1].getValue(*program)
		if opcode.code == 1 {
			// Addition
			result := param1 + param2
			opcode.parameters[2].storeValue(program, result)
		} else if opcode.code == 2 {
			// Multiplication
			result := param1 * param2
			opcode.parameters[2].storeValue(program, result)
		} else if opcode.code == 5 {
			// Jump-if-true
			if param1 != 0 {
				jump = param2
			}
		} else if opcode.code == 6 {
			// Jump-if-false
			if param1 == 0 {
				jump = param2
			}
		} else if opcode.code == 7 {
			// Less-than
			if param1 < param2 {
				opcode.parameters[2].storeValue(program, 1)
			} else {
				opcode.parameters[2].storeValue(program, 0)
			}
		} else if opcode.code == 8 {
			// Equals
			if param1 == param2 {
				opcode.parameters[2].storeValue(program, 1)
			} else {
				opcode.parameters[2].storeValue(program, 0)
			}
		}
	} else if opcode.code == 3 {
		// Store input at a location
		input := getInput()
		opcode.parameters[0].storeValue(program, input)
	} else if opcode.code == 4 {
		// Print value at the postition
		fmt.Printf("Diagnostic code: %d\n", opcode.parameters[0].getValue(*program))
	} else if opcode.code == 99 {
		// Halt and exit
		fmt.Println("Halt!")
		os.Exit(0)
	}

	return jump
}

func (parameter Parameter) getValue(program []string) int {
	var value int
	if parameter.mode == 0 {
		value, _ = strconv.Atoi(program[parameter.value])
	} else {
		value = parameter.value
	}
	return value
}

func (parameter Parameter) storeValue(program *[]string, toStore int) {
	toStoreInt := strconv.Itoa(toStore)
	if parameter.mode == 0 {
		(*program)[parameter.value] = toStoreInt
	} else {
		fmt.Println("Tried to use a storage parameter with immediate mode")
		os.Exit(1)
	}
}

func parseOpcode(program []string, index int) Opcode {
	opcode := Opcode{}
	instruction := program[index]

	modes := ""
	if len(instruction) > 1 {
		opcode.code, _ = strconv.Atoi(string(instruction[len(instruction) - 2:])) // Get the 2 digit opcode
		modes = string(instruction[:len(instruction) - 2])
	} else {
		opcode.code, _ = strconv.Atoi(string(instruction[len(instruction) - 1:])) // Get the 1 digit opcode
	}

	paramIndex := index + 1
	parameters := program[paramIndex:paramIndex + opcode.noOfParameters()]
	opcode.parseParameters(parameters, modes)

	return opcode
}

func (opcode *Opcode) noOfParameters() int {
	code := opcode.code
	if code == 1 || code == 2 || code == 7 || code == 8 {
		return 3
	} else if code == 5 || code == 6 {
		return 2
	} else if code == 3 || code == 4 {
		return 1
	}
	return 0
}

func (opcode *Opcode) parseParameters(parameters []string, modes string) {
	for index, item := range parameters {
		paramValue, _ := strconv.Atoi(item)
		mode := 0
		if len(modes) > index {
			mode, _ = strconv.Atoi(string(modes[len(modes)-1 - index]))
		}
		newParam := Parameter{paramValue, mode}
		opcode.parameters = append(opcode.parameters, newParam)
	}
}


func fillProgram() []string {
	file, _ := os.Open("input")
	defer file.Close()

	var returnSlice []string

	result, _ := ioutil.ReadAll(file)
	resultString := fmt.Sprintf("%s", result)
	stringSplice := strings.Split(resultString, ",")

	for _, item := range stringSplice {
		returnSlice = append(returnSlice, string(item))
	}

	return returnSlice
}