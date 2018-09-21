package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("Hey! I need one argument!")
	} else {
		str := args[1]
		var num float64

		t1 := time.Now().UnixNano()
		for i := 0; i < 1000000; i++ {
			num = oldParser(str)
		}
		t2 := time.Now().UnixNano()
		for i := 0; i < 1000000; i++ {
			num = newParser(str)
		}
		t3 := time.Now().UnixNano()

		fmt.Println("Old parser: ", t2-t1)
		fmt.Println("New parser: ", t3-t2)
		fmt.Println(num) // just to avoid fails and code optimization
	}
}

func oldParser(str string) float64 {
	num, _ := strconv.ParseFloat(str, 64)
	return num
}

var floatRegExp, _ = regexp.Compile("[0-9]+[.|,]?[0-9]+")
var floatSeparatorRegExp, _ = regexp.Compile("[.|,]")

func newParser(str string) float64 {
	parsed := floatRegExp.FindString(str)
	parsed = floatSeparatorRegExp.ReplaceAllString(parsed, ".")

	num, _ := strconv.ParseFloat(parsed, 64)

	return num
}
