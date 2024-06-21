package main

import "fmt"

// func sum(s []int, ch chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	ch <- sum // send sum to ch
// }

func main() {
	ch := make(chan int, 1)
	ch <- 1           // ch = { 1 }
	ch <- 2           // wait
	fmt.Println(<-ch) // unreachable statement
	fmt.Println(<-ch)
}

// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(10, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for { // while(true) runs forever
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	defer close(c)
// 	defer close(quit)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c) // blocked
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

// ch := make(chan int)
// 	go sum([]int{1, 2, 3}, ch)
// 	go sum([]int{4, 5, 6}, ch)
// 	x := <-ch // receive from ch
// 	y := <-ch // receive from ch
// 	fmt.Println(x, y, x+y)
