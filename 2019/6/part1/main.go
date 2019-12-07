package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type orbitMap map[string]string

func (o orbitMap) countDirectAndIndirect() int {
	counter := 0

	for _, v := range o {
		key := v

		for {
			counter++
			if key == "COM" {
				break
			}
			key = o[key]
		}
	}
	return counter
}

func New(input io.Reader) orbitMap {
	o := make(orbitMap)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		orbit := strings.Split(scanner.Text(), ")")
		o[orbit[1]] = orbit[0]
	}
	return o
}

func main() {
	inputPathPtr := flag.String("input-path", "input", "path to input file")
	flag.Parse()

	input, err := os.Open(*inputPathPtr)

	if err != nil {
		fmt.Printf("Failed to open input file %s with %v\n", *inputPathPtr, err)
		return
	}
	defer input.Close()
	o := New(input)

	fmt.Println(o.countDirectAndIndirect())
}
