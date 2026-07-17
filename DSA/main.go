package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ১. মূল ফাইলটি ওপেন করুন
	inputFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// ২. নতুন ফাইলটি তৈরি করুন যেখানে প্রথম ৫০k ডাটা সেভ হবে
	outputFile, err := os.Create("250k_data.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush() // সব ডাটা ফাইলে রাইট হওয়া নিশ্চিত করবে

	count := 0
	target := 250000

	// ৩. লুপ চালিয়ে প্রথম ৫০,০০০ লাইন রিড এবং রাইট করুন
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

	fmt.Printf("Successfully saved %d lines to first_50k.txt\n", count)
}
