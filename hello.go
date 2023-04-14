package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	now := time.Now()
	m, _ := time.ParseDuration("+2h")
	m1 := now.Add(m)

	fmt.Println(m1.Format("2006-01-02 15:04:05"))

	fmt.Println("Hello, time!")

	// : * ? " < > | \ /  & =
	re,_ := regexp.Compile("[/:*?\"<>\\\\|'&=]")

	txt := re.ReplaceAllString(`"/\\2&=|?aaa=111 *'/test\". 3.html:www?e<aaa>"`,"")

	fmt.Println(txt)

	fmt.Println("test v2!")


}