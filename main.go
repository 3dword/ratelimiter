package main

import (
	"fmt"
	"math"
	"./bucket"
	"time"
)

func main() {

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


