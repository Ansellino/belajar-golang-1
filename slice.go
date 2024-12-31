package main

import "fmt"

func main(){
	// membuat slice dari array yang sudah ada
	names := [...]string{"Januari","Februari","Maret","April","Mei","Juni"} 
	slice1 := names[4:6] // 5 dan 6
	fmt.Println(slice1)

	slice2 := names[:3] // 1 dan 2
	fmt.Println(slice2)

	slice3 := names[3:] // 4, 5, dan 6
	fmt.Println(slice3)

	slice4 := names[:] // semua
	fmt.Println(slice4)

	days := [...]string{"Senin","Selasa","Rabu","Kamis","Jumat","Sabtu","Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println("Day Slice 1:", daySlice1)
	fmt.Println("Day Slice 2:", daySlice2)
	fmt.Println("Day :", days)

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Ansel"
	newSlice[1] = "Ansel1"
	// newSlice[2] = "Ansel2" akan error, harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Ansel2")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice) // ini yang sering di pake
}