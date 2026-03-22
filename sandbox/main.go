package main

import (
	"fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
	ch := make (chan int)

	go func(){
		for i := 0; i < 10; i++ { // 0 1 2 3 4
			ch <- i
		}
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}

