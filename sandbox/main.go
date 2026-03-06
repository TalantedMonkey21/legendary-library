package main

import (
	"time"
	"fmt"
)

type Counter struct{
	value int
	name *string
}
type Note struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Created_at time.Time `json:"created_at"`
}

var Notes = []Note{
	{Id: 0, Title: "Pushkin", Body: "russian writer", Created_at: time.Now()},
	{Id: 1, Title: "Tolstoy", Body: "russian writer", Created_at: time.Now()},
	{Id: 2, Title: "Dostoevsky", Body: "russian writer", Created_at: time.Now()},
}
func (c *Counter) change() {
	fmt.Println(&c.value)
}
func main() {
	s := "Jhon"
	counter := Counter{name: &s}
	fmt.Println(counter.value)
	fmt.Println(counter.name)
	fmt.Println(*counter.name)
	s = "Max"
	fmt.Println()
	fmt.Println(&counter.value)
	fmt.Println(counter.value)
	fmt.Println(counter.name)
	fmt.Println(*counter.name)
	fmt.Println()
	counter.change()

	fmt.Println(len(Notes))
}
