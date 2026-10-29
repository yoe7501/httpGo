package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: please provide a city")
		os.Exit(2)
	}
}

func run() error {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: [options] <city>\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return errors.New("please provide a city")
	}
	city := args[0]

	fmt.Println("Fetching weather for:", city)
	return nil
}
