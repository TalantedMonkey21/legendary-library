package main

import (
	"fmt"
)

type Person struct {
    Name string
    Age  int
}

func main() {
	ext := make([]int, 2, 6)
	ext[0] = 10
	ext[1] = 20

	src := []int{1, 2, 3}
	dst := append(ext, src...)

	ext[0] = 999
	fmt.Println(dst[0])
	fmt.Println(ext)
	fmt.Println(dst)
}

