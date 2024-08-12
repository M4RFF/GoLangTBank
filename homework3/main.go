package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Options struct {
	From      string
	To        string
	Offset    int64
	Limit     int64
	BlockSize int64
	Conv      string
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.From, "from", "", "file to read. by default - stdin")
	flag.StringVar(&opts.To, "to", "", "file to write. by default - stdout")
	flag.Int64Var(&opts.Offset, "offset", 0, "number of bytes to skip from the input")
	flag.Int64Var(&opts.Limit, "limit", 0, "number of bytes to read. by default -offset")
	// 4096 bytes is the standard memory page size, optimal for Input/Output operations
	flag.Int64Var(&opts.BlockSize, "block-size", 4096, "block in bytes for reading and writing")
	flag.StringVar(&opts.Conv, "conv", "", "text transformations to: upper_case, lower_case, trim_spacesx")

	flag.Parse()

	if strings.Contains(opts.Conv, "upper_case") && strings.Contains(opts.Conv, "lower_case") {
		return nil, errors.New("can not use upper_case and lower_case transformations")
	}

	return &opts, nil
}

func transformData(data []byte, conv string) []byte {
	text := string(data)

	switch {
	case strings.Contains(conv, "upper_case"):
		text = strings.ToUpper(text)
	case strings.Contains(conv, "lower_case"):
		text = strings.ToLower(text)
	case strings.Contains(conv, "trim_spacesx"):
		text = strings.TrimSpace(text)
	}

	return []byte(text)
}

func main() {
	opts, err := ParseFlags()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)
		os.Exit(1)
	}

	fmt.Println(opts)

	var input io.Reader
	var output io.Writer

	// Input
	if opts.From != "" { // check if a user has provided a path to an input file
		file, err := os.Open(opts.From) // open the file for reading and put the result to the "file"
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "can not open input file", err) // print out an error
			os.Exit(1)                                                     // means that there's an error in the code and it stops immediately execution
		}
		defer file.Close()
		input = file // assigns input to file
	} else {
		input = os.Stdin // assigns input toos.Stdin
		if opts.Offset > 0 {
			fmt.Fprintln(os.Stderr, "can not seek in stdin when offset is specified")
			os.Exit(1)
		}
	}

	// Output
	if opts.To != "" { // check if a user has provided a path to an output file
		file, err := os.Create(opts.To) // create or open the file for writing and put the result to the "file"
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "can not open output file", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file // assigns output to file
	} else {
		output = os.Stdout // assigns output to os.Stdout
	}

	if opts.Offset > 0 {
		if _, err := input.(io.Seeker).Seek(opts.Offset, io.SeekStart); err != nil {
			fmt.Fprintln(os.Stderr, "cannot seek input:", err)
			os.Exit(1)
		}
	}

	// creates a limited reader
	reader := io.LimitReader(input, opts.Limit) // limits the number of bytes that can be read from input. maximum number of bytes to read equal to opts.Limit
	// creates a buffer for reading
	buffer := make([]byte, opts.BlockSize) // there's an array of bytes that holds the data from input and opts.BlockSize says how many bytes will be read

	for {
		n, err := reader.Read(buffer) // reads data from the buffer, n is a number of bytes
		if n > 0 {
			data := transformData(buffer[:n], opts.Conv)  //takes the data and then read it into buffer and then applies transformation
			if _, err := output.Write(data); err != nil { // write the transformed data to the output
				fmt.Fprintln(os.Stderr, "can not write to output:", err)
				os.Exit(1)
			}
		}
		if err != nil {
			if err != io.EOF { // if the end of the file is reached, then in terminal will be able to see an error
				fmt.Fprintln(os.Stderr, "error reading input:", err)
			}
			break // we leave the loop in two opportunities: the end of the file is reached or error
		}
	}
}
