package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strAccountNumber string = "6083870000101"
	var strWeightNumber string = "2121212121212"

	var intMode int = 11
	var luhnCheckDigit int = 0

	if len(strAccountNumber) == len(strWeightNumber) {
		var totalWeight int = 0

		for i := 1; i <= len(strAccountNumber); i++ {

			j, _ := strconv.Atoi(strAccountNumber[i-1 : i])
			k, _ := strconv.Atoi(strWeightNumber[i-1 : i])

			totalWeight = totalWeight + (j * k)

			fmt.Printf("\nTotal Weight: %d", totalWeight)
			fmt.Printf("	j: %d k: %d", j, k)
		}

		var remainder int = totalWeight % intMode

		if remainder == 0 {
			luhnCheckDigit = remainder
		} else {
			luhnCheckDigit = intMode - remainder

			if luhnCheckDigit > 9 {
				luhnCheckDigit = -1
			}
		}
	}

	fmt.Printf("\n\nLuhn check digit: %d", luhnCheckDigit)
}
