package main

import (
	"fmt"
	"golang-lab/functions"
	"golang-lab/methods"
	"golang-lab/variables"
	"golang-lab/logger"
	"os"
)

func main() {

	defer logger.Close()

	
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

	if action == "methods" {
		fmt.Println("---------- METHODS ----------")
		methods.RunExamples()

		fmt.Println("---------- METHODS CHALLANGES ----------")
		methods.RunChallanges()
	}

}
