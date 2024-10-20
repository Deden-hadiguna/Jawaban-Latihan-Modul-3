package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Fungsi untuk menghitung permutasi
func permutation(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / factorial(n-r)
}

// Fungsi untuk menghitung kombinasi
func combination(n, r int) int {
	if n < r {
		return 0
	}
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Println("Masukkan empat bilangan bulat (a, b, c, d):")
	fmt.Scan(&a, &b, &c, &d)

	// Menghitung hasil
	permutA := permutation(a, c)
	combinA := combination(a, c)
	permutB := permutation(b, d)
	combinB := combination(b, d)

	// Menampilkan hasil
	fmt.Printf("Permutasi dari %d terhadap %d: %d\n", a, c, permutA)
	fmt.Printf("Kombinasi dari %d terhadap %d: %d\n", a, c, combinA)
	fmt.Printf("Permutasi dari %d terhadap %d: %d\n", b, d, permutB)
	fmt.Printf("Kombinasi dari %d terhadap %d: %d\n", b, d, combinB)
}
