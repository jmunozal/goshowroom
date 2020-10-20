package main

import (
	"fmt"
	"github.com/miekg/dns"
)

func main() {
	var msg dns.Msg
	fqdn := dns.Fqdn("stacktitan.com")
	msg.SetQuestion(fqdn, dns.TypeA)
	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		panic(err)
	}
	if len(in.Answer) < 1 {
		fmt.Println("No records")
		return
	}
	for _, answer := range in.Answer {
		// https://stackoverflow.com/questions/24492868/what-is-the-meaning-of-dot-parenthesis-syntax-in-golang
		if a, ok := answer.(*dns.A); ok {
			fmt.Println(a)
		}
	}
}