package main

import (
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	expected := []float64{1.0, 8.0, 27.0}

	pipeline(in, out)

	go func() {
		for _, v := range []uint8{1, 2, 3} {
			in <- v
		}
		close(in)
	}()

	var results []float64
	for val := range out {
		results = append(results, val)
	}

	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, results)
	}
}

func TestPipelineEmpty(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	pipeline(in, out)

	close(in) // сразу закрываем вход

	var results []float64
	for val := range out {
		results = append(results, val)
	}

	if len(results) != 0 {
		t.Errorf("Ожидался пустой результат, получено %v", results)
	}
}
