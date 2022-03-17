package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012
func main() {
	// TODO: answer here
	var (
		r float32 = 2
		t float32 = 20
		pi float32 = 3.14
		
	)
	luas :=  (pi *r * r) * t
	fmt.Printf("jadi volimenya adalah : %f", luas)
}