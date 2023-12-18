package main

import "fmt"

func main() {
	response := well([]string{"good", "bad", "good"})

	fmt.Println(response)
}

func well(x []string) string {
	var goodVariables int

	for _, v := range x {
		if v == "good" {
			goodVariables++
		}
	}

	if goodVariables <= 2 {
		return "good"
	}
	return ""
}
