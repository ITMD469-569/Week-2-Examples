package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var num1String = "1"
	var num1 int = stringToInt(num1String)
	fmt.Println(1 + num1)

}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func stringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return f
}
