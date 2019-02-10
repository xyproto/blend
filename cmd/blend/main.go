package main

import (
	"fmt"
	"github.com/xyproto/mix"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Required arguments: infile1 infile2 outfile\n")
		os.Exit(1)
	}

	fn1 := os.Args[1]
	fn2 := os.Args[2]
	out := os.Args[3]

	fmt.Printf("Mixing %s and %s.\n", fn1, fn2)
	err := mix.Files(fn1, fn2, out, 0.5)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s\n", out)
}
