package chTask1

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const nameTask string = "17.3.1> "

// Шаг наращивания счётчика
const step int64 = 100

// Конечное значение счетчика
const endCounterValue int64 = 1000

func StartTask(channelMess chan string) {
	var counter int64 = 0
	var wg sync.WaitGroup
	channelMess <- fmt.Sprintf(nameTask + " === Задание 17.3.1 (HW-04) === ")
	channelMess <- fmt.Sprintf(nameTask+"Конечное значение счетчика:%v, Шаг наращивания счётчика:%v", endCounterValue, step)

	increment := func(i int) {
		defer wg.Done()

		atomic.AddInt64(&counter, step)
		channelMess <- fmt.Sprintf(nameTask+"%v: id:%v, step:%v", time.Now().Format("15:04:05.000"), i, step)
	}
	// Не всегда вычисление этой переменной будет приводить к верному
	// результату в счётчике, но для правильных значений
	// и для удобства - можно
	var iterationCount int = int(endCounterValue / step)
	for i := 1; i <= iterationCount; i++ {
		wg.Add(1)
		go increment(i)
	}
	// Ожидаем поступления сигнала
	wg.Wait()
	// Печатаем результат, надеясь, что будет 1000
	channelMess <- fmt.Sprintf(nameTask+"Печатаем результат, надеясь, что будет 1000: %v", counter)
	channelMess <- fmt.Sprintf(nameTask + " === end === ")
}
