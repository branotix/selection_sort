package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
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
	quickSort(numbers, 0, len(numbers)-1)
	duration := time.Since(startTime)

	fmt.Printf("Sorting completed Time taken: %v\n", duration)
}
