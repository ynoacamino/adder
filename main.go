package main

import (
	"flag"
	"fmt"
	"strings"
)

func sumOne(a, b, cin bool) (bool, bool) {
	result := a != b != cin
	cout := (a && b) || (cin && (a != b))

	return result, cout
}

func sum4Bits(first, second [4]bool) ([4]bool, bool) {
	respose := [4]bool{}
	currentCout := false
	for i := len(first) - 1; i >= 0; i -= 1 {
		currentRes, cout := sumOne(first[i], second[i], currentCout)
		currentCout = cout
		respose[i] = currentRes
	}

	return respose, currentCout
}

func stringToBits(number string) [4]bool {
	res := [4]bool{}
	arr := strings.Split(number, "")
	for i := 0; i < len(arr); i++ {
		res[3-i] = arr[3-i] == "1"
	}
	fmt.Println("\t\t\t", res)

	return res
}

func main() {
	var firstStr, secondStr string
	flag.StringVar(&firstStr, "first", "0000", "First 4-bit binary number")
	flag.StringVar(&secondStr, "second", "0000", "Second 4-bit binary number")
	flag.Parse()

	firstNum := stringToBits(firstStr)
	secondNum := stringToBits(secondStr)

	result, cout := sum4Bits(firstNum, secondNum)

	fmt.Println("Result: \t\t", result)
	fmt.Println("Cout: \t\t\t", cout)
}
