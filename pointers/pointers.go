package pointers

import "fmt"

func AllThingsPointers() {
	num, str := 13, "Hello"

	//Point to num
	test := &num
	*test = 21
	fmt.Println(num)
	test2 := &str
	*test2 = "Why now?"
	fmt.Println(str)
	fmt.Println(*test2)
}
