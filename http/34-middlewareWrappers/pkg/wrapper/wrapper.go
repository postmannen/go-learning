package wrapper

import (
	"fmt"
	"net/http"
)

//MyWrapper is for wrapping some other function into this one,
//and return them both combined.
func MyWrapper(hf http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Wrapper: Lets pretend we did some authorization checking here",
			" before the real rootHandler is called")

		//Then we execute the passed in HandlerFunc, and it will be passed into W and R
		hf(w, r)
	}
}
