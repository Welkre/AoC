package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func throwRock(action string) [2]int {
	switch action {
	case "X": return [2]int{2,3};
	case "Y": return [2]int{8, 4};
	case "Z": return [2]int{3, 8};
	}
	return [2]int{0, 0}
}
func throwPaper(action string) [2]int {
	switch action {
	case "X": return [2]int{1, 1};
	case "Y": return [2]int{5, 5};
	case "Z": return [2]int{9, 9};
	}
	return [2]int{0, 0}
}
func throwScissors(action string) [2]int {
	switch action {
	case "X": return [2]int{7, 2};
	case "Y": return [2]int{2, 6};
	case "Z": return [2]int{6, 7};
	}
	return [2]int{0, 0}
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	score := 0
	realScore := 0
	for scanner.Scan() {
		text := scanner.Text()

		splitText := strings.Fields(text)
		switch splitText[0] {
		case "A": {
			out := throwRock(splitText[1])
			score += out[0]
			realScore += out[1]
			}
		case "B": {
			out := throwPaper(splitText[1])
			score += out[0]
			realScore += out[1]
			}
		case "C": {
			out := throwScissors(splitText[1])
			score += out[0]
			realScore += out[1]
			}
		}
	}
	fmt.Printf("Original Score: %d\n", score)
	fmt.Printf("Real Score: %d\n", realScore)

}
