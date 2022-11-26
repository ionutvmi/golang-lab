package variables

import "fmt"

func RunMaps() {

	var scores = map[string]int{
		"Yankees": 10,
		"Red Sox": 5,
		"Dodgers": 7,
		"Cubs":    0,
	}

	fmt.Println(len(scores))
	fmt.Println(scores)

	val, ok := scores["Red Sox"]
	fmt.Println(val, ok)

	val, ok = scores["Cubs"]
	fmt.Println(val, ok)

	val, ok = scores["White Sox"]
	fmt.Println(val, ok)

	delete(scores, "Red Sox")

	val, ok = scores["Red Sox"]
	fmt.Println(val, ok)

	fmt.Println(scores)

}
