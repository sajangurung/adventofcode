package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		t := fmt.Sprintf("sum = %d + %d", sum, i)
		fmt.Println(t)
		sum = sum + i
	}

	fmt.Println(sum)

	if err:= scanner.Err(); err != nil {
		log.Fatal(err)
	}


}
