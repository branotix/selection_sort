package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()
	outputFile, err := os.Create("250k_data.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	count := 0
	target := 250000

	for scanner.Scan() {
		if count >= target {
			break
		}

		// প্রতি লাইনের ডাটা এবং একটি নতুন লাইন (\n) রাইট করা হচ্ছে
		_, err := writer.WriteString(scanner.Text() + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		count++
	}

	// কোনো রিডিং এরর আছে কিনা চেক করা
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for i := 0; i <= 10; i++ {
		fmt.Println("added")
	}
	fmt.Printf("Successfully saved %d lines to first_50k.txt\n", count)
}
