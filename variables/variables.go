package variables

import (
	"fmt"
	"strings"
)

func example1() {
	var x int = 10 // full variable declaration format
	var y = 20     // type isn't needed, since the default type for the numeric literal 20 is int
	var z int      // type is specified, but value isn't, initializing z to the zero value for int (0)

	fmt.Println(x, y, z, "nice", "inline")

	f := 3.2
	m := 10
	fmt.Println(f, m)

	a := 10
	b := 20

	fmt.Println("a, b = ", a, b)

	// swap
	b, a = a, b

	fmt.Println("a, b = ", a, b)

}

func example2() {
	s1 := "This is a string \t inline\nAdd a new\v line"

	s2 := `This
is a multiline string
it has
    multiple
lines`

	fmt.Println(s1)
	fmt.Println(s2)

}

func example3() {
	var nums []int
	otherNums := []int{4, 5, 6}
	preAllocated := make([]int, 15)

	for i := range preAllocated {
		preAllocated[i] = i * 2
	}

	fmt.Println(nums, otherNums, preAllocated)

	// copies by reference
	a := preAllocated

	a[0] = 99
	fmt.Println("a = ", a)
	fmt.Println("preAllocated = ", preAllocated)

	// !! also copy by reference
	b := a[0:3]

	b[0] = 55

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	newSliceCopy := make([]int, len(a))
	copy(newSliceCopy, a)

	fmt.Println("newSliceCopy = ", newSliceCopy)
	newSliceCopy[0] = 200

	fmt.Println("a = ", a)
	fmt.Println("newSliceCopy = ", newSliceCopy)

	otherCopy := a[0:0]
	fmt.Println("a = ", a)
	fmt.Println("otherCopy = ", otherCopy)

	// slice expression

	fmt.Println("a[0:1]", a[0:1])
	fmt.Println("a[:1]", a[:1])
	fmt.Println("a[2:]", a[2:])
	//fmt.Println("a[:-2]", a[:-2])
	//fmt.Println("a[-3:]", a[-3:])

}

func example4() {
	s1 := "A💗BC"

	fmt.Println("s1=", s1)
	fmt.Println("[]byte(s1)=", []byte(s1))

	fmt.Println("[]rune(s1)=", []rune(s1))
	fmt.Println("string([]rune(s1))=", string([]rune(s1)))

	chars := strings.Split(s1, "")
	fmt.Println("chars=", chars)

	fmt.Printf("chars[1] = %v\n", chars[1])
	fmt.Printf("type chars[1] = %T\n", chars[1])
}

func example5() {
	itemToPrice := map[string]float64{
		"tomatos": 12.53,
		"apples":  4,
	}

	itemToPrice["cucumbers"] = 11.59

	fmt.Println("itemToPrice=", itemToPrice)

	for key, val := range itemToPrice {
		fmt.Println(key, "=", val)
	}

	fmt.Println("itemToPrice=", itemToPrice)

	delete(itemToPrice, "apples")

	fmt.Println("after delete itemToPrice=", itemToPrice)

	if v, ok := itemToPrice["tomatos"]; ok {
		fmt.Println("The key tomatos exists and it has the value", v)
	}

	if v, ok := itemToPrice["apples"]; ok {
		fmt.Println("The key apples exists and it has the value", v)
	}

}

func challange() {
	var first byte = 10
	var second = "hello"

	third := 3.2

	third = third + float64(first)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}

// RunExamples executes the code included in the examples
func RunExamples() {
	example1()
	example2()
	example3()
	example4()
	example5()

}

func RunChallanges() {
	challange()
	runStrings()
	runSlices()
	runMaps()
}
