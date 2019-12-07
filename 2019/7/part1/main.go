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
	saveInstruction           = 3
	outputInstruction         = 4
	jumpIfTrueIntruction      = 5
	jumpIfFalseIntruction     = 6
	lessThanIntruction        = 7
	equalsInstruction         = 8
	haltInstruction           = 99
)

const (
	positionMode = iota
	immediateMode
)

type program struct {
	opsCodes            []int
	initialState        []int
	position            int
	firstParameter      int
	secondParameter     int
	memoryAddress       int
	output              []int
	modes               []int
	currentModePosition int
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
		p.initialState = make([]int, len(p.opsCodes))
		copy(p.initialState, p.opsCodes)
	}
	return nil
}

func (p *program) eval(inputValue []int) error {
	currentInputIndex := 0

	for p.hasNext() {
		opsCode := p.next()

		p.modes = make([]int, 0, 2)
		p.modes = append(p.modes, (opsCode/10000)%10)
		p.modes = append(p.modes, (opsCode/1000)%10)
		p.modes = append(p.modes, (opsCode/100)%10)
		p.currentModePosition = 2
		opsCode = (opsCode % 100)

		switch opsCode {
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

		case saveInstruction:
			position := p.next()
			err := p.store(inputValue[currentInputIndex], position)
			if err != nil {
				return err
			}
			currentInputIndex++
		case outputInstruction:
			position := p.next()

			value, err := p.getValue(position)
			if err != nil {
				return err
			}

			p.output = append(p.output, value)
		case jumpIfTrueIntruction, jumpIfFalseIntruction:
			position := p.next()

			value, err := p.getValue(position)
			if err != nil {
				return err
			}

			jump := false
			if opsCode == jumpIfTrueIntruction && value != 0 {
				jump = true
			}

			if opsCode == jumpIfFalseIntruction && value == 0 {
				jump = true
			}

			position = p.next()
			if jump {
				value, err := p.getValue(position)
				if err != nil {
					return err
				}
				p.position = value
			}
		case lessThanIntruction, equalsInstruction:
			position := p.next()

			firstValue, err := p.getValue(position)
			if err != nil {
				return err
			}

			position = p.next()

			secondValue, err := p.getValue(position)
			if err != nil {
				return err
			}

			result := 0

			if opsCode == lessThanIntruction && firstValue < secondValue {
				result = 1
			}

			if opsCode == equalsInstruction && firstValue == secondValue {
				result = 1
			}

			memoryAddress := p.next()
			err = p.store(result, memoryAddress)

			if err != nil {
				return nil
			}
		default:
			return fmt.Errorf("invalid opscode %d position: %d", opsCode, p.position)
		}
	}
	return nil
}

func (p *program) hasNext() bool {
	return p.position < len(p.opsCodes)
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

func (p *program) next() int {
	value := p.opsCodes[p.position]
	p.position++
	return value
}

func (p *program) nextMode() int {
	if len(p.modes) == 0 || p.currentModePosition == -1 {
		return positionMode
	}
	mode := p.modes[p.currentModePosition]
	p.currentModePosition--
	return mode
}

func (p *program) getValue(position int) (int, error) {
	if p.nextMode() == immediateMode {
		return position, nil
	}

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
func (p *program) reset() {
	p.position = 0
	p.memoryAddress = 0
	p.firstParameter = 0
	p.secondParameter = 0

	p.output = make([]int, 0)
	p.modes = make([]int, 0)
	p.currentModePosition = 0

	p.opsCodes = make([]int, len(p.initialState))
	copy(p.opsCodes, p.initialState)
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

func main() {
	inputPathPtr := flag.String("input-path", "input", "path to input file")
	inputValuePtr := flag.String("input-value", "1,0", "value to input to code")
	sequencePtr := flag.Bool("sequence", false, "run phase settingsr sequence of inputs")
	generatePtr := flag.Bool("generate", false, "generate sequence of inputs")
	flag.Parse()

	input, err := os.Open(*inputPathPtr)

	if err != nil {
		fmt.Printf("Failed to open input file %s with %v\n", *inputPathPtr, err)
		return
	}
	defer input.Close()

	inputsStr := strings.Split(*inputValuePtr, ",")
	inputsInt := make([]int, 0)

	for _, v := range inputsStr {
		value, err := strconv.Atoi(v)

		if err != nil {
			fmt.Printf("failed to converted input to integer with %s\n", err)
			return
		}
		inputsInt = append(inputsInt, value)
	}

	var (
		result  string
		output  []int
		execErr error
	)

	if !*sequencePtr && !*generatePtr {
		result, output, execErr = exec(input, inputsInt)
	}

	if *sequencePtr {
		p := &program{}
		err := p.parser(input)

		if err != nil {
			fmt.Printf("failed to parse opscode with %s\n", err)
			return
		}

		result, output, execErr = execSequence(p, inputsInt)
	}

	if *generatePtr {
		p := &program{}
		err := p.parser(input)

		if err != nil {
			fmt.Printf("failed to parse opscode with %s\n", err)
			return
		}

		result, output, execErr = generateMaxSequence(p, inputsInt)
	}

	if execErr != nil {
		fmt.Printf("Failed to exec opscode with %v\n", err)
		return
	}

	fmt.Println(result)
	for _, v := range output {
		fmt.Printf("Output: %d\n", v)
	}
}

func exec(input io.Reader, inputValue []int) (string, []int, error) {
	p := program{}
	err := p.parser(input)

	if err != nil {
		return "", []int{}, err
	}

	err = p.eval(inputValue)

	return p.String(), p.output, err
}

func generateMaxSequence(p *program, inputValue []int) (string, []int, error) {

	var (
		result    string
		output    []int
		maxOutput int
	)
	fmt.Println(inputValue)

	var n = len(inputValue) - 1
	var i, j int
	for c := 1; c < 120; c++ {

		i = n - 1
		j = n

		for inputValue[i] > inputValue[i+1] {
			i--
		}
		for inputValue[j] < inputValue[i] {
			j--
		}

		inputValue[i], inputValue[j] = inputValue[j], inputValue[i]
		j = n
		i += 1

		for i < j {
			inputValue[i], inputValue[j] = inputValue[j], inputValue[i]
			i++
			j--
		}

		r, o, err := execSequence(p, inputValue)

		if err != nil {
			return "", []int{}, err
		}

		if o[0] > maxOutput {
			maxOutput = o[0]
			result = r
			output = o
		}
	}

	return result, output, nil
}

func execSequence(p *program, inputValue []int) (string, []int, error) {
	var err error

	currentInputValue := 0
	for i := 0; i < len(inputValue); i++ {
		p.reset()
		err = p.eval([]int{inputValue[i], currentInputValue})

		if err != nil {
			return "", []int{}, err
		}
		currentInputValue = p.output[0]
	}

	return p.String(), p.output, err
}
