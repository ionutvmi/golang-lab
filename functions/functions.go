package functions

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

type activity struct {
	name    string
	execute func(int, int) (int, error)
}

func example1() {
	fmt.Println(os.Args)

	a, err := strconv.Atoi(os.Args[2])
	fmt.Println("err a=", err)

	b, err := strconv.Atoi(os.Args[3])
	fmt.Println("err b=", err)

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	res, err := division(a, b)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("a / b = ", res)
	}

	print := func(args ...any) {
		fmt.Println(args...)
	}

	print("Hello", "there", "inline", "functions")

	activities := []activity{
		{
			name:    "adder",
			execute: adder,
		},
		{
			name:    "subtractor",
			execute: subtractor,
		},
		{
			name:    "multiplier",
			execute: multiplier,
		},
		{
			name:    "division",
			execute: division,
		},
	}

	for _, currentActivity := range activities {
		val, err := currentActivity.execute(a, b)
		if err != nil {
			print("ERROR:", currentActivity.name, err)
		} else {
			print(currentActivity.name, val)
		}
	}

	// dynamically extract the name of the function
	print(runtime.FuncForPC(reflect.ValueOf(activities[0].execute).Pointer()).Name())

}

func division(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}

func adder(a, b int) (int, error) {
	x := a + b

	// check for overflow
	if (x > a) != (b > 0) {
		return x, errors.New("addition out of bounds")
	}
	return x, nil
}

func subtractor(a, b int) (int, error) {
	x := a - b

	// check for overflow
	if (x < a) != (b > 0) {
		return x, errors.New("subtraction out of bounds")
	}
	return x, nil
}

func multiplier(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}
	x := a * b

	// check for overflow
	if x/b != a {
		return x, errors.New("multiplication out of bounds")
	}

	return x, nil
}

func RunExamples() {
	example1()
}

func RunChallanges() {

}
