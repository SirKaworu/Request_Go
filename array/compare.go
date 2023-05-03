package main

import (
	"bufio" //input stream buffer
	"fmt"
	"os" //akses os untuk os.stdin
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin) //membaca input terminal, //new reader efisien seperti readstring

	// Input array 1
	fmt.Print("Masukkan array 1 : ")
	arr1Input, _ := reader.ReadString('\n')              //membaca input sampai newline
	arr1 := strings.Fields(strings.TrimSpace(arr1Input)) //Field untuk memecah, TrimSpace untuk menghapus whitespace

	// Input array 2
	fmt.Print("Masukkan array 2 : ")
	arr2Input, _ := reader.ReadString('\n')
	arr2 := strings.Fields(strings.TrimSpace(arr2Input))

	// Bandingkan elemen
	minLength := len(arr1)
	if len(arr2) < minLength {
		minLength = len(arr2) //melihat panjang arr terpendek mencegah out of range
	} //rentang index yang valid pada kedua array, 0 hingga minLength - 1 saat looping

	// Validasi input
	if len(arr1) != len(arr2) {
		fmt.Println("Jumlah elemen pada kedua array berbeda.")
		return
	}

	var beda []int                   // slice function untuk menampung index yang berbeda
	for i := 0; i < minLength; i++ { //mengambil nilai minimum sebagai batas loop
		if arr1[i] != arr2[i] {
			beda = append(beda, i) //append menambahkan ke slice
		}
	}

	if len(beda) > 0 { //jika terdapat index yang berbeda
		fmt.Printf("Kedua array berbeda pada index %v\n", beda)
	} else {
		fmt.Println("Kedua array sama")
	}
}
