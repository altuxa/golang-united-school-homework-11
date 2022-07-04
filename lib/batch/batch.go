package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

// SOLUTION 1

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, 0, n)
	sem := make(chan struct{}, pool)
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(num int) {
			user := getOne(int64(num))
			m.Lock()
			res = append(res, user)
			m.Unlock()
			<-sem
			wg.Done()
		}(i)
	}
	wg.Wait()
	close(sem)
	return res
}

// SOLUTION 2

// func getBatch(n int64, pool int64) (res []user) {
// 	res = make([]user, 0, n)
// 	wait := make(chan user, n)
// 	buf := make(chan struct{}, pool)
// 	for i := 0; i < int(n); i++ {
// 		buf <- struct{}{}
// 		go func(num int) {
// 			user := getOne(int64(num))
// 			<-buf
// 			wait <- user
// 		}(i)
// 	}
// 	close(buf)
// 	for i := 0; i < int(n); i++ {
// 		res = append(res, <-wait)
// 	}
// 	close(wait)
// 	return res
// }
