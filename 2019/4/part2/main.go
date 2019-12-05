package main

import (
	"flag"
	"fmt"
)

func main() {
	minPtr := flag.Int("min", 130254, "minimum starting range for password")
	maxPtr := flag.Int("max", 678275, "Max ending range for password")
	flag.Parse()

	result, err := validPasswordCount(*minPtr, *maxPtr)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func validPasswordCount(min, max int) (int, error) {
	if min <= 100000 || max >= 999999 {
		return 0, fmt.Errorf("min (%d) and Max (%d) need to be between 100000 and 999999", min, max)
	}
	validPasswords := 0
	for i := min; i < max; i++ {
		if isValidPassword(i) {
			validPasswords++
		}
	}
	return validPasswords, nil
}

func isValidPassword(password int) bool {

	if password%10 < (password/100000)%10 {
		return false
	}
	lastDigit := 0
	doubles := make(map[int]int)
	for n := 100000; n >= 1; n = n / 10 {
		digit := (password / n) % 10

		if lastDigit == 0 {
			lastDigit = digit
			continue
		}

		if digit < lastDigit {
			return false
		}

		if digit == lastDigit {
			doubles[digit]++
		}

		lastDigit = digit
	}

	for _, v := range doubles {
		if v == 1 {
			return true
		}
	}
	return false
}
