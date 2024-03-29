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

type pixels [][]int

type image struct {
	layers       []pixels
	wide         int
	tall         int
	minZeroLayer pixels
}

func (img *image) parse(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanRunes)
	img.layers = make([]pixels, 0)
	p := make(pixels, img.tall)

	currentTallIndex := 0
	currentWideIndex := 0
	layerTotalZeros := 0
	totalZeros := -1

	for scanner.Scan() {
		digit := strings.TrimSpace(scanner.Text())

		if digit == "" {
			break
		}
		pixelColor, err := strconv.Atoi(digit)
		if err != nil {
			return fmt.Errorf("failed to convert pixel to integer with %w", err)
		}

		p[currentTallIndex] = append(p[currentTallIndex], pixelColor)
		currentWideIndex++

		if pixelColor == 0 {
			layerTotalZeros++
		}

		if currentWideIndex >= img.wide {
			currentTallIndex++
			currentWideIndex = 0
		}

		if currentTallIndex >= img.tall {
			currentTallIndex = 0

			img.layers = append(img.layers, p)

			if totalZeros > layerTotalZeros || totalZeros < 0 {
				totalZeros = layerTotalZeros
				img.minZeroLayer = p
			}

			layerTotalZeros = 0
			p = make(pixels, img.tall)
		}
	}
	return nil
}

func (img *image) checkCorruption() int {
	oneCounter := 0
	twoCounter := 0

	for i := range img.minZeroLayer {
		for _, v := range img.minZeroLayer[i] {
			if v == 1 {
				oneCounter++
			}

			if v == 2 {
				twoCounter++
			}
		}
	}
	return oneCounter * twoCounter
}

func main() {
	inputPathPtr := flag.String("input-path", "input", "path to input file")
	widePtr := flag.Int("wide", 25, "image will")
	tallPtr := flag.Int("tall", 6, "image tall")
	flag.Parse()

	input, err := os.Open(*inputPathPtr)

	if err != nil {
		fmt.Printf("Failed to open input file %s with %v\n", *inputPathPtr, err)
		return
	}
	defer input.Close()

	img := &image{wide: *widePtr, tall: *tallPtr}
	err = img.parse(input)
	if err != nil {
		fmt.Printf("failed to converted input to integer with %s\n", err)
		return
	}

	fmt.Println(img.checkCorruption())
}
