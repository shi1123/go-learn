package main

import "fmt"

func main() {
	arr := [5]string{"I","am","stupid","and","weak"}
	fmt.Println(arr)
	arr[2] = "smart"
	arr[4] = "strong"
	fmt.Println(arr)
}