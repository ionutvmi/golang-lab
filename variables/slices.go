package variables

import "fmt"

func RunSlices() {
	var values []int

	fmt.Println(len(values))

	values = append(values, 20, 50, 200)

	fmt.Println(len(values))

	start := values[0:2]
	end := values[len(values)-2:]

	fmt.Println(values)
	fmt.Println(start)
	fmt.Println(end)

}
