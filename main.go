package main

import (
	"fmt"
	"golang-lab/functions"
	"golang-lab/variables"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Missing action parameter")
		os.Exit(1)
		return
	}

	action := os.Args[1]

	if action == "variables" {
		fmt.Println("---------- VARIABLES ----------")
		variables.RunExamples()

		fmt.Println("---------- VARIABLES CHALLANGES ----------")
		variables.RunChallanges()
	}

	if action == "functions" {
		fmt.Println("---------- FUNCTIONS ----------")
		functions.RunExamples()

		fmt.Println("---------- FUNCTIONS CHALLANGES ----------")
		functions.RunChallanges()
	}
}
