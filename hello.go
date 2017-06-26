package main

import "fmt"

func main() {
	fmt.Println("hello world go")
	chan1 := make(chan bool, 2)
	chan2 := make(chan int, 1)
	chan1 <- false
	select {
	case b := <-chan1:
		// b <- chan1
		fmt.Println(b, "striopss")
	case chan2 <- 3:
		fmt.Println(<-chan2, "int")
	default:
		fmt.Println("exit select")
	}

	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int, 2)
	go sunArr(s[:len(s)/2], c)
	// z := <-c
	go sunArr(s[len(s)/2:], c)
	// y := <-c
	// x, y := <-c, <-c
	// fmt.Println(x, y)
}

func sunArr(arr []int, c chan int) {
	// defer close(c)
	sum := 0
	for _, element := range arr {
		sum += element
		fmt.Println(sum)
	}
	c <- sum
}
