package main

import "fmt"

const size = 30

//0,1,1,2,3,5,8,13,21,...

func main() {
	var fibonacci []int
	for i := 0; i < size; i++ {
		if i == 0 || i == 1 {
			fibonacci = append(fibonacci, i)
		} else {
			fibonacci = append(fibonacci, fibonacci[i-2]+fibonacci[i-1])
		}
	}
	fmt.Printf("%v", fibonacci)
}
