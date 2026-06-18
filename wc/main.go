package main 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)	

	var flag string

	if len(os.Args) > 1 {
		flag = os.Args[1]
	}

	switch flag {
		case "-l": 
			count := 0

			for scanner.Scan() {
				count++
			}

			fmt.Println(count, os.Stdin)
			return 
		case "-w": 
			count := 0 
			scanner.Split(bufio.ScanWords)

			for scanner.Scan() {
				count++
			}	
			fmt.Println(count, os.Stdin)
			return 
		case "-m": 
			count := 0 
			scanner.Split(bufio.ScanRunes)

			for scanner.Scan() {
				count++
			}
			fmt.Println(count, os.Stdin)
			return
		default:	
			input := []string{}

			for scanner.Scan() {
				input = append(input, scanner.Text())
			}
			lines := len(input)
			words := countWords(input)
			characters := countCharacters(input)
			//bytes := countBytes(input)

			fmt.Println(lines, words, characters)
			//fmt.Println("input:", input)
	}
}



func countWords(input []string) int {
	count  := 0 

	for _, line := range input {
		scanner := bufio.NewScanner(strings.NewReader(line))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			count++
		}
	}
	return count 
}

func countCharacters(input []string) int {
	count := 0 
	for _, line := range input {
		count += len(line)
	}
	return count
}


