package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	example2()

}

func example1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	if input.Err() != nil {
		fmt.Println("Error scanning: ", input.Err().Error())
	}
	fmt.Printf("scanning")
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Printf("scanned")

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func example2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "example 2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n >= 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fmt.Printf("Dupilcate name: %s in file %s\n", input.Text(), f.Name())
		}
	}
}

func example3() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "example3: %v\n", err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
