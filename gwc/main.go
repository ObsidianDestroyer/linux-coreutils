package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	applicationFlags := []cli.Flag{
		&cli.BoolFlag{
			Name: "c",
			Usage: "Print count of bytes",
			Aliases: []string{"bytes"},
		},
	}

	app := &cli.App{
		Name:  "gwc",
		Usage: "Prints the number of characters of each line, words and bytes for each FILE and rmultiple files. According to a sequence of non-zero sampled characters, separated by whitespace.",
		Flags:  applicationFlags,
		Action: execute,
	}
	err := app.Run(os.Args)
	check(err)
}

func execute(cli *cli.Context) error {
	filePath := cli.Args().Get(0)
	if filePath == "" {
		fmt.Println("Please, type a path to file.")
		os.Exit(3)
	} else {
		byteStream := readFile(filePath)
		count := countBytes(byteStream)
		fmt.Print(count)
	}
	mustCountBytes := cli.Bool("c")
	if mustCountBytes == true {

	}
	return nil
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		cli.Exit("Encountered unhandled error", 1)
	}
}

func readFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	check(err)
	return file
}

func countBytes(byteStream []byte) int {
	return bytes.Count(byteStream, []byte(""))
}
