package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"log"
)


func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close();

	scanner := bufio.NewScanner(file)

	fabric := [1000][1000]int{}

	for scanner.Scan() {
		line := scanner.Text()
		s := RegSplit(line)

		x := ToInt(s[1])
		y := ToInt(s[2])
		w := ToInt(s[3])
		h := ToInt(s[4])

		for i := y; i < y+h; i++ {
			for j := x; j < x+w; j++ {
				fabric[i][j] = fabric[i][j] + 1
			}
		}
	}

	sum := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if fabric[i][j] > 1 {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

func RegSplit(text string) []string {
	r, _ := regexp.Compile("[0-9]+")

	s := r.FindAllString(text, -1)
	return s
}

func ToInt(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
