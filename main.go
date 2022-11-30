package main

import (
	"fmt"
	"golang-lab/functions"
	"golang-lab/logger"
	"golang-lab/methods"
	"golang-lab/variables"
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

		fmt.Println("---------- VARIABLES CHALLENGES ----------")
		variables.RunChallenges()
	}

	if action == "functions" {
		fmt.Println("---------- FUNCTIONS ----------")
		functions.RunExamples()

		fmt.Println("---------- FUNCTIONS CHALLENGES ----------")
		functions.RunChallenges()
	}

	if action == "methods" {
		fmt.Println("---------- METHODS ----------")
		methods.RunExamples()

		fmt.Println("---------- METHODS CHALLENGES ----------")
		methods.RunChallenges()
	}

}
