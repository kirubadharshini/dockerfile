package function

import (
	"fmt"
	"net/http"
)

// Entry point for the Cloud Function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World! storage")
}
