package chTask2

import (
	"fmt"
	"sync"
	"time"
)

const nameTask string = "17.3.2> "

const step int = 1
const iterationAmount int = 1000
const iterationAmountGo int = 1050

func StartTask(channelMess chan string) {

	channelMess <- fmt.Sprintf(nameTask + " === Задание 17.3.2 (HW-04) === ")

	var counter int
	var c = sync.NewCond(&sync.Mutex{})
	finished := false

	increment := func(i int) {
		c.L.Lock()
		if counter < iterationAmount && !finished {
			counter += step
			channelMess <- fmt.Sprintf(nameTask+"%v: id:%04d from %04d, counter:%v", time.Now().Format("15:04:05.000"), i, iterationAmountGo, counter)
		}
		if counter == iterationAmount && !finished {
			finished = true
			c.Broadcast()
		}
		c.L.Unlock()
	}

	for i := 1; i <= iterationAmountGo; i++ {
		go increment(i)
	}

	c.L.Lock()
	for !finished {
		c.Wait()
	}
	c.L.Unlock()

	channelMess <- fmt.Sprintf(nameTask+"%v: counter:%v", time.Now().Format("15:04:05.000"), counter)
	channelMess <- fmt.Sprintf(nameTask + " === end === ")
}
