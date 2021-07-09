package melon

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMelon(t *testing.T) {
	wg := new(sync.WaitGroup)
	me := New(5000, OptionAnchor(100, 70))
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 200; i++ {
		fmt.Println("goroutine", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			concurrency(ctx, me)
		}()
	}
	go monitor(ctx, cancel, me)
	fmt.Println("goroutine wait")
	wg.Wait()
	fmt.Println(me.Stats())
}

func monitor(ctx context.Context, cancel context.CancelFunc, me *melon) {
	var i int
	for {
		i++
		time.Sleep(time.Second)
		begin := time.Now()
		good := me.Good()
		interval := time.Since(begin).Nanoseconds() / 1e6
		fmt.Println("stat good", i, good, me.index, interval)
		if i > 200 {
			cancel()
			fmt.Println("canceled")
			return
		}
	}
}

func concurrency(ctx context.Context, me *melon) {
	var i int
	var sweet bool
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		sweet = true
		i++
		if i%20 == 0 {
			sweet = false
		}
		me.Feed(sweet)
	}
}

func BenchmarkMelon(b *testing.B) {
	me := New(5000, OptionAnchor(100, 70))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 == 0 {
			me.Feed(true)
		} else {
			me.Feed(false)
		}
		me.Good()
	}
}
