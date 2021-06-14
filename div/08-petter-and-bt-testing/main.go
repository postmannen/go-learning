package main

import (
	"fmt"
	"net/http"
	"sync"
)

func minSide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>Document</title>
		</head>
	<body>
		<h1> Dette er en test </h1>		
	</body>
	</html>
	`)
}

func minSide2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `
	<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<title>Document</title>
		</head>
	<body>
		<h1> Dette er en test </h1>		
	</body>
	</html>
	`)
}

func one() {
	http.HandleFunc("/a", minSide)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed starting server with err %s \n", err)
	}
	wg.Done()
}

func two() {
	http.HandleFunc("/b", minSide2)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Printf("failed starting server with err %s \n", err)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go one()

	wg.Add(1)
	go two()

	wg.Wait()

}
