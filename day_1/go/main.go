package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	most := 0
	mostSecond := 0
	mostThird := 0
	curr := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		num, _ := strconv.Atoi(text)
		
		switch text {
		case "": curr = 0;
		default: curr += num; 
		}
		
		if curr >= most {
			mostThird = mostSecond
			mostSecond = most
			most = curr
		}
	}

	final := most + mostSecond + mostThird

	fmt.Println(final)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
