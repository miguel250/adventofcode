package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"
)

type point struct {
	x, y int
}

type wire struct {
	points    []point
	lastPoint point
}

func (w *wire) Step(direction string) error {
	p := point{}
	switch direction {
	case RIGHT:
		p.x = w.lastPoint.x + 1
		p.y = w.lastPoint.y
	case LEFT:
		p.x = w.lastPoint.x - 1
		p.y = w.lastPoint.y
	case UP:
		p.x = w.lastPoint.x
		p.y = w.lastPoint.y + 1
	case DOWN:
		p.x = w.lastPoint.x
		p.y = w.lastPoint.y - 1
	default:
		return fmt.Errorf("unknow direction %s", direction)
	}

	w.points = append(w.points, p)
	w.lastPoint = p
	return nil
}

type Panel struct {
	intersectionPoints []point
}

func (p *Panel) process(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	firstWire := &wire{}
	firstWireProcessed := false
	lastPoint := point{}
	for scanner.Scan() {

		values := strings.TrimSpace(scanner.Text())
		actions := strings.Split(values, ",")

		for _, action := range actions {
			w := &wire{lastPoint: lastPoint}
			direction := string(action[0])
			move, err := strconv.Atoi(action[1:])

			if err != nil {
				return fmt.Errorf("unable to process action %s", action)
			}

			for i := 0; i < move; i++ {
				if firstWireProcessed {
					w.Step(direction)
					lastPoint = w.lastPoint
					continue
				}
				firstWire.Step(direction)
			}
			if firstWireProcessed {
				p.intersections(firstWire, w)
			}
		}
		firstWireProcessed = true
	}

	return nil
}

func (p *Panel) shortersDistance() int {
	distance := 0

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for _, point := range p.intersectionPoints {
		d := abs(point.x) + abs(point.y)
		if distance == 0 || d < distance {
			distance = d
		}
	}
	return distance
}

func (p *Panel) intersections(wireOnePoints, wireTwoPoints *wire) {
	for _, wireOnePoint := range wireOnePoints.points {
		if wireOnePoint.x == 0 && wireOnePoint.y == 0 {
			continue
		}
		for _, wireTwoPoint := range wireTwoPoints.points {
			if wireTwoPoint.x == 0 && wireTwoPoint.y == 0 {
				continue
			}
			if wireOnePoint.x == wireTwoPoint.x && wireOnePoint.y == wireTwoPoint.y {
				p.intersectionPoints = append(p.intersectionPoints, point{wireTwoPoint.x, wireOnePoint.y})
				break
			}
		}
	}
}

func New(input io.Reader) (*Panel, error) {
	p := &Panel{}
	err := p.process(input)
	return p, err
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

	p, err := New(input)

	if err != nil {
		fmt.Printf("Failed to create panel with %s\n", err)
	}

	fmt.Println(p.shortersDistance())

}
