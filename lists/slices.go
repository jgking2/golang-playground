package lists

import "fmt"

func AllThingsSlices() {
	var slice []int
	result := append(slice, 4)
	slice2 := []int{0, 1, 2, 3}
	//Append two slices
	fmt.Println(append(result, slice2...))
	for index, value := range slice2 {
		fmt.Printf("Index: %[1]v Value: %[2]v\n", index, value)
	}
}
