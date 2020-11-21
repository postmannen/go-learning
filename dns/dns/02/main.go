package main

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/miekg/dns"
)

var domainsToAddresses map[string]string = map[string]string{
	"nas.erter.local.":  "10.0.0.20",
	"a.realfarmer.org.": "1.2.3.4",
	"c.realfarmer.org.": "erter.org.",
}

type handler struct{}

func (d *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	fmt.Printf("*** %#v\n", r.Question[0])
	switch r.Question[0].Qtype {
	case dns.TypeA:
		log.Printf("case: Arecord: r.Question = %#v\n", r.Question)
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := domainsToAddresses[domain]

		// net.ParseIP will return nil if it is not an ip,
		// and we can use that for checking later to use
		// cname or a-record
		ipaddr := net.ParseIP(address)

		if ok {
			if ipaddr != nil {
				msg.Answer = append(msg.Answer, &dns.A{
					Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
					A:   ipaddr,
				})
			}

			if ipaddr == nil {
				msg.Answer = append(msg.Answer, &dns.CNAME{
					Hdr: dns.RR_Header{
						Name:   domain,
						Rrtype: dns.TypeCNAME,
						Class:  dns.ClassINET,
						Ttl:    60,
					},
					Target: address,
				})
			}
		}

	}
	w.WriteMsg(&msg)
}

func main() {
	srv := &dns.Server{Addr: ":" + strconv.Itoa(53), Net: "udp"}
	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
