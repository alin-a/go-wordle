package console

import (
	"fmt"
	"log"
	"strings"
)

func readInputFromConsole() string {
	var input string

	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	return strings.ToUpper(input)
}
