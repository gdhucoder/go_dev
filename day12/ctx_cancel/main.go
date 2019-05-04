package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct {
	r   *http.Response
	err error
}

// timeout ctx can be used as a series of operation
// such as: io, db, or combo...

func process() {

	// define a timeout ctx
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("http request failed, err:", err)
		return
	}
	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		// Timeout! err: Get http://www.google.com: net/http: request canceled while waiting for connection
		fmt.Println("Timeout! err:", res.err)
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("Server Response: %s", out)
	}
	return
}
func main() {
	process()
}
