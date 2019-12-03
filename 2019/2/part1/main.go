package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	additionInstruction       = 1
	multiplicationInstruction = 2
	haltInstruction           = 99
)

type program struct {
	opsCodes        []int
	position        int
	firstParameter  int
	secondParameter int
	memoryAddress   int
}

func (p *program) parser(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(splitOnComma)

	p.opsCodes = make([]int, 0, 120)

	for scanner.Scan() {
		opsCode, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err != nil {
			return fmt.Errorf("failed to covered opsCode to an integer with: %w", err)
		}

		p.opsCodes = append(p.opsCodes, opsCode)
	}
	return nil
}

func (p *program) eval() error {
	for p.hasNext() {
		switch opsCode := p.next(); opsCode {
		case haltInstruction:
			return nil
		case additionInstruction:
			err := p.processOpsCode()
			if err != nil {
				return err
			}

			result := p.firstParameter + p.secondParameter
			err = p.store(result, p.memoryAddress)

			if err != nil {
				return err
			}

		case multiplicationInstruction:
			err := p.processOpsCode()

			if err != nil {
				return err
			}

			result := p.firstParameter * p.secondParameter
			err = p.store(result, p.memoryAddress)

			if err != nil {
				return err
			}

		default:
			return fmt.Errorf("invalid opscode %d", opsCode)
		}
	}
	return nil
}

func (p *program) hasNext() bool {
	return p.position < len(p.opsCodes)
}

func (p *program) next() int {
	value := p.opsCodes[p.position]
	p.position++
	return value
}

func (p *program) processOpsCode() error {
	firstPosition := p.next()
	secondPosition := p.next()
	p.memoryAddress = p.next()

	firstParameter, err := p.getValue(firstPosition)

	if err != nil {
		return err
	}
	p.firstParameter = firstParameter

	secondParameter, err := p.getValue(secondPosition)

	if err != nil {
		return err
	}

	p.secondParameter = secondParameter
	return nil
}

func (p *program) getValue(position int) (int, error) {
	if err := p.isValidPosition(position); err != nil {
		return 0, err
	}
	return p.opsCodes[position], nil
}

func (p *program) store(value, position int) error {
	if err := p.isValidPosition(position); err != nil {
		return err
	}
	p.opsCodes[position] = value
	return nil
}

func (p *program) isValidPosition(position int) error {
	if position > len(p.opsCodes) || position < 0 {
		return fmt.Errorf("failed to store value in position %d", position)
	}
	return nil
}

func (p *program) String() string {
	var b bytes.Buffer
	size := len(p.opsCodes)

	for i, v := range p.opsCodes {
		fmt.Fprintf(&b, "%d", v)
		if size-1 > i {
			fmt.Fprint(&b, ",")
		}
	}
	return b.String()
}

func splitOnComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			return i + 1, data[:i], nil
		}
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken

}

func exec(input io.Reader) (string, error) {
	p := program{}
	err := p.parser(input)

	if err != nil {
		return "", err
	}

	err = p.eval()
	return p.String(), err
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

	result, err := exec(input)

	if err != nil {
		fmt.Printf("Failed to exec opscode with %v\n", err)
		return
	}

	fmt.Println(result)
}
