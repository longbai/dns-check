package main

import (
	"errors"
	"net"
	"time"

	"github.com/miekg/dns"
)

func Query(d DnsServer, domain string) (ip net.IP, err error) {
	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeA)

	c := new(dns.Client)
	c.DialTimeout = 60 * time.Second
	c.ReadTimeout = 60 * time.Second
	c.WriteTimeout = 60 * time.Second
	r, _, err := c.Exchange(m, d.Address())
	if err != nil {
		return
	}

	if r.Rcode != dns.RcodeSuccess {
		err = errors.New("failed response")
		return
	}

	for _, k := range r.Answer {
		if a, ok := k.(*dns.A); ok {
			ip = a.A
			return
		}
	}
	err = errors.New("no A")
	return
}
