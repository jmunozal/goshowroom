package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main () {

	url := flag.String("url", "", "url")
	t := flag.Int("delay", 0, "time (in seconds) between calls (>1)")
	token := flag.String("token", "", "auth (Bearer) token")

	flag.Parse()

	if *url == "" {
		fmt.Println("url is mandatory")
		os.Exit(1)
	}
	if *t == 0 || *t < 1 {
		fmt.Println("time is mandatory and > 1")
		os.Exit(1)
	}
	if *token == "" {
		fmt.Println("token is mandatory")
		os.Exit(1)
	}

	sleepTime := strconv.Itoa(*t) + "s"

	for {
		fmt.Println("call to ", *url, doCall(url, token))
		t, _ := time.ParseDuration(sleepTime)
		time.Sleep(t)
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