package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// GenerateIBAN generates a valid IBAN number for the given UK bank account details
func GenerateIBAN(countryCode, bankCode, accountNumber string) (string, error) {
	// Step 1: Move the first 4 characters to the end
	iban := accountNumber + countryCode + bankCode + "00"

	// Step 2: Convert letters to numbers (A=10, B=11, ..., Z=35)
	for i := 0; i < len(iban); i++ {
		if iban[i] >= 'A' && iban[i] <= 'Z' {
			iban = iban[:i] + string(iban[i]-'A'+'0'+10) + iban[i+1:]
		}
	}

	// Step 3: Calculate the modulus 97 of the updated IBAN
	remainder := mod97(iban)

	// Step 4: Subtract the remainder from 98 and pad with leading zeros
	checkDigits := fmt.Sprintf("%02d", 98-remainder)

	// Step 5: Generate the final IBAN number
	iban = countryCode + checkDigits + bankCode + accountNumber

	return iban, nil
}

// mod97 performs modulo 97 operation on the given number string
func mod97(number string) int {
	const modValue = 97
	parsedInt := new(big.Int)
	parsedInt.SetString(number, 10)
	remainder := new(big.Int)
	remainder.Mod(parsedInt, big.NewInt(modValue))

	return int(remainder.Int64())
}

// GenerateLuhnCheckDigit generates a Luhn check digit for the given number
func GenerateLuhnCheckDigit(number string) int {
	sum := 0
	double := false

	// Iterate from right to left, doubling every second digit
	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	// Calculate the check digit as the smallest number that, when added to the sum, makes it divisible by 10
	checkDigit := (10 - (sum % 10)) % 10
	return checkDigit
}

// GenerateAccountNumbers generates a slice of unique random account numbers within the specified range
func GenerateAccountNumbers(min, max, count int) ([]int, error) {
	if min > max || count > (max-min+1) {
		return nil, fmt.Errorf("invalid account number range or count")
	}

	numbers := make([]int, count)

	// Generate account numbers using Fisher-Yates shuffle algorithm
	for i := 0; i < count; i++ {
		accountNumber := min + i
		numbers[i] = accountNumber
	}
	// for i := min; i <= max; i++ {
	// 	numbers[i-min] = strconv.Itoa(i)
	// }

	rand.Seed(time.Now().UnixNano())

	// Shuffle the account numbers
	for i := len(numbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	// Append Luhn check digits
	// for i := range numbers {
	// 	accountNumber := numbers[i]
	// 	checkDigit := GenerateLuhnCheckDigit(accountNumber)
	// 	accountNumber += strconv.Itoa(checkDigit)
	// 	numbers[i] = accountNumber
	// }

	return numbers, nil
}

func main() {
	// Example account details
	// countryCode := "GB"
	// bankCode := "BARC"
	minAccountNumber := 3000000
	maxAccountNumber := 4000000
	numAccounts := 1000000

	fmt.Println("Generating", numAccounts, "account numbers...")

	accountNumbers, err := GenerateAccountNumbers(minAccountNumber, maxAccountNumber, numAccounts)
	if err != nil {
		fmt.Println("Error generating account numbers:", err)
		return
	}

	file, err := os.Create("numbers.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, accountNumber := range accountNumbers {
		// Generate IBAN number
		// iban, err := GenerateIBAN(countryCode, bankCode, accountNumber)
		// if err != nil {
		// 	fmt.Println("Error generating IBAN:", err)
		// 	return
		// }

		_, err := writer.WriteString(fmt.Sprintf("%d\n", accountNumber))
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println("Account Number:", accountNumber)
		// fmt.Println("IBAN:", iban)
		// fmt.Println()
	}
}
