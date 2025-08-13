package main

import (
	"reflect"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
	}()

	go func() {
		defer close(ch3)
		ch3 <- 4
		ch3 <- 5
	}()

	resultCh := MergeChannels(ch1, ch2, ch3)

	var result []int
	for v := range resultCh {
		result = append(result, v)
	}

	expected := []int{1, 2, 3, 4, 5}
	if len(result) != len(expected) {
		t.Fatalf("Ожидалось %v элементов, получено %v", len(expected), len(result))
	}

	// Проверка, что все элементы присутствуют
	for _, val := range expected {
		found := false
		for _, r := range result {
			if r == val {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Элемент %d отсутствует в результате", val)
		}
	}
}

func TestMergeChannelsEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	close(ch1)
	close(ch2)

	resultCh := MergeChannels(ch1, ch2)
	result := []int{}
	for v := range resultCh {
		result = append(result, v)
	}

	if !reflect.DeepEqual(result, []int{}) {
		t.Errorf("Ожидался пустой результат, получено %v", result)
	}
}

func TestMergeStringChannels(t *testing.T) {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		defer close(ch1)
		ch1 <- "a"
		ch1 <- "b"
	}()

	go func() {
		defer close(ch2)
		ch2 <- "c"
		ch2 <- "d"
	}()

	merged := MergeChannels(ch1, ch2)
	var results []string
	for s := range merged {
		results = append(results, s)
	}

	expected := []string{"a", "b", "c", "d"}
	if len(results) != len(expected) {
		t.Errorf("Ожидалось 4 значения, получено %d", len(results))
	}
	// Проверяем содержимое (порядок может быть произвольным!)
	expectedChar := map[string]bool{
		"a": true,
		"b": true,
		"c": true,
		"d": true,
	}

	for _, s := range results {
		if !expectedChar[s] {
			t.Errorf("Неожиданное значение %q", s)
		}
		delete(expectedChar, s) // Для проверки дубликатов
	}

	if len(expectedChar) > 0 {
		t.Errorf("Не все ожидаемые значения получены: %v", expected)
	}
}
