package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menentukan posisi titik terhadap lingkaran
func checkPosition(x, y, cx, cy, r float64) string {
	distance := math.Sqrt(math.Pow(x-cx, 2) + math.Pow(y-cy, 2))
	if distance < r {
		return "Titik di dalam lingkaran"
	} else if distance == r {
		return "Titik di batas lingkaran"
	} else {
		return "Titik di luar lingkaran"
	}
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	// Input untuk lingkaran 1
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1):")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk lingkaran 2
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2):")
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk titik sembarang
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&x, &y)

	// Mengecek posisi titik terhadap lingkaran 1 dan 2
	posisiLingkaran1 := checkPosition(x, y, cx1, cy1, r1)
	posisiLingkaran2 := checkPosition(x, y, cx2, cy2, r2)

	// Menentukan hasil akhir
	if posisiLingkaran1 == "Titik di dalam lingkaran" && posisiLingkaran2 == "Titik di dalam lingkaran" {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if posisiLingkaran1 == "Titik di dalam lingkaran" {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if posisiLingkaran2 == "Titik di dalam lingkaran" {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
