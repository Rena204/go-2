package main

import "fmt"

func main() {
	sentence := "hello world"
	counts := make(map[rune]int)
	total := 0

	// Подсчет количества каждой буквы в предложении
	for _, char := range sentence {
		if char != ' ' {
			counts[char]++
			total++
		}
	}

	// Вывод результатов
	for char, count := range counts {
		percentage := float64(count) / float64(total) * 100
		fmt.Printf("%c - %d %.2f%%\n", char, count, percentage)
	}
}
