package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
	filename := os.Args[1]
	lines := GetLines(filename)

	line := lines[0]

	// part1
	line = GetUnitsAfterReaction(line)
	fmt.Println(len(line))
}

func GetLines(fileName string) []string {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func GetUnitsAfterReaction(line string) string {
	found := true

	for ; found ; {
		found = false
		for i := 0; i < len(line) - 1; i++ {
			if(line[i] == line[i+1] + 32 || line[i] == line[i+1] - 32 ){
				line = line[:i] + line[i+2:]
				found = true
			}
		}
	}

	return line;
}
