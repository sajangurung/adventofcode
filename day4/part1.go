package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, _ := os.Open(filename)
	defer file.Close()

	// Sort lines
	lines := Sort(file)

	totals := make(map[int] int)
	minutes := make(map[int] []int)

	id := 0
	start := 0

	for _, line := range lines {
		split := RegSplit(line, "[# :\\]]+")

		minute, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal(err)
		}
		action := split[3][0]
		if action == 'G' {
			tempId, err := strconv.Atoi(split[4])
			if err != nil {
				log.Fatal(err)
			}
			id = tempId

			if _, ok := minutes[id]; !ok {
				minutes[id] = make([]int, 60)
			}

		} else if action == 'f' {
			start = minute
		} else if action == 'w' {
			end := minute
			temp := minutes[id]
			for i := start; i < end; i++ {
				 temp[i] = temp[i] + 1
				 totals[id] ++
			}
			minutes[id] = temp
		}
	}

	guardId := MaxId(totals)

	guardMinutes := minutes[guardId]

	maxMinutesAsleep := MaxMinutesAsleep(guardMinutes)

	fmt.Println(guardId,maxMinutesAsleep, guardId * maxMinutesAsleep)
}

func RegSplit(text string, delimeter string) []string {
	reg := regexp.MustCompile(delimeter)
	indexes := reg.FindAllStringIndex(text, -1)
	laststart := 0
	result := make([]string, len(indexes)+1)
	for i, element := range indexes {
		result[i] = text[laststart:element[0]]
		laststart = element[1]
	}
	result[len(indexes)] = text[laststart:len(text)]
	return result
}

func MaxId(arr map[int] int) int {
	maxValue := 0
	Key := 0

	for k,v := range arr {
		if v > maxValue {
			maxValue = v
			Key = k
		}
	}

	return Key
}

func MaxMinutesAsleep(arr []int) int {
	maxValue := 0
	Key := 0

	for k,v := range arr {
		if v > maxValue {
			maxValue = v
			Key = k
		}
	}

	return Key
}

func Sort(file *os.File) []string {
	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err:= scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Strings(lines)

	return lines
}
