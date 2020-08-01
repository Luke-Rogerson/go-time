package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("The time now is %d:%d\n", t.Hour(), t.Minute())

}
