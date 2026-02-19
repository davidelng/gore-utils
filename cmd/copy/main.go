package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/davidelng/gore-utils/pkg/copy"
)

func main() {
	isVerbose := flag.Bool("v", false, "Display progress bar and detailed output")
	isInteractive := flag.Bool("i", false, "Prompt before overwriting")
	isDryRun := flag.Bool("d", false, "Test the process instead of executing the copy")
	bufSize := flag.Int64("b", 32, "Buffer size in kilobytes")

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: copy [-v] [-b 32] [-d] [-i] <source> <dest>\n")
		os.Exit(1)
	}

	_ = &copy.Config{
		Source:      args[0],
		Dest:        args[1],
		BufferSize:  *bufSize,
		Verbose:     *isVerbose,
		Interactive: *isInteractive,
		DryRun:      *isDryRun,
	}
}
