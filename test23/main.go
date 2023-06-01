package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var SortCode, BIC string = "608387", "CNFVGB21XXX"
var AccountNumberLength, CoronaAccountNumber int = 8, 1018

func main() {
	checkDigit := "00"
	countryCode := BIC[4:6]
	accountNumber := fmt.Sprintf("%08d", CoronaAccountNumber)
	bankCode := BIC[0:4]

	draftIBAN := fmt.Sprintf("%s%s%s%s%s", bankCode, SortCode, accountNumber, countryCode, checkDigit)
	checkDigit = GenerateIBANChecksumDigit(draftIBAN)
	IBAN := fmt.Sprintf("%s%s%s%s%s", countryCode, checkDigit, bankCode, SortCode, accountNumber)
	fmt.Println(draftIBAN)
	fmt.Println(IBAN)
}

func GenerateIBANChecksumDigit(draftIBAN string) string {
	checkDigit := ""

	if draftIBAN != "" {
		// Step - 1: Replace each letter in the draft IBAN with two digits. A is replaced with 10, B = 11, C = 12 through to Z = 35
		numericIBAN := regexp.MustCompile("\\D").ReplaceAllStringFunc(strings.ToUpper(draftIBAN), func(letter string) string {
			return strconv.Itoa(int(letter[0]) - 55)
		})

		// Step - 2: Perform MOD97 on numeric encoded IBAN
		remainder := ""
		value9 := 9
		value97 := 97
		for _, digit := range numericIBAN {
			if len(remainder) < value9 {
				remainder += string(digit)
			} else {
				remainder = strconv.Itoa((toInt(remainder) % value97)) + string(digit)
			}
		}

		output := toInt(remainder) % value97

		// Step - 3: check digit = 98 - MOD97
		value98 := 98
		value2 := 2
		checkDigit = strconv.Itoa(value98 - output)
		checkDigit = strings.Repeat("0", value2-len(checkDigit)) + checkDigit
	}

	return checkDigit
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
