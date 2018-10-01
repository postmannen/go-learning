package rootpage

import (
	"fmt"
	"net/http"
)

//RootHandler is the handler for the root page
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root: This is the root page")
}
