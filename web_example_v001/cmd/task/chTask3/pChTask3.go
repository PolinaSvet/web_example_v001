package chTask3

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const nameTask string = "17.6.1> "

type Counter struct {
	value int
	mu    sync.Mutex
}

var ch1CountSent *Counter
var ch2CountSent *Counter
var ch1CountRece *Counter
var ch2CountRece *Counter

func NewCounter() *Counter {
	return &Counter{value: 0}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Load() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func sendMessage(chMess chan string, flag *bool, ch1, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if *flag {
			return
		}
		x := rand.Intn(100)

		select {
		case ch2 <- x:
			ch2CountSent.Increment()
		}
		select {
		case ch1 <- x:
			ch1CountSent.Increment()
		}

		time.Sleep(1 * time.Second)
	}
}

func receMessage(chMess chan string, flag *bool, ch1, ch2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if *flag {
			return
		}
		select {
		case msg := <-ch2:
			ch2CountRece.Increment()
			ch2CRece := ch2CountRece.Load()
			ch2CSent := ch2CountSent.Load()

			chMess <- fmt.Sprintf(nameTask+"%v: Из 2 канала: %d, ch2[%d, %d]\n", time.Now().Format("15:04:05.000"), msg, ch2CSent, ch2CRece)

			<-ch2
		case msg := <-ch1:
			ch1CountRece.Increment()
			ch1CRece := ch1CountRece.Load()
			ch1CSent := ch1CountSent.Load()

			chMess <- fmt.Sprintf(nameTask+"%v: Из 1 канала: %d, ch1[%d, %d]\n", time.Now().Format("15:04:05.000"), msg, ch1CSent, ch1CRece)
		}
	}
}

func StartTask(channelMess chan string, flag *bool) {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1CountSent = NewCounter()
	ch2CountSent = NewCounter()
	ch1CountRece = NewCounter()
	ch2CountRece = NewCounter()

	var wg sync.WaitGroup

	channelMess <- fmt.Sprintf(nameTask + " === Задание 17.6.1 (HW-04) === ")

	wg.Add(2)
	go sendMessage(channelMess, flag, ch1, ch2, &wg)
	go receMessage(channelMess, flag, ch1, ch2, &wg)

	for {
		if *flag {
			wg.Wait()
			channelMess <- fmt.Sprintf(nameTask + " === Все горутины завершили работу === ")
			return
		}
	}
}
