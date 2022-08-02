package gwc

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(pathToFile string) {
	file, err := os.ReadFile(pathToFile)
	check(err)
	fmt.Print(string(file))
}

