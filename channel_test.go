package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	channel <- "wanta"
	fmt.Println(<- channel)

	defer close(channel)


	go func () {
		time.Sleep(2 * time.Second)
		channel <- "wanta zulfikri"
		fmt.Println(",emgambil data dichannel")

	}()
	data := <- channel 
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}