package methods

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Operation string
	Stop      string
	Start     string
	File      string
}

func (c Config) Valid() error {
	if c.File == "" {
		return errors.New("no file")
	}

	if c.Operation != "count" && c.Operation != "get" {
		return errors.New("unknown operation")
	}

	return nil
}

func parseConfig(configParts []string) Config {
	c := Config{}

	for _, item := range configParts {
		if strings.Index(item, "op:") == 0 {
			c.Operation = item[3:]
		}

		if strings.Index(item, "start:") == 0 {
			c.Start = item[6:]
		}

		if strings.Index(item, "stop:") == 0 {
			c.Stop = item[5:]
		}

		if strings.Index(item, "file:") == 0 {
			c.File = item[5:]
		}
	}

	return c
}

func runScan() {

	if len(os.Args) < 2 {
		fmt.Println("no arguments")
		return
	}

	c := parseConfig(os.Args[1:])
	fmt.Println(c)
	fmt.Println(c.Valid())
}
