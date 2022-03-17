package main

import (
	"fmt"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var (
		result float32
		sisi float32
		panjang float32
		lebar float32
		alas float32
		tinggi float32
		r float32 
		choice float32
	)
	fmt.Println("masukan pilihan :")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("masukan sisi ")
		fmt.Scan(&sisi)
		result = sisi * sisi
		fmt.Println(result)
	case 2:
		fmt.Println("masukan panjang : ")
		fmt.Scan(&panjang)
		fmt.Println("masukan lebar : ")
		fmt.Scan(&lebar)
		result = panjang * lebar
		fmt.Println(result)
	case 3:
		fmt.Println("masukan alas :")
		fmt.Scan(&alas)
		fmt.Println("masukan tinggi :")
		fmt.Scan(&tinggi)
		result = alas * tinggi
		fmt.Println(result)
	case 4:
		fmt.Println("masukan jari-jari :")
		fmt.Scan(&r)
		var pi float32 = 3.14
		result = pi * r * r
		fmt.Printf("%f",result)
	default:
		fmt.Println("Invalid")

	}
	
}