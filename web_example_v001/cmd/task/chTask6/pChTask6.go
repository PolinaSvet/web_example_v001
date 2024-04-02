package chTask6

import (
	"fmt"
	"sync"
	"time"
)

const nameTask string = "17.7.1> "

const maxValueLimit int = 10000
const numGoroutinesLimit int = 100

type Counter struct {
	value    int
	maxValue int
	mu       sync.Mutex
}

func NewCounter(maxValue int) *Counter {
	if maxValue > maxValueLimit {
		maxValue = maxValueLimit
	}
	return &Counter{maxValue: maxValue}
}

func (c *Counter) Increment(ch chan struct{}, i int, chMess chan string) {
	for {
		c.mu.Lock()
		if c.value >= c.maxValue {
			c.mu.Unlock()
			break
		}
		c.value++
		chMess <- fmt.Sprintf(nameTask+"%v: id: %d, counter : %d\n", time.Now().Format("15:04:05.000"), i, c.value)
		c.mu.Unlock()
		ch <- struct{}{}
	}
}

func StartTask(channelMess chan string, maxValue_ int, numGoroutines_ int) {

	channelMess <- fmt.Sprintf(nameTask + " === Задание 17.7.1 (HW-04) === ")

	maxValue := maxValue_           //100
	numGoroutines := numGoroutines_ //5

	if numGoroutines > numGoroutinesLimit {
		numGoroutines = numGoroutinesLimit
	}

	counter := NewCounter(maxValue)
	ch := make(chan struct{})

	for i := 0; i < numGoroutines; i++ {
		go counter.Increment(ch, i, channelMess)
	}

	for i := 0; i < maxValue; i++ {
		<-ch
	}

	close(ch)

	channelMess <- fmt.Sprintf(nameTask + " === end === ")
}
