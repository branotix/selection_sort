package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func main() {

	file, err := os.Open("./Sorted_Data/sorted_250k.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err == nil {
			numbers = append(numbers, val)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	startTime := time.Now()
	selectionSort(numbers)
	duration := time.Since(startTime)

	fmt.Printf("Sorting completed Time taken: %v\n", duration)

}

// outputFile, err := os.Create("sorted_250k.txt")
// 	if err != nil {
// 		fmt.Println("Error creating output file:", err)
// 		return
// 	}
// 	defer outputFile.Close()

// 	writer := bufio.NewWriter(outputFile)
// 	defer writer.Flush()

// 	for _, num := range numbers {
// 		_, err := writer.WriteString(strconv.Itoa(num) + "\n")
// 		if err != nil {
// 			fmt.Println("Error writing to file:", err)
// 			return
// 		}
// 	}

// 	fmt.Println("successed")
