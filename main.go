// Mark command line tool
// Available at http://github.com/a8m/mark
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/a8m/mark"
	"io"
	"os"
)

var (
	input     = flag.String("i", "", "")
	output    = flag.String("o", "", "")
	smarty    = flag.Bool("smartypants", false, "")
	fractions = flag.Bool("fractions", false, "")
)

var usage = `Usage: mark-cli [options...] <input>

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
	// Reader
	var reader *bufio.Reader
	if *input != "" {
		file, err := os.Open(*input)
		if err != nil {
			usageAndExit(fmt.Sprintf("Error to open file input: %s.", *input))
		}
		defer file.Close()
		reader = bufio.NewReader(file)
	} else {
		if stat, err := os.Stdin.Stat(); err == nil && stat.Size() > 0 {
			reader = bufio.NewReader(os.Stdin)
		} else {
			usageAndExit("")
		}
	}
	// Collect data
	var data string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			usageAndExit("Failed to reading input.")
		}
		data += line
	}
	// Writer
	var file *os.File
	var err error
	if *output != "" {
		file, err = os.Create(*output)
		if err != nil {
			usageAndExit("Error to create the wanted output file.")
		}
	} else {
		file = os.Stdout
	}
	// Mark rendering
	opts := mark.DefaultOptions()
	opts.Smartypants = *smarty
	opts.Fractions = *fractions
	m := mark.New(data, opts)
	if _, err := file.WriteString(m.Render()); err != nil {
		filename := *output
		if filename == "" {
			filename = "STDOUT"
		}
		usageAndExit(fmt.Sprintf("Error writing output to: %s.", filename))
	}
}

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprintf(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}
