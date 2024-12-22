package main

import (
	"fmt"
	"time"
)

// Data mahasiswa dan nilai semester
var students = map[string][2][5]float64{
	"Alice":   {[5]float64{85.5, 92.2, 88.1, 0, 0}, [5]float64{89.0, 90.5, 91.3, 0, 0}},
	"Bob":     {[5]float64{78.3, 82.7, 84.9, 0, 0}, [5]float64{81.2, 83.4, 85.6, 0, 0}},
	"Charlie": {[5]float64{91.5, 88.7, 90.2, 0, 0}, [5]float64{90.1, 87.9, 89.3, 0, 0}},
	"David":   {[5]float64{75.8, 83.1, 79.6, 0, 0}, [5]float64{82.4, 80.7, 83.2, 0, 0}},
	"Eve":     {[5]float64{88.2, 92.5, 90.8, 0, 0}, [5]float64{90.3, 91.1, 89.4, 0, 0}},
	"Frank":   {[5]float64{83.7, 85.9, 87.2, 0, 0}, [5]float64{86.5, 88.3, 85.9, 0, 0}},
	"Grace":   {[5]float64{90.4, 88.6, 92.3, 0, 0}, [5]float64{89.7, 91.2, 90.8, 0, 0}},
	"Henry":   {[5]float64{82.1, 84.5, 83.9, 0, 0}, [5]float64{85.2, 83.8, 86.4, 0, 0}},
	"Iris":    {[5]float64{87.6, 89.3, 88.5, 0, 0}, [5]float64{88.9, 90.2, 89.7, 0, 0}},
	"Jack":    {[5]float64{79.8, 81.4, 80.7, 0, 0}, [5]float64{82.3, 83.1, 81.9, 0, 0}},
	"Kelly":   {[5]float64{86.3, 88.9, 87.5, 0, 0}, [5]float64{88.1, 89.4, 88.7, 0, 0}},
	"Leo":     {[5]float64{84.2, 86.7, 85.4, 0, 0}, [5]float64{86.8, 87.5, 85.9, 0, 0}},
	"Mia":     {[5]float64{89.5, 91.2, 90.3, 0, 0}, [5]float64{90.7, 92.1, 91.4, 0, 0}},
	"Noah":    {[5]float64{81.9, 83.6, 82.8, 0, 0}, [5]float64{84.2, 85.7, 83.9, 0, 0}},
	"Olivia":  {[5]float64{88.7, 90.4, 89.6, 0, 0}, [5]float64{90.1, 91.5, 90.8, 0, 0}},
	"Peter":   {[5]float64{80.5, 82.8, 81.7, 0, 0}, [5]float64{83.4, 84.9, 82.6, 0, 0}},
	"Quinn":   {[5]float64{85.9, 87.3, 86.5, 0, 0}, [5]float64{87.8, 88.6, 87.2, 0, 0}},
	"Rachel":  {[5]float64{88.1, 89.7, 88.9, 0, 0}, [5]float64{89.5, 90.8, 89.9, 0, 0}},
	"Sam":     {[5]float64{82.4, 84.1, 83.2, 0, 0}, [5]float64{84.7, 85.9, 84.3, 0, 0}},
	"Tara":    {[5]float64{87.2, 88.9, 88.1, 0, 0}, [5]float64{89.3, 90.5, 89.7, 0, 0}},
	"Uma":     {[5]float64{83.6, 85.2, 84.4, 0, 0}, [5]float64{85.8, 86.9, 85.5, 0, 0}},
	"Victor":  {[5]float64{81.3, 83.7, 82.5, 0, 0}, [5]float64{84.1, 85.3, 83.9, 0, 0}},
	"Wendy":   {[5]float64{86.8, 88.5, 87.7, 0, 0}, [5]float64{88.9, 90.1, 89.3, 0, 0}},
	"Xavier":  {[5]float64{84.7, 86.3, 85.5, 0, 0}, [5]float64{86.9, 88.1, 87.3, 0, 0}},
	"Yara":    {[5]float64{89.3, 90.9, 90.1, 0, 0}, [5]float64{91.2, 92.4, 91.6, 0, 0}},
	"Zack":    {[5]float64{82.9, 84.5, 83.7, 0, 0}, [5]float64{85.1, 86.3, 85.5, 0, 0}},
}

// Fungsi iteratif untuk menghitung rata-rata
func calculateAverageIterative(scores [5]float64) float64 {
	sum := 0.0
	for _, score := range scores {
		sum += score
	}
	return sum / float64(len(scores))
}

// Fungsi rekursif untuk menghitung rata-rata
func calculateAverageRecursive(scores [5]float64, n int) float64 {
	if n == 0 {
		return 0
	}
	return scores[n-1] + calculateAverageRecursive(scores, n-1)
}

func main() {
	var name string
	var choice int

	// Looping agar program terus berjalan hingga pengguna memilih untuk keluar
	for {
		// Meminta input nama mahasiswa
		fmt.Println("Enter the name of the student (e.g., Alice, Bob, Charlie, David, Eve) or type 'exit' to quit:")
		fmt.Scanln(&name)

		// Cek jika pengguna ingin keluar
		if name == "exit" {
			fmt.Println("Exiting the program...")
			break
		}

		// Validasi apakah mahasiswa ada dalam data
		grades, exists := students[name]
		if !exists {
			fmt.Println("Student not found! Please enter a valid name.")
			continue
		}

		// Meminta input pilihan metode perhitungan
		fmt.Println("Choose the calculation method:")
		fmt.Println("1. Iterative")
		fmt.Println("2. Recursive")
		fmt.Scanln(&choice)

		var sem1Average, sem2Average float64
		start := time.Now()

		if choice == 1 {
			// Menggunakan metode iteratif
			sem1Average = calculateAverageIterative(grades[0])
			sem2Average = calculateAverageIterative(grades[1])
		} else if choice == 2 {
			// Menggunakan metode rekursif
			sem1Average = calculateAverageRecursive(grades[0], len(grades[0])) / float64(len(grades[0]))
			sem2Average = calculateAverageRecursive(grades[1], len(grades[1])) / float64(len(grades[1]))
		} else {
			fmt.Println("Invalid choice! Please select either 1 (Iterative) or 2 (Recursive).")
			continue
		}

		elapsed := time.Since(start)
		fmt.Printf("Average for Semester 1: %.2f\n", sem1Average)
		fmt.Printf("Average for Semester 2: %.2f\n", sem2Average)
		// Menampilkan waktu dengan presisi 10 angka di belakang koma
		fmt.Printf("Calculation completed in: %.10f seconds\n", elapsed.Seconds())
	}
}
