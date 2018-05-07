package main

import (
	"time"
	"fmt"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer1 completed")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 completed")
	}()
	//timer2.
	timer3 := time.NewTimer(8 * time.Second)
	<-timer3.C
	fmt.Println("Timer3 completed")

	stop2 := timer2.Stop()


	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
