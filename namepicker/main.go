package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		fmt.Println("Error: File is empty or improperly formatted.")
		return
	}
	runs, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid run count format:", scanner.Text())
		return
	}

	var names []string
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(names) == 0 {
		fmt.Println("No names found in the file.")
		return
	}

	selectionCounts := make(map[string]int)
	for i := 0; i < runs; i++ {
		selectedName := names[rand.Intn(len(names))]
		selectionCounts[selectedName]++
	}

	var mostSelectedName string
	maxCount := 0
	for name, count := range selectionCounts {
		if count > maxCount {
			mostSelectedName = name
			maxCount = count
		}
	}

	fmt.Println("Final selected name:", mostSelectedName)
}
