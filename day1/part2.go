package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sum := 0
	var sums []int

	for {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			sums = append(sums, sum)

			line := scanner.Text()
			i, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal(err)
			}

			sum += i

			for _, v := range sums {
				if v == sum {
					fmt.Println(sum)
					return
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
