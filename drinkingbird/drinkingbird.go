package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main () {

	url := flag.String("url", "", "url")
	time := flag.Int("delay", 0, "time (in seconds) between calls (>1)")
	token := flag.String("token", "", "auth (Bearer) token")

	flag.Parse()

	if *url == "" {
		fmt.Println("url is mandatory")
		os.Exit(1)
	}
	if *time == 0 || *time < 1 {
		fmt.Println("time is mandatory and > 1")
		os.Exit(1)
	}
	if *token == "" {
		fmt.Println("token is mandatory")
		os.Exit(1)
	}

	for {
		fmt.Println("call to ", *url, doCall(url, token))
	}


}

func doCall (url *string, token *string) string {

	var client = &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ :=http.NewRequest("GET", *url, nil);
	btoken := "Bearer" + *token
	req.Header.Add("Authentication", btoken)
	response, err := client.Do(req)
	if (err != nil) {
		fmt.Println("Error calling URL")
	}
	return response.Status

}