package gwc

import (
	"os"
	"log"

	"github.com/urfave/cli/v2"
	
	"github.com/ObsidianDestroyer/gwc/count_bytes"
)


func main() {
	commonFlags := []cli.Flag{
		&cli.BoolFlag{
			Name: "c",
			Usage: "Print count of bytes",
			Aliases: []string{"bytes"},
		},
		&cli.BoolFlag{
			Name: "m",
			Usage: "Print count of chars",
			Aliases: []string{"chars"},
		},
		&cli.BoolFlag{
			Name: "l",
			Usage: "Print count of lines",
			Aliases: []string{"lines"},
		},
	}
	count_bytes.ReadFile()
	app := &cli.App{
		Name: "gwc",
		Usage: "Prints the number of characters of each line, words and bytes for each FILE and rmultiple files. According to a sequence of non-zero sampled characters, separated by whitespace.",
		Flags: commonFlags,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
