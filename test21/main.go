package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Increment  = 679186565
	Modulus    = 1073741789
	Multiplier = 921746065

	maxGenerateAttempts = 5
	rangeMin            = 3000000000
	rangeMax            = 5000000000

	fileName = "numbers.txt"
)

func main() {

	var genAccountNumber int64 = 3956675448 //3334600767 //3386627187

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a buffered writer to write the numbers to the file.
	writer := bufio.NewWriter(file)

	for i := 1; i <= 1000; i++ {

		//fmt.Printf("Last Account number: %d\n", genAccountNumber)

		//for attempts := 1; attempts <= maxGenerateAttempts; attempts++ {

		accountNumber := nextInSequence(rangeMin, genAccountNumber)

		_, err := writer.WriteString(fmt.Sprintf("%d\n", accountNumber))
		if err != nil {
			fmt.Println(err)
			return
		}

		// if accountNumber >= rangeMin && accountNumber <= rangeMax {
		// 	fmt.Println(accountNumber)
		// } else {
		// 	fmt.Printf("%d Retry in %d.%d\n", accountNumber, i, attempts)
		// }

		genAccountNumber = accountNumber
		//}
		//fmt.Println()
	}

	// Flush the buffer to ensure all data is written to the file.
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Successfully wrote generated numbers to %s.", fileName)
}

func nextInSequence(rangeMin, number int64) int64 {
	last := number - rangeMin
	return rangeMin + (Multiplier*last+Increment)%Modulus
}
