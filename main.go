package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type Response struct {
	Urls []struct {
		Url string `json:"url"`
	} `json:"url_list"`
	HasNext bool `json:"has_next"`
}

func main() {

	var concurrency int
	flag.IntVar(&concurrency, "c", 30, "Set Concurrency or Speed")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i <= concurrency; i++ {
		wg.Add(1)
		go func() {
			otx()
			wg.Done()
		}()
		wg.Wait()
	}

}

func otx() {
	//time.Sleep(time.Microsecond * 10)
	var url_list = make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urls := scanner.Text()
		u, err := url.Parse(urls)
		if err != nil {
			return
		}
		url_list = append(url_list, u.Hostname())

	}

	for _, domain := range url_list {
		var result Response
		var next_page = 1
		for {
			format_url := fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/url_list?limit=100&page=%v", domain, next_page)
			// Get request
			resp, err := http.Get(format_url)
			if err != nil {
				fmt.Println("No response from request")
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body) // response body is []byte

			if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
				fmt.Println("Can not unmarshal JSON")
			}
			// Loop through all the urls
			for _, rec := range result.Urls {
				fmt.Println(rec.Url)
			}
			if result.HasNext == true {
				next_page++
			} else {
				break
			}

		}
	}
}
