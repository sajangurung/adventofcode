package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	countTwo := 0
	countThree := 0

	for scanner.Scan() {
		// Create a map of characters(int32) with their count value
		m := make(map[rune] int)
		line := scanner.Text()
		for _, char := range line {
			m[char] = m[char] + 1
		}

		twoExists := false
		threeExists := false
		for _, count := range m {
			if count == 2 && !twoExists {
				countTwo++
				twoExists = true
			} else if count == 3 && !threeExists {
				countThree++
				threeExists = true
			}
		}
	}

	checksum := countTwo * countThree
	fmt.Println(countTwo)
	fmt.Println(countThree)
	fmt.Println(checksum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


}
