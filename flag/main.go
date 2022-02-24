package main

import (
	"flag"
	"fmt"
	"strconv"
)

// ~/workspace/go-workspace/gcode/flag/ [master+*] go run main.go --input 1122 --partB true
//partB
//2
//partB
//2
//partB
//4
//partB
//4
// ~/workspace/go-workspace/gcode/flag/ [master+*]

// ~/workspace/go-workspace/gcode/flag/ [master+*] go run main.go  --partB true
//partB
//2
//partB
//4
//partB
//6
//partB
//8
// ~/workspace/go-workspace/gcode/flag/ [master+*]

var input = flag.String("input", "1234", "The input to the problem.")
var partB = flag.Bool("partB", true, "Whether to use the Part B logic.")

func main() {
	flag.Parse()
	digits := make([]int, len(*input))

	for i, c := range *input {
		curVal, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Printf("Couldn't parse: %v\n", err)
			return
		}
		digits[i] = curVal
	}

	for _, curVal := range digits {
		if *partB {
			fmt.Println("partB")
			fmt.Println(2 * curVal)
		} else {
			fmt.Println("noPartB")
			fmt.Println(curVal)
		}
	}
}
