package main

import (
	"fmt"
	"time"
)

type Semaphore chan struct{}

func NewSemaphore(max int) Semaphore {
	return make(Semaphore, max)
}

func (s Semaphore) Add(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{} // отправляем в канал
	}
}

func (s Semaphore) Down(n int) {
	for i := 0; i < n; i++ {
		<-s // забираем из канала
	}
}

func main() {
	workers := 5
	sem := NewSemaphore(workers)

	for i := 1; i <= 10; i++ {
		sem.Add(1) // ждём свободного места
		go func(id int) {
			defer sem.Down(1) // обязательно освобождаем!
			fmt.Printf("Worker %d started\n", id)
			time.Sleep(time.Second) // имитация работы
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	// Ждём завершения всех горутин
	sem.Add(workers) // ждём, пока все workers мест не освободятся
}
