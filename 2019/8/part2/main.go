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

type pixels [][]int

type image struct {
	layers       []pixels
	wide         int
	tall         int
	minZeroLayer pixels
	topLayers    []pixels
}

func (img *image) String() string {
	var b bytes.Buffer
	for i := 0; i < len(img.layers); i++ {
		for j := 0; j < img.tall; j++ {
			for k := 0; k < img.wide; k++ {
				v := img.layers[i][j][k]
				if v == 1 {
					b.WriteString("█")
				}
				if v == 0 {
					b.WriteString(" ")
				}
			}
			b.WriteString("\n")
		}
	}
	return b.String()
}

func (img *image) parse(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(splitOnDigit)
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

func (img *image) decode() string {
	p := make(pixels, img.tall)

	for i := 0; i < img.wide; i++ {
		for j := 0; j < img.tall; j++ {
			for k := 0; k < len(img.layers); k++ {
				if img.layers[k][j][i] != 2 {
					p[j] = append(p[j], img.layers[k][j][i])
					break
				}
			}
		}
	}

	img.layers = []pixels{p}
	return img.String()
}

func splitOnDigit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	for i := 0; i < len(data); i++ {
		return i + 1, []byte{data[i]}, nil
	}
	if !atEOF {
		return 0, nil, nil
	}
	return 0, data, bufio.ErrFinalToken

}

func main() {
	inputPathPtr := flag.String("input-path", "input", "path to input file")
	widePtr := flag.Int("wide", 25, "image will")
	tallPtr := flag.Int("tall", 6, "image tall")
	decodePtr := flag.Bool("decode", true, "decode image")
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

	if *decodePtr {
		fmt.Print(img.decode())
	}
}
