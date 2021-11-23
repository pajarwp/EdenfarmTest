package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Println("Enter list number in array: ")
	str, _ := input.ReadString('\n')
	var listNumber []int
	for i := 0; i < len(str)-1; i++ {
		intStr, _ := strconv.Atoi(string(str[i]))
		listNumber = append(listNumber, intStr)
	}
	biggestNumber := getBiggestNumber(listNumber)
	fmt.Println(biggestNumber)

}

func getBiggestNumber(listNumber []int) int {
	var arrayResult []int
	for i, _ := range listNumber {
		if listNumber[i] == listNumber[len(listNumber)-1-i] {
			arrayResult = append(arrayResult, listNumber[i])
		} else if listNumber[i] == listNumber[len(listNumber)-1-i] && len(listNumber) != 0 {
			break
		}
	}
	var biggestNumber int
	for i := 0; i < len(arrayResult)/2; i++ {
		for j := i + 1; j < len(arrayResult)/2; j++ {
			if arrayResult[i] > arrayResult[j] && arrayResult[i] > biggestNumber {
				biggestNumber = arrayResult[i]
			} else if arrayResult[i] < arrayResult[j] && arrayResult[j] > biggestNumber {
				biggestNumber = arrayResult[j]
			}
		}
	}
	return biggestNumber
}
