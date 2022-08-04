package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var helpUsage = "Prints the number of characters of each line, words and bytes for each FILE and rmultiple files. " +
	"According to a sequence of non-zero sampled characters, separated by whitespace."

func main() {
	applicationFlags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "c",
			Usage:   "Print count of bytes",
			Aliases: []string{"bytes"},
		},
	}

	app := &cli.App{
		Name:   "gwc",
		Usage:  helpUsage,
		Flags:  applicationFlags,
		Action: execute,
	}
	err := app.Run(os.Args)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		cli.Exit("Encountered unhandled error", 1)
	}
}

func execute(cli *cli.Context) error {
	var completeString string
	var count int

	filePath := cli.Args().Get(0)
	fileName := filepath.Base(filePath)

	if filePath == "" {
		fmt.Println("Please, type a path to file.")
		os.Exit(3)
	}

	byteStream := readFile(filePath)

	if cli.Bool("c") == true {
		count = countBytes(byteStream)
		completeString += fmt.Sprint(count)
	}

	fmt.Println(completeString, fileName)
	return nil
}

func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	check(err)
	return file
}

func countBytes(byteStream []byte) int {
	return bytes.Count(byteStream, []byte(""))
}
