package chTask5

import (
	"fmt"
	"time"
)

const nameTask string = "17.6.3> "

func sender(ch chan int, chMess chan string) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
	chMess <- fmt.Sprintf(nameTask + " === end === ")
}

func receiver(ch chan int, chMess chan string) {
	for num := range ch {
		chMess <- fmt.Sprintf(nameTask+"%v: Получено из канала : %d\n", time.Now().Format("15:04:05.000"), num)
	}
}

func StartTask(channelMess chan string) {
	ch := make(chan int)

	channelMess <- fmt.Sprintf(nameTask + " === Задание 17.6.3 (HW-04) === ")

	go sender(ch, channelMess)
	go receiver(ch, channelMess)

	// Для ожидания завершения программы
	select {}

}
