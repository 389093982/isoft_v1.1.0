package temp

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		fmt.Println("start to execute dataService temp put method...")
		put(w, r)
		return
	}
	if m == http.MethodPatch {
		fmt.Println("start to execute dataService temp patch method...")
		patch(w, r)
		return
	}
	if m == http.MethodPost {
		fmt.Println("start to execute dataService temp post method...")
		post(w, r)
		return
	}
	if m == http.MethodDelete {
		fmt.Println("start to execute dataService temp del method...")
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
