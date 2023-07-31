package main

import "fmt"

var sortCodes = map[string][]string{
	"emi": {"041317", "041319", "658370"},
	"crr": {"123456", "456789", "123789"},
}

func main() {
	code := "1234561"
	platform := "emi1"
	fmt.Println("Key: ", getKeyByValue(sortCodes, code))
	fmt.Println("Value: ", getValueByKey(sortCodes, platform))
}

func getKeyByValue(m map[string][]string, targetValue string) string {
	for key, values := range m {
		for _, value := range values {
			if value == targetValue {
				return key
			}
		}
	}
	return ""
}

func getValueByKey(m map[string][]string, key string) []string {
	return m[key]
}
