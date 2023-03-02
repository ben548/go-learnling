package main

import "fmt"

func main() {
	//fmt.Println("before panic")
	//panic("panicing")
	//fmt.Println("after panic")

	//fmt.Println("before panic")
	//defer func() {
	//	fmt.Println("after panic")
	//}()
	//panic("panicing")

	fmt.Println("before panic")
	defer func() {
		fmt.Println("after panic")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()
	panic("panicing")
}
