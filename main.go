package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/dimchansky/utfbom"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Usage: "Remove BOM",
		Action: func(c *cli.Context) error {
			reader := utfbom.SkipOnly(bufio.NewReader(os.Stdin))
			if _, err := io.Copy(os.Stdout, reader); err != nil {
				return err
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
