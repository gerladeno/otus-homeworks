package hw06_pipeline_execution //nolint:golint,stylecheck

import (
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	sleepPerStage = time.Millisecond * 100
	fault         = sleepPerStage / 2
)

func TestPipeline(t *testing.T) {
	// Stage generator
	g := func(name string, f func(v interface{}) interface{}) Stage {
		return func(in In) Out {
			out := make(Bi)
			var wg = sync.WaitGroup{}
			go func() {
				defer close(out)
				defer wg.Wait()
				for v := range in {
					wg.Add(1)
					v := v
					go func() {
						defer wg.Done()
						time.Sleep(sleepPerStage)
						out <- f(v)
					}()
				}
			}()
			return out
		}
	}

	stages := []Stage{
		g("Dummy", func(v interface{}) interface{} { return v }),
		g("Multiplier (* 2)", func(v interface{}) interface{} { return v.(int) * 2 }),
		g("Adder (+ 100)", func(v interface{}) interface{} { return v.(int) + 100 }),
		g("Stringifier", func(v interface{}) interface{} { return strconv.Itoa(v.(int)) }),
	}

	t.Run("simple case", func(t *testing.T) {
		in := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.ElementsMatch(t, []string{"102", "104", "106", "108", "110"}, result)
		require.Less(t,
			int64(elapsed),
			int64(sleepPerStage)*int64(len(stages))+int64(fault))
	})

	t.Run("done case", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		// Abort after 200ms
		abortDur := sleepPerStage * 2
		go func() {
			<-time.After(abortDur)
			close(done)
		}()

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Len(t, result, 0)
		require.Less(t, int64(elapsed), int64(abortDur)+int64(fault))
	})

	t.Run("chan not closed", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		data := []int{1, 2, 3, 4, 5}

		abortDur := 1200 * time.Millisecond
		go func() {
			<-time.After(abortDur)
			close(done)
		}()

		go func() {
			for _, v := range data {
				in <- v
			}
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)
		require.Len(t, result, 5)
		require.Less(t, int64(abortDur), int64(elapsed)+int64(fault))
		require.Less(t, int64(elapsed), int64(abortDur)+int64(fault))
	})

	t.Run("empty slice of workers", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		data := []string{"1", "2", "3", "4", "5"}

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		var stages []Stage
		result := make([]string, 0, 10)
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}
		require.Equal(t, result, data)
	})

	t.Run("empty input", func(t *testing.T) {
		in := make(Bi)
		done := make(Bi)
		var data []string

		go func() {
			for _, v := range data {
				in <- v
			}
			close(in)
		}()

		result := make([]string, 0, 10)
		start := time.Now()
		for s := range ExecutePipeline(in, done, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)
		require.Len(t, result, 0)
		require.Less(t, int64(elapsed), int64(fault))
	})

	t.Run("many input values case", func(t *testing.T) {
		in := make(Bi)
		n := 10000
		data := make([]int, n)

		go func() {
			for i, _ := range data {
				in <- i
			}
			close(in)
		}()

		result := make([]string, 0, n)
		start := time.Now()
		for s := range ExecutePipeline(in, nil, stages...) {
			result = append(result, s.(string))
		}
		elapsed := time.Since(start)

		require.Len(t, result, n)
		require.Less(t,
			int64(elapsed),
			int64(sleepPerStage)*int64(len(stages))+int64(fault))
	})
}
