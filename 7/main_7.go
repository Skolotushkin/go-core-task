package main

import "sync"

// MergeChannels
// Объединяет N каналов в один
func MergeChannels[T any](channels ...<-chan T) <-chan T {
	mergedCh := make(chan T)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(channels))
		for _, ch := range channels {
			go func(ch <-chan T, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)
		}
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}
