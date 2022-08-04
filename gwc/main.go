package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
		&cli.BoolFlag{
			Name:    "m",
			Usage:   "Print count of chars",
			Aliases: []string{"chars"},
		},
		&cli.BoolFlag{
			Name:    "l",
			Usage:   "Print count of lines",
			Aliases: []string{"lines"},
		},
		&cli.BoolFlag{
			Name:    "L",
			Usage:   "Print max length of line",
			Aliases: []string{"max-line-length"},
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
	var bytesCount int
	var charsCount int
	var linesCount int
	var maxLineLength int

	filePath := cli.Args().Get(0)
	fileName := filepath.Base(filePath)

	if filePath == "" {
		fmt.Println("Please, type a path to file.")
		os.Exit(3)
	}

	byteStream := readFile(filePath)

	if cli.Bool("c") == true {
		bytesCount = countBytes(byteStream)
	}
	if cli.Bool("m") == true {
		charsCount = countChars(byteStream)
	}
	if cli.Bool("l") == true {
		linesCount = countLines(byteStream)
	}
	if cli.Bool("L") == true {
		maxLineLength = countMaxLineLength(byteStream)
	}

	fmt.Println(bytesCount, charsCount, linesCount, maxLineLength, fileName)
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

func countChars(byteStream []byte) int {
	return len(byteStream)
}

func countLines(byteStream []byte) int {
	return len(strings.Split(string(byteStream), "\n"))
}

func countMaxLineLength(byteStream []byte) int {
	lines := strings.Split(string(byteStream), "\n")
	linesLength := make([]int, len(lines))
	maxLength := linesLength[0]
	for _, line := range lines {
		linesLength = append(linesLength, len(line))
	}
	for index := 0; index < len(linesLength); index++ {
		if maxLength < linesLength[index] {
			maxLength = linesLength[index]
		}
	}
	return maxLength
}
