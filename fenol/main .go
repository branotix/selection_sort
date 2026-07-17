package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Selection Sort Algorithm
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
	// ১. ফাইল ওপেন এবং রিড করা
	file, err := os.Open("note.txt")
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

	fmt.Printf("Loaded %d numbers. Starting Selection Sort...\n", len(numbers))

	// ২. সিলেকশন সর্ট এবং সময় গণনা
	startTime := time.Now()
	selectionSort(numbers)
	duration := time.Since(startTime)

	fmt.Println("Sorting completed successfully!")
	fmt.Printf("Time taken: %v\n", duration)

	// ৩. সর্ট হওয়া পুরো ডাটা নতুন একটি ফাইলে সেভ করা
	outputFile, err := os.Create("sorted_50k.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush() // সব ডাটা ফাইলে রাইট হওয়া নিশ্চিত করা

	for _, num := range numbers {
		// int কে string এ কনভার্ট করে ফাইলে লেখা হচ্ছে
		_, err := writer.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("✅ Full sorted data has been successfully saved to 'sorted_50k.txt'. Open the file to check!")
}
