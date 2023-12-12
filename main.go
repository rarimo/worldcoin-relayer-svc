package main

import (
	"os"

	"github.com/rarimo/worldcoin-relayer-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
