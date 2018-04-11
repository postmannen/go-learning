package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

//template to represent a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

/*
	ServeHTTP handles the HTTP request.
	Here we are compiling the template inside a ServeHTTP method.
	Create a ServeHTTP method for the struct 'templateHandler'. This method is the same specified in the 'Handler' interface,
	so here we are also making our type 'templateHandler' accepted as a 'Handler' as specified in the interface by implementing the ServeHTTP method.
*/
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//make sure the compiling of the template is only called once.
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	//pass the request r as data into the template.
	t.templ.Execute(w, r)
}

func main() {

	/*
		The struct templateHandler is now also a type Handler since it satisfies the Handler interface by having a ServerHTTP method.
		We use the struct type directly by using & before the struct name, so we don't create a reference to it with a variable, which
		is not needed since it will be used only once right here in the code. We also fill the filname directly into the structs
		filename field.
		http.Handle wants a string for the url, and a type http.Handler.
		Since the struct 'templateHandler' have a method called ServeHTTP, that method will be executed when the struct is given
		as input to the http.Handle below.
	*/
	http.Handle("/", &templateHandler{filename: "chat.html"})

	//"addr" : the field to use as a parameter like --file=... , ":8080" is the default value
	addr := flag.String("addr", ":8080", "The addr of the application")
	flag.Parse()

	r := newRoom()
	//r which is our new room is also a http.Handler since it has a ServeHTTP method
	http.Handle("/room", r)

	//start the room
	go r.run()

	fmt.Println("Starting the webserver on : ", *addr)
	//we can bundle the output for checking directly in calling a function.
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("Error : ListenAndServe failed : ", err)
	}
}
