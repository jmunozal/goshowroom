package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

func main() {

	var (
		flDomain      = flag.String("domain", "", "The domain to perform guessing against.")
		flWordList    = flag.String("wordlist", "", "The wordlist to use for guessing.")
		flWorkerCount = flag.Int("c", 100, "The amount of workers to use.")
		//flServerAddr	= flag.String("server", "8.8.8.8:53", "The DNS server to use.")
		dnsResolvers = []string{"208.67.222.222:53", "1.1.1.1:53", "8.8.8.8:53", "8.8.4.4:53", "9.9.9.9:53"}
	)
	flag.Parse()

	if *flDomain == "" || *flWordList == "" {
		fmt.Println("-domain and -wordlist are required")
		os.Exit(1)
	}

	go spinner(100 * time.Millisecond)

	var results []result
	fqdns := make(chan string, *flWorkerCount)
	gather := make(chan []result)
	tracker := make(chan empty)

	fh, err := os.Open(*flWordList)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	scanner := bufio.NewScanner(fh)

	for i := 0; i < *flWorkerCount; i++ {
		flServerAddr := getResolver(&dnsResolvers)
		//		log.Println("Resolver: ", *flServerAddr)
		go worker(tracker, fqdns, gather, flServerAddr)
	}

	for scanner.Scan() {
		fqdns <- fmt.Sprintf("%s.%s", scanner.Text(), *flDomain)
	}

	go func() {
		for r := range gather {
			w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
			results = append(results, r...)
			for _, res := range r {
				fmt.Fprintf(w, "\r %s \t %s\n", res.IPAddress, res.Hostname)
			}
		}
		var e empty
		tracker <- e
	}()

	close(fqdns)

	for i := 0; i < *flWorkerCount; i++ {
		<-tracker
	}
	close(gather)
	<-tracker

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, ' ', 0)
	for _, r := range results {
		fmt.Fprintf(w, "\r %s \t %s\n", r.Hostname, r.IPAddress)
	}
	w.Flush()
}

type result struct {
	IPAddress string
	Hostname  string
}

func lookupA(fqdn string, serverAddr *string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(&m, *serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		errorText := "no answer server " + *serverAddr
		return ips, errors.New(errorText)
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			ips = append(ips, a.A.String())
			return ips, nil
		}
	}
	return ips, nil
}

func lookupCNAME(fqdn string, serverAddr *string) ([]string, error) {
	var m dns.Msg
	var ips []string
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeCNAME)
	in, err := dns.Exchange(&m, *serverAddr)
	if err != nil {
		return ips, err
	}
	if len(in.Answer) < 1 {
		errorText := "no answer server " + *serverAddr
		return ips, errors.New(errorText)
	}
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.CNAME); ok {
			ips = append(ips, a.Target)
		}
	}
	return ips, nil
}

func lookup(fqdn string, serverAddr *string) []result {
	var results []result
	var cfqdn = fqdn
	for {
		cnames, err := lookupCNAME(fqdn, serverAddr)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue // process the next cname
		}
		// processing an A register
		ips, err := lookupA(cfqdn, serverAddr)
		if err != nil {
			break // error A register
		}
		// process IPs
		for _, ip := range ips {
			results = append(results, result{IPAddress: ip, Hostname: fqdn})
		}
		break // end process
	}
	return results
}

type empty struct{}

func worker(tracker chan empty, fqdns chan string, gather chan []result, serverAddr *string) {
	for fqdn := range fqdns {
		results := lookup(fqdn, serverAddr)
		if len(results) > 0 {
			gather <- results
		}
	}
	var e empty
	tracker <- e
}

func getResolver(d *[]string) *string {
	return &(*d)[rand.Intn(len(*d))]
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
