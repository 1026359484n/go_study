package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground")
	var t = time.Now()
	hour, min, sec := t.Clock()
	fmt.Println("The time is", t)
	fmt.Println("The time is", t.AddDate(0,0,1))
	fmt.Println([]string{"The time is", strconv.Itoa(hour)}, "点", strconv.Itoa(min), "分", strconv.Itoa(sec), "秒")
	fmt.Println(append([]string{"The time is", strconv.Itoa(hour)}, "点", strconv.Itoa(min), "分", strconv.Itoa(sec), "秒"))
}
