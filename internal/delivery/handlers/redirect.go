package handlers

import (
	"fmt"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
