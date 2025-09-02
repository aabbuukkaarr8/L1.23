package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	i := 3
	copy(slice[i:], slice[i+1:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
