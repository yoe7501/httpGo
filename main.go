package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

var ErrBadUsage = errors.New("bad usage")

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		if errors.Is(err, ErrBadUsage) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func run() error {
	var verbose bool
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: [options] <city>\n\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(&verbose, "v", false, "more output information")

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return fmt.Errorf("%w: City name required", ErrBadUsage)
	}
	city := args[0]
	if verbose {
		fmt.Printf("Fetching weather for: %s with extra information !\n\n", city)
		return nil
	}
	fmt.Println("Fetching weather for:", city)

	return nil
}
