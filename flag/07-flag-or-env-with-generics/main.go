package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CheckEnv[T any](key string, v T) any {
	val, ok := os.LookupEnv(key)
	if !ok {
		return v
	}

	switch any(v).(type) {
	case int:
		n, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("error: failed to convert env to int: %v\n", n)
		}
		return n
	case string:
		return val
	}

	return nil
}

func getConfig(fs *flag.FlagSet) {
	fs.VisitAll(func(f *flag.Flag) {
		fmt.Printf("printing flag values, name: %v, value:%v\n", f.Name, f.Value)
	})

}

type config struct {
	httpAddr    string
	httpTimeout int
}

func main() {
	c := config{
		httpAddr:    ":8080",
		httpTimeout: 10,
	}

	flag.StringVar(&c.httpAddr, "http-listen-addr", CheckEnv("HTTP_ADDR", c.httpAddr).(string), "http service listen address")
	flag.IntVar(&c.httpTimeout, "http-timeout", CheckEnv("HTTP_TIMEOUT", c.httpTimeout).(int), "http timeout requesting http services")

	flag.Parse()
	getConfig(flag.CommandLine)

	fmt.Printf("printing struct: %#v\n", c)
}
