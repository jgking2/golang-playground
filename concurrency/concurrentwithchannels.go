package concurrency

import (
	"fmt"
	"time"
)

var uselessStrings = [6]string{"hello", "world", "what", "it", "do", ":)"}

func ExploreConcurrencyWithChannels() {

	testChannel := make(chan int)

	go func() {
		outPutUselessStuffInChannel(testChannel)
	}()

	//Grab the data
	for data := range testChannel {
		fmt.Println(data)
	}
}

func ExploreConcurrencyWithMultipleChannels() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		outPutUselessStuffInChannel(chan1)
	}()

	go func() {
		outPutUselessStrings(chan2)
	}()

Channels:
	for {
		select {
		case num := <-chan1:
			if num == 0 {
				break
			}
			fmt.Println(num)
		case str := <-chan2:
			if str == "" {
				break Channels
			}
			fmt.Println(str)
		}
	}
}

func outPutUselessStuffInChannel(channel chan<- int) {
	defer close(channel)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		channel <- i
	}
}

func outPutUselessStrings(channel chan<- string) {
	defer close(channel)
	for _, value := range uselessStrings {
		time.Sleep(time.Second * 1)
		channel <- value
	}
}
