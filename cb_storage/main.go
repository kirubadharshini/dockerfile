package function

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	counter int
	mu      sync.Mutex // To prevent race conditions
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	count := counter
	mu.Unlock()

	fmt.Fprintf(w, "Hello, World! This function has run %d times.", count)
}
