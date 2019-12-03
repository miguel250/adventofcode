package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	inputPathPtr := flag.String("input-path", "input", "path to input file")
	flag.Parse()

	input, err := os.Open(*inputPathPtr)

	if err != nil {
		fmt.Printf("Failed to open input file %s with %v\n", *inputPathPtr, err)
		return
	}

	defer input.Close()

	calculateFuel(input, os.Stdout)
}

func calculateFuel(input io.Reader, output io.Writer) error {
	scanner := bufio.NewScanner(input)
	sum := 0

	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return fmt.Errorf("failed to covered mass to an integer with: %w", err)
		}

		fuel := fuelEquation(mass)
		sum += fuel

		for fuel > 0 {
			fuel = fuelEquation(fuel)
			if fuel > 0 {
				sum += fuel
			}
		}

	}

	str := fmt.Sprintf("%d\n", sum) // convert to integer to round down
	output.Write([]byte(str))

	return nil
}

func fuelEquation(mass int) int {
	return int(mass/3) - 2
}
