package web

import (
	"fmt"
	"net/http"
)

//DataStruct data for web handler
type DataStruct struct {
	Data1 string
}

//IndexA for handling the A page
func (d *DataStruct) IndexA(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", d)
	d.Data1 = "Apen hoppet ned av busken"
}

//IndexB for handling the B page
func (d *DataStruct) IndexB(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", d)
	d.Data1 = "Apen hoppet opp i busken igjen"
}

//
