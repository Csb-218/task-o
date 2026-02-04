package main

import (
	"bytes"
	"net/http"
	"sync"
)

func test_concurrent() {
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			body := []byte(`{"title":"Task ` + string(rune('A'+i)) + `"}`)
			http.Post(
				"http://localhost:3000/todos",
				"application/json",
				bytes.NewBuffer(body),
			)
		}(i)
	}

	wg.Wait()
}
