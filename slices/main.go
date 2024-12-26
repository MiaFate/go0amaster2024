package main

import "fmt"

func main() {
	weekDays := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	daySlice := weekDays[0:5]
	fmt.Println(daySlice)

	daySlice = append(daySlice, "Friday", "Saturday", "other days")
	fmt.Println(daySlice)

	daySlice = append(daySlice[:2], daySlice[3:]...)
	fmt.Println(daySlice)

	fmt.Println(len(daySlice))
	fmt.Println(cap(daySlice))

	//make
	names := make([]string, 5, 10)
	names[0] = "Mia"
	fmt.Println(names)

	slice := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	fmt.Println(copy(slice2, slice))
	fmt.Println(slice)
	fmt.Println(slice2)

}
