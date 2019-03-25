package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func ExploreConcurrency() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.AfterFunc(time.Second*7, wg.Done)
		infiniteLoop(0)
	}()

	go func() {
		time.AfterFunc(time.Second*5, wg.Done)
		infiniteLoop(40)
	}()

	wg.Wait()
}

func infiniteLoop(startAt int) {
	for i := startAt; ; {
		i++
		fmt.Printf("Hello %d\n", i)
		time.Sleep(time.Millisecond * 500)
	}
}
