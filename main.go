package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	input     = flag.String("i", "", "")
	output    = flag.String("u", "", "")
	smarty    = flag.Bool("smartypants", false, "")
	fractions = flag.Bool("fractions", false, "")
)

var usage = `Usage: mark-tool [options...] <input>

Options:
  -i  Specify file input, otherwise use last argument as input file. 
      If no input file is specified, read from stdin.
  -o  Specify file output. If none is specified, write to stdout.

  -smartypants  Use "smart" typograhic punctuation for things like 
                quotes and dashes.
  -fractions    Traslate fraction like to suitable HTML elements
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(usage))
	}
	flag.Parse()
	// Read input
	var reader *bufio.Reader
	if *input != "" {
		file, err := os.Open(*input)
		if err != nil {
			// Error and exit
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		reader = bufio.NewReader(os.Stdin)
	}
	var data string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			// Error and exit
		}
		data += line
	}
	// Write output
	var file *os.File
	var err error
	if *output != "" {
		file, err = os.Create(*output)
		if err != nil {
			// Error and exit
		}
	} else {
		file = os.Stdout
	}
	file.WriteString(data)
	fmt.Println(*input, *output, *smarty, *fractions)
}
