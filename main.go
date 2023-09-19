package main

import (
	"fmt"
	"strings"
)

func main() {
	result := AlayGen("hello", "world", "zzz")
	fmt.Println(result)

	n := 41
	fibResult := Fibonacci(n)
	fmt.Printf("Fibonacci ke-%d: %d\n", n, fibResult)
}

// Fungsi untuk mengubah huruf-huruf menjadi alay
func convertToAlay(word string) string {
	// Mengganti karakter-karakter tertentu menjadi alay
	word = strings.ReplaceAll(word, "a", "4")
	word = strings.ReplaceAll(word, "e", "3")
	word = strings.ReplaceAll(word, "i", "1")
	word = strings.ReplaceAll(word, "l", "!")
	word = strings.ReplaceAll(word, "n", "N")
	word = strings.ReplaceAll(word, "s", "5")
	word = strings.ReplaceAll(word, "x", "*")
	return word
}

// Fungsi variadic untuk menggabungkan kata-kata dan mengubah menjadi alay
func AlayGen(words ...string) string {
	var result string
	for _, word := range words {
		result += convertToAlay(word) + " "
	}
	return strings.TrimSpace(result)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
