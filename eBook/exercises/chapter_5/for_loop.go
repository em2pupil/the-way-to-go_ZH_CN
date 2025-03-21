package main

import "fmt"

func main() {
	// 1:
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}

	for j := 0; j < 15; j++ {
		fmt.Println(j)
	}

	// 2:
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}

	j := 0
start:
	fmt.Println(j)
	if j < 15 {
		j++
		goto start
	}
}
