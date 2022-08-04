package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

var appName = "gwc"
var helpUsage = "Prints the number of characters of each line, words and bytes for each FILE and rmultiple files. " +
	"According to a sequence of non-zero sampled characters, separated by whitespace."

var IsLetter = regexp.MustCompile("^[a-zA-Z!@#$&()\\-`.+,/\"]*$").MatchString

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
		&cli.BoolFlag{
			Name:    "w",
			Usage:   "Print max length of line",
			Aliases: []string{"words"},
		},
	}
	app := &cli.App{
		Name:   appName,
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
	var wordsCount int
	var maxLineLength int

	filePath := cli.Args().Get(0)
	fileName := filepath.Base(filePath)
	completedResult := ""
	counts := make([]int, 0)

	if filePath == "" {
		fmt.Println("Please, type a path to file.")
		os.Exit(3)
	}
	byteStream := readFile(filePath)

	if cli.Bool("c") == true { // count bytes
		bytesCount = countBytes(byteStream)
		counts = append(counts, bytesCount)
	}
	if cli.Bool("m") == true { // count chars
		charsCount = countChars(byteStream)
		counts = append(counts, charsCount)
	}
	if cli.Bool("l") == true { // count lines
		linesCount = countLines(byteStream)
		counts = append(counts, linesCount)
	}
	if cli.Bool("L") == true { // count max len line
		maxLineLength = countMaxLineLength(byteStream)
		counts = append(counts, maxLineLength)
	}
	if cli.Bool("w") == true { // count words
		wordsCount = countWords(byteStream)
		counts = append(counts, wordsCount)
	}
	for _, count := range counts {
		if count != 0 {
			completedResult += fmt.Sprint(count) + " "
		}
	}
	fmt.Println(completedResult + fileName)
	return nil
}

func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	check(err)
	return file
}

func getLinesArrayFromByteStream(byteStream []byte) []string {
	return strings.Split(string(byteStream), "\n")
}

func countBytes(byteStream []byte) int {
	return bytes.Count(byteStream, []byte(""))
}

func countChars(byteStream []byte) int {
	return len(byteStream)
}

func countLines(byteStream []byte) int {
	return len(getLinesArrayFromByteStream(byteStream))
}

func countMaxLineLength(byteStream []byte) int {
	lines := getLinesArrayFromByteStream(byteStream)
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

func countWords(byteStream []byte) int {
	lines := getLinesArrayFromByteStream(byteStream)
	words := make([]string, 0)
	for _, line := range lines {
		splittedWords := strings.Fields(line)
		for _, word := range splittedWords {
			if IsLetter(word) {
				words = append(words, word)
			}
		}
	}
	return len(words)
}
