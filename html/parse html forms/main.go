import (
	"log"
	"net/http"
)

func billCreateWeb(w http.ResponseWriter, r *http.Request) {
	err := tmpl["init.html"].ExecuteTemplate(w, "createBillCompletePage", "SOME DATA HERE ????")
	if err != nil {
		log.Println("createBillCompletePage: template execution error = ", err)
	}
	r.ParseForm()
	//r.FormValue("value_of_name=XXXX>") to get the value of name=XXX. Returns only a single value
	//r.Form["<value_of_name=XXXX>"]	to get the value of name=XXX . Can get multiple values and stores them in a slice
	minInput := r.Form["knapp"]
	//log.Printf("billCreateWeb : New line button pressed, %v\n", minInput[0])

	//Check if a string of values are parsed from form. Will return runtime error if not checked since the slice is not created.
	if minInput != nil && minInput[0] == "verdi" {
		err := tmpl["init.html"].ExecuteTemplate(w, "createBillLine", "SOME DATA HERE ????")
		if err != nil {
			log.Println("createBillCompletePage: template execution error = ", err)
		}
	}
}

/* HTML Example for the code above
{{define "createBillLine"}}
    <form>
            <input type="text" name="tekstboks" value="tekstverdi">
        <input type="submit" name="knapp" value="verdi">
    </form>
{{end}
*/