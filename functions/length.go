package functions

import (
	"fmt"
	"os"
)

func printParams(params []string) {
	for _, val := range params {
		fmt.Println(val)
	}
}

func printFileFromParams(params []string) {
	for _, filePath := range params {
		data, err := os.ReadFile(filePath)

		// filePathParts := strings.Split(filePath, "/")
		// fileName := filePathParts[len(filePathParts)-1]

		// fileName := path.Base(filePath)

		fileName := filePath // jebaited

		if err != nil {
			fmt.Println(fileName, "-", err.Error())
			continue
		}

		fi, err := os.Stat(filePath)

		if err != nil {
			fmt.Println(fileName, "-", err.Error())
			continue
		}

		fmt.Println(fileName, "-", data[:10], fi.Size())
	}
}

func runLength() {
	printParams(os.Args[1:])
	printFileFromParams(os.Args)
}
