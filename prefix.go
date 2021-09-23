package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var separator string
	flag.StringVar(&separator, "separator", "", "separator between words")
	flag.StringVar(&separator, "s", "", "separator between words")
	flag.Parse()

	suffix_file := flag.Arg(0)
	if suffix_file == "" {
		fmt.Fprintln(os.Stderr, "usage: prefix [--separator=<char>] <suffix_file>")
		return
	}

	var prefix_f io.Reader
	var suffix_f io.Reader
	var err error

	prefix_f = os.Stdin

	suffix_f, err = os.Open(suffix_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open word list: %s\n", err)
		return
	}

	sc_suffix := bufio.NewScanner(suffix_f)
	sc_prefix := bufio.NewScanner(prefix_f)

	suffixs := make([]string, 0)

	for sc_suffix.Scan() {
		suffixs = append(suffixs, sc_suffix.Text())
	}

	prefixs := make(chan string)

	go func() {
		for prefix := range prefixs {
			for _, suffix := range suffixs {
				fmt.Println(prefix + separator + suffix)
			}
		}
	}()

	for sc_prefix.Scan() {
		prefixs <- sc_prefix.Text()
	}
}
