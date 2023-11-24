package main

import (
	"fmt"
	"testing"
)

func TestExecuteDiff(t *testing.T) {
	file1 := []string{
		"Coding Challenges helps you become a better software engineer through that build real applications.",
		"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
		"I’ve used or am using these coding challenges as exercise to learn a new programming language or technology.",
		"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
	}
	file2 := []string{
		"Helping you become a better software engineer through coding challenges that build real applications.",
		"I share a weekly coding challenge aimed at helping software engineers level up their skills through deliberate practice.",
		"These are challenges that I’ve used or am using as exercises to learn a new programming language or technology.",
		"Each challenge will have you writing a full application or tool. Most of which will be based on real world tools and utilities.",
	}
	output := ExecuteDiff(file1, file2)
	fmt.Println(output)
}
