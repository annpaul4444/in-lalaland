package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello from the other side")
	fmt.Println(time.Now().Format("02.01.2006 Monday"))
}
