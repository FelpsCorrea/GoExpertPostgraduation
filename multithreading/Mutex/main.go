package main

import (
	"fmt"
	"net/http"
	"sync"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++

		// atomic.AddUint64(&number, 1) // Alternativa ao mutex

		m.Unlock()
		fmt.Fprintf(w, "Number: %d\n", number)
	})

	http.ListenAndServe(":8080", nil)
}

// go run -race main.go testa se o código possui race conditions
// Testa 10000 requisições com 100 concorrentes
// ab -n 10000 -c 100 http://localhost:8080/ testa a aplicação
