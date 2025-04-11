package main

import (
	"context"
	"fmt"
	"sync"
)

func get(ctx context.Context, id int) (string, error) {
	// пофиг что тут
	return fmt.Sprintf("result-%d", id), nil
}

// listToCall := []int{
//        1,2,3//... 1_000_000 элементов'
//    }

// Задание 3 добавить Mutex
// Задание 4 вызывать не больше 4х запросов одновременно и не больше 1000 в минуту (leacky bucket rate-limit)
func callAll(listToCall []int) map[int]string {
	results := map[int]string{}
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, id := range listToCall {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
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
	listToCall := []int{1, 2, 3, 4, 5}

	results := callAll(listToCall)

	for id, result := range results {
		fmt.Printf("ID: %d, Result: %s\n", id, result)
	}
}
