package main

import (
	"math/rand"
	"time"
)

// randomNumberGenerator
// Запускает горутину, которая отправляет случайные числа в небуферизированный канал
func randomNumberGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			ch <- rand.Intn(100) // случайное число от 0 до 99
		}
	}()
	return ch
}
