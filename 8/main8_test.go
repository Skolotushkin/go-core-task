package main

import (
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	t.Run("Basic Add/Down", func(t *testing.T) {
		sem := NewSemaphore(2)

		sem.Add(1)
		if len(sem) != 1 {
			t.Errorf("Expected 1 Addn slot, got %d", len(sem))
		}

		sem.Down(1)
		if len(sem) != 0 {
			t.Errorf("Expected all slots Downd, got %d", len(sem))
		}
	})

	t.Run("Blocking When Full", func(t *testing.T) {
		sem := NewSemaphore(1)
		sem.Add(1) // заполняем семафор

		done := make(chan bool)
		go func() {
			sem.Add(1) // этот вызов должен заблокироваться
			done <- true
		}()

		select {
		case <-done:
			t.Error("Add should have blocked")
		case <-time.After(100 * time.Millisecond):
			// Ожидаемое поведение - операция блокируется
		}

		sem.Down(1) // освобождаем семафор
		<-done      // теперь горутина должна разблокироваться
	})
	t.Run("Zero Capacity", func(t *testing.T) {
		sem := NewSemaphore(0)
		done := make(chan bool)

		go func() {
			sem.Add(1)
			done <- true
		}()

		select {
		case <-done:
			t.Error("Add should block on zero capacity semaphore")
		case <-time.After(100 * time.Millisecond):
			// Ожидаемое поведение
		}
	})
}
