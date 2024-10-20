package main

import (
	"fmt"
)

// Fungsi f(x) = x^2
func f(x int) int {
	return x * x
}

// Fungsi g(x) = x - 2
func g(x int) int {
	return x - 2
}

// Fungsi h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Fungsi untuk menghitung komposisi (fogoh)(x)
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi untuk menghitung komposisi (gohof)(x)
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi untuk menghitung komposisi (hofog)(x)
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int
	fmt.Println("Masukkan tiga bilangan bulat (a, b, c) yang dipisahkan oleh spasi:")
	fmt.Scan(&a, &b, &c)

	// Menghitung hasil komposisi
	resultFogoh := fogoh(a)
	resultGohof := gohof(b)
	resultHofog := hofog(c)

	// Menampilkan hasil
	fmt.Printf("(fogoh)(%d) = %d\n", a, resultFogoh)
	fmt.Printf("(gohof)(%d) = %d\n", b, resultGohof)
	fmt.Printf("(hofog)(%d) = %d\n", c, resultHofog)
}
