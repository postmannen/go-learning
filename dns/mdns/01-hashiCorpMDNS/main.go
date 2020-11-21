package main

import (
	"log"

	"github.com/hashicorp/mdns"
)

func main() {
	//host, _ := os.Hostname()
	host := "apekatt"
	info := []string{"My awesome service"}
	service, err := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8000, nil, info)
	if err != nil {
		log.Printf("NewMDNSService failed: %v\n", err)
	}

	config := &mdns.Config{
		Zone: service,
	}

	server, err := mdns.NewServer(config)
	if err != nil {
		log.Printf("NewServer failed: %v\n", err)
	}

	defer server.Shutdown()

	select {}
}
