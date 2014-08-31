/*
Package toggles does not do much yet
*/
package toggles

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println('Handler!')
	r.ParseForm()
	// fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hey Ian, I love this path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Hey Ian, I love this form: %s\n", r.Form)
}

func getToggle(s string) {
	fmt.Println("GetToggle: ", s)
}
