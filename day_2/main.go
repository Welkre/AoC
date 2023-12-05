package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	}
	return false
}

func FindGameNumber(text string) int {
	if IsNumber(string(text[6])) {
		if IsNumber(string(text[7])) {
			strNum := fmt.Sprintf("%s%s%s", string(text[5]), string(text[6]), string(text[7]))
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println(err)
			}
			return num
		}
		strNum := fmt.Sprintf("%s%s", string(text[5]), string(text[6]))
		num, err := strconv.Atoi(strNum)
		if err != nil {
			fmt.Println(err)
		}
		return num
	}
	num, err := strconv.Atoi(string(text[5]))
	if err != nil {
		fmt.Println(err)
	}
	return num
}

func ParseGameData(text string) bool {
	fullGameData := strings.SplitAfter(text, ":")
	gamesData := strings.Split(fullGameData[1], ";")
	green := 0
	blue := 0
	red := 0

	for i := 0; i < len(gamesData); i++ {
		cubes := gamesData[i]
		// fmt.Println(cubes)
		for j := 0; j < len(cubes); j++ {
			switch string(cubes[j]) {
			case "g": {
				if GetCubeNumber(string(cubes[j-3:j-1])) > green { 
					green = GetCubeNumber(string(cubes[j-3:j-1]))
					}
				}
			case "b": {
				if GetCubeNumber(string(cubes[j-3:j-1])) > blue { 
					blue = GetCubeNumber(string(cubes[j-3:j-1]))
					}
				}
			case "r": {
				if string(cubes[j+2]) == "d" {
					if GetCubeNumber(string(cubes[j-3:j-1])) > red { 
						red = GetCubeNumber(string(cubes[j-3:j-1]))
						}
					}
				}
			default: continue
			}
		}
	}
	// fmt.Printf("Green: %d Blue: %d Red: %d \n", green, blue, red)
	if red > 12 {
		return false
	}
	if blue > 14 {
		return false
	}
	if green > 13 {
		return false
	}
	return true
}

func GetCubeNumber(s string) int {
	if IsNumber(string(s[0])) {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return num
	}
	num, err := strconv.Atoi(string(s[1]))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return num
	
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(file)
	partOneRes := 0

	for fileScanner.Scan() {
		gameNumber := FindGameNumber(fileScanner.Text())

		// fmt.Println(fileScanner.Text())

		if ParseGameData(fileScanner.Text()) {
			partOneRes += gameNumber
		}


	}
	fmt.Println(partOneRes)
}
