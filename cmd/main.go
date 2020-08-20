package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lifina/go-template/internal/cmd/bgo"
)

func main() {
	baseName := filepath.Base(os.Args[0])

	err := bgo.NewCommand(baseName).Execute()
	if err != nil {
		if err != context.Canceled {
			fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		}
		os.Exit(1)
	}
}
