package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func get(ctx context.Context, id int) (string, error) {
	// пофиг что тут
	return fmt.Sprintf("result-%d", id), nil
}

// listToCall := []int{
//        1,2,3//... 1_000_000 элементов'
//    }

// Задание 4 вызывать не больше 4х запросов одновременно и не больше 1000 в минуту (leacky bucket rate-limit)
func callAll(listToCall []int) map[int]string {

	const rate = 4
	const period = time.Minute / 1000
	results := map[int]string{}
	var mutex sync.Mutex
	var wg sync.WaitGroup

	semaphore := make(chan struct{}, rate)
	limiter := time.NewTicker(period)
	defer limiter.Stop()

	for _, id := range listToCall {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			<-limiter.C

			semaphore <- struct{}{}

			defer func() {
				<-semaphore
			}()

			s, err := get(context.Background(), id)
			if err != nil {
				fmt.Println("error", id, err)
				return
			}
			mutex.Lock()
			results[id] = s
			mutex.Unlock()
		}(id)

	}
	wg.Wait()
	return results
}

func main() {
	listToCall := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	results := callAll(listToCall)

	for id, result := range results {
		fmt.Printf("ID: %d, Result: %s\n", id, result)
	}
}
