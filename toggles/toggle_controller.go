/*
Package toggles does not do much yet
 */
package toggles

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey Ian, I love this path: %s!", r.URL.Path[1:])
}

func getToggle(s string) {
	fmt.Println("GetToggle: ", s)
}

