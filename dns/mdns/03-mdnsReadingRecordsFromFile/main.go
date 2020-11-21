package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/davecheney/mdns"
)

func mustPublish(rr string) {
	if err := mdns.Publish(rr); err != nil {
		log.Fatalf(`Unable to publish record "%s": %v`, rr, err)
	}
}

// aRecords structure
type aRecords struct {
	ARecords []ARecord `json:"records"`
}

// ARecord structure
type ARecord struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	TTL  string `json:"ttl"`
}

// printReverseIP will just print the ip address in reverse order
func (r *ARecord) printReverseIP() string {
	sp := strings.Split(r.IP, ".")

	for left, right := 0, len(sp)-1; left < right; left, right = left+1, right-1 {
		sp[left], sp[right] = sp[right], sp[left]
	}

	s := concatenateSlice(sp)

	return s
}

// concatenateSlice will take all the string elements of
// a slice, and return them as a single string.
func concatenateSlice(s []string) string {
	var output string
	for _, v := range s {
		output += v
	}

	return output
}

// publishRecordA Publish an A record
func publishRecordA(r ARecord) {
	mustPublish(r.Name + ". " + r.TTL + " IN A " + r.IP)
	mustPublish(r.printReverseIP() + ".in-addr.arpa. " + r.TTL + " IN PTR " + r.Name + ".")
}
func main() {
	fh, err := os.Open("./recordsA.json")
	if err != nil {
		log.Printf("error: os.Open failed: %v\n", err)
		return
	}

	var records aRecords

	js := json.NewDecoder(fh)
	err = js.Decode(&records)
	if err != nil {
		log.Printf("error: json.Decode failed: %v\n", err)
		return
	}

	for _, v := range records.ARecords {
		publishRecordA(v)
	}

	select {}
}
