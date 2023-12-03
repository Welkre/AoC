package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func GetNumber(text string) string {
	var result string

	for i := 0; i < len(text); i++ {
		if _, err := strconv.Atoi(string(text[i])); err == nil {
			result += string(text[i])
			break
		}
	}

	for i := len(text)-1; i >= 0; i-- {
		if _, err := strconv.Atoi(string(text[i])); err == nil {
			result += string(text[i])
			break
		}
	}
	return result
}

func ConvertStringToNumber(text string) string {
	// Parse the given piece of text for a token that matches either 
	// the first written number or the first integer literal
	
	NUMBERS_AS_STRINGS := map[string]int{"zero": 0,"one": 1,"two": 2,"three": 3,"four": 4,"five": 5,"six": 6,"seven": 7,"eight": 8,"nine": 9}
	priority := make(map[string]int)

	textReturn := text

	for i := range NUMBERS_AS_STRINGS {
		index := strings.Index(text, i)
		if index == -1 {
			continue
		}
		priority[i] = index
	}

	keys := make([]string, 0 , len(priority))

	for key := range priority {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return priority[keys[i]] < priority[keys[j]]
	})

	for _, k := range keys{
		stringNum := fmt.Sprintf("%d", NUMBERS_AS_STRINGS[k])
		textReturn = strings.ReplaceAll(textReturn, k, stringNum)
	}

	return textReturn

}


func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	bufScanner := bufio.NewScanner(file)

	total := 0

	for bufScanner.Scan() {
		text := bufScanner.Text()
		convertedText := ConvertStringToNumber(text)

		num, err := strconv.Atoi(GetNumber(convertedText))
		if err != nil {
			fmt.Println(err)
		}

		total += num
	}
	fmt.Printf("The total calibration value is: %d\n", total)
}

