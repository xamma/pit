package main

import (
	"fmt"
	"os"

	"github.com/xamma/pit/internal/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "pit:", err)
		os.Exit(1)
	}
}
