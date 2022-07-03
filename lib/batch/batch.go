package batch

import (
	"sync"
	"sync/atomic"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	buf := make(chan struct{}, pool)
	var wg sync.WaitGroup
	var ops int64
	var m sync.Mutex
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		buf <- struct{}{}
		go func() {
			m.Lock()
			user := getOne(int64(ops))
			atomic.AddInt64(&ops, 1)
			res = append(res, user)
			m.Unlock()
			<-buf
			wg.Done()
		}()
	}
	wg.Wait()
	close(buf)
	return res
}
