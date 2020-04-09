package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		expect := 3

		assertCounter(t, counter, expect)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter()
		//A WaitGroup waits for a collection of goroutines to finish.
		//The main goroutine calls Add to set the number of goroutines to wait for.
		//Then each of the goroutines runs and calls Done when finished.
		//At the same time, Wait can be used to block until all goroutines have finished.
		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *counter, expect int) {
	if got.Value() != expect {
		t.Errorf("got %d, want %d", got.Value(), expect)
	}
}
