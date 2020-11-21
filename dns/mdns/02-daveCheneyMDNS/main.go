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

type aRecords struct {
	ARecords []ARecord `json:"records"`
}

type ARecord struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	TTL  string `json:"ttl"`
}

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

	// r1 := ARecord{
	// 	IP:   "10.0.0.26",
	// 	Name: "ws.local",
	// 	TTL:  "60",
	// }
	//
	// r2 := ARecord{
	// 	IP:   "10.0.0.1",
	// 	Name: "router.local",
	// 	TTL:  "60",
	// }
	//
	// r := []ARecord{r1, r2}

	for _, v := range records.ARecords {
		publishRecordA(v)
	}

	select {}
}
