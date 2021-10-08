package main

import (
	"fmt"
	"time"
)
// channel 实现消费者和生产者
func main() {
	c := make(chan int,10)
	go func() {
		for i := 0; ; i++ {
			c <- i
			time.Sleep(time.Second)
		}

	}()
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Second)
			num,err := <- c
			if err {
				fmt.Println("channel close")
			}
			fmt.Printf("recieve %d\n", num)
		}
	}()
	time.Sleep(time.Second * 10)
	close(c)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}
