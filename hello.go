package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	m, _ := time.ParseDuration("+2h")
	m1 := now.Add(m)

	fmt.Println(m1.Format("2006-01-02 15:04:05"))

}