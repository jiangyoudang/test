package main

import (
	"fmt"
	_ "time"
)

var quit chan int = make(chan int)


func loop()  {
	for i:=1; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	quit <- 0

}

func main() {
	go loop()
	go loop()

//	time.Sleep(2*time.Second)

	for i:=0; i<2; i++ {
		<- quit
	}
}

