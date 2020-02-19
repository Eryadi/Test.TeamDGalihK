package main

import (
	"fmt"
	"time"
)

func main() {
	// Cara 1
	var nama1 string
	nama1 = "Willy"
	// Cara 2
	nama2 := "Willy"
	fmt.Println(nama1, nama2)

	// angka int
	angka := 9
	fmt.Println(angka)
	// angka float
	angkaF := 9.2
	fmt.Println(angkaF)

	// array
	arr := []string{"asd", "fgh"}
	fmt.Println(arr)

	// _ variable buangan

	// loop variable
	for _, v := range arr {
		fmt.Printf("Value: %s\n", v)
	}

	// loop biasa
	// for i := 0; i < 5; i++ {

	// }

	// contoh operator
	// var value = ((2 + 6) / 4 * 2) % 4
	// == != < <= > >=
	// && ||

	// selection if
	// if value == 0 {
	// 	fmt.Println("NOL")
	// } else {
	// 	fmt.Println(value)
	// }

	// selection switch/case
	// menu := 1
	// switch menu {
	// case 1:
	// case 2:
	// default:
	// }

	// panggil fungsi
	hasil := hitung(2, 5)
	fmt.Println(hasil)

	var1, var2, var3, _ := bagiDanJumlah(2.0, 1.0)
	fmt.Println(var1, var2, var3)

	var user User
	user.Nama = "Willy"
	user.Umur = 21

	fmt.Printf("%+v\n", user)

	// map[key]value
	var map1 = make(map[string]string)
	// JAVA Map<String,String>
	map1["Nama"] = "Willy"
	// map1["Umur"] = 21
	fmt.Println(map1["Nama"])

	// map[key]value
	var map2 = make(map[string]interface{})
	// JAVA Map<String,String>
	map2["Nama"] = "Willy"
	map2["Umur"] = 21
	fmt.Println(map2["Umur"])

	// tanggal + formatting
	var tglLahir time.Time
	tglLahir = time.Now()
	// JAVA YYYY-MM-DD
	// format wajib 2006-01-02 15:04:05
	fmt.Println(tglLahir.Format("2006-01-02"))
}

// return 1 value
func hitung(a int, b int) int {
	return a + b
}

// return multiple values
func bagiDanJumlah(a float64, b float64) (float64, float64, float64, float64) {
	return a / b, a + b, a - b, a * a
}

// User object model
type User struct { // <-- public
	Nama string `json:"pro_name" db:"Pro_Name"`
	Umur int
}

// A object model
type a struct { // <-- private
	nama string
	umur int
}

// JAVA
// class User {
// 	string Nama;
// 	int Umur;
// }

// JAVA
// public int hitung(int a, int b) {
// 	return a+b
// }
