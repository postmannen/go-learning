package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	gitRepo := flag.String("gitRepo", "", "Your git repo")
	flag.Parse()

	//s := "https://github.com/RaaLabs/mixwebserver.git"

	ss := strings.Split(*gitRepo, "/")
	name := ss[len(ss)-1]

	name = strings.TrimSuffix(name, ".git")
	fmt.Println(name)
}
