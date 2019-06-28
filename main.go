package main

import (
	"fmt"
	"math"
	"./bucket"
	"time"
)

func main() {
	fmt.Println(int64(math.Pow(10,9)))

	b := bucket.New(uint64(10), time.Now().Unix())
	for i:=0; i < 100; i++{
		go func() {
			if bucket.IsLimited(b) {
				fmt.Println("limited, not ok")
			} else {
				fmt.Println("ok , capacity = ", b.GetCapacity())
			}
		}()
	}
}


