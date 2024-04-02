package chTask4

import (
	"fmt"
	"time"
)

const nameTask string = "17.6.2> "

func backgroundTask(chMess chan string, flag *bool) {
	for {
		if *flag {
			chMess <- fmt.Sprintf(nameTask + " === end === ")
			return
		}
		chMess <- fmt.Sprintf(nameTask+"Текущее время: %v", time.Now().Format(time.Stamp))
		time.Sleep(1 * time.Second)
	}
}

func StartTask(channelMess chan string, flag *bool) {

	ch1 := make(chan int)
	ch2 := make(chan int)

	defer func() {
		close(ch1)
		close(ch2)
	}()

	channelMess <- fmt.Sprintf(nameTask + " === Задание 17.6.2 (HW-04) === ")

	go backgroundTask(channelMess, flag)

	for {
		select {
		case msg := <-ch1:
			channelMess <- fmt.Sprintf(nameTask+"Получено из канала 1: %d\n", msg)
		case msg := <-ch2:
			channelMess <- fmt.Sprintf(nameTask+"Получено из канала 2: %d\n", msg)
		}
	}

}
