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
		count, _ := o.orbitCount(v)
		counter = counter + count
	}
	return counter
}

func (o orbitMap) orbitCount(name string) (int, []string) {
	objects := make([]string, 0)
	counter := 0
	for {
		objects = append(objects, name)
		counter++
		if name == "COM" {
			break
		}
		name = o[name]
	}
	return counter, objects
}

func (o orbitMap) orbitalTransferCount(object1, object2 string) int {

	objectOneDist, objectOneOrbits := o.orbitCount(object1)
	objectTwoDist, objectTwoOrbits := o.orbitCount(object2)

	pointOne := objectOneOrbits
	pointTwo := objectTwoOrbits

	if objectTwoDist > objectOneDist {
		pointOne = objectTwoOrbits
		pointTwo = objectOneOrbits
	}

	count := -1
	for i, v := range pointOne {
		for j, u := range pointTwo {
			if v == u {
				count = j
				break
			}
		}

		if count > -1 {
			count = i + count
			break
		}
	}

	// minus 2 here because we are counting both objects as well
	return count - 2
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
	orbitalTransfershPtr := flag.Bool("transfers", false, "Get count of orbital transfers")
	objectOnePtr := flag.String("object1", "YOU", "object1 name")
	objectTwoPtr := flag.String("object2", "SAN", "object1 name")
	flag.Parse()

	input, err := os.Open(*inputPathPtr)

	if err != nil {
		fmt.Printf("Failed to open input file %s with %v\n", *inputPathPtr, err)
		return
	}
	defer input.Close()
	o := New(input)

	if !*orbitalTransfershPtr {
		fmt.Println(o.countDirectAndIndirect())
		return
	}

	fmt.Println(o.orbitalTransferCount(*objectOnePtr, *objectTwoPtr))
}
