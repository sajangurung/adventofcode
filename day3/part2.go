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
	claims := [1338] bool{}
	for scanner.Scan() {
		line := scanner.Text()
		s := RegSplit(line)

		id := ToInt(s[0])
		x := ToInt(s[1])
		y := ToInt(s[2])
		w := ToInt(s[3])
		h := ToInt(s[4])

		for i := y; i < y+h; i++ {
			for j := x; j < x+w; j++ {
				if fabric[i][j] > 0 {
					claims[fabric[i][j]] = true
					claims[id] = true
				} else {
					fabric[i][j] = id
				}
			}
		}
	}

	for i := 1; i <= 1337; i++ {
		if !claims[i] {
			fmt.Println(i)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
