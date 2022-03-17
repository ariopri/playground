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
		r float32
		t float32
		pi float32 = 3.14
		
	)
	fmt.Printf("Masukan jari-jari alas tabung : %.0f ",r)
	fmt.Scan(&r)
	fmt.Printf("Masukan tinggi tabung : %.0f",t)
	fmt.Scan(&t)
	volume :=  (pi *r * r) * t
	fmt.Printf("jadi volimenya adalah : %f", volume)
}