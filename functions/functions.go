package functions

import (
	"fmt"
	"os"
	"strconv"
)

func example1() {
	fmt.Println(os.Args)

	a, err := strconv.Atoi(os.Args[2])
	fmt.Println("err a=", err)

	b, err := strconv.Atoi(os.Args[3])
	fmt.Println("err b=", err)

	fmt.Println("a=", a)
	fmt.Println("b=", b)
}

func RunExamples() {
	example1()
}

func RunChallanges() {

}
