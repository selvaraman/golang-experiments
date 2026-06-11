package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL  string
	Body string
	Err  error
}

func fetch(ch chan<- Result, url string, ctx context.Context, client *http.Client, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	res, err := client.Do(req)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		ch <- Result{URL: url, Err: fmt.Errorf("non-200 status code: %d", res.StatusCode)}
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	result := Result{URL: url, Body: string(body)}
	ch <- result
}

func main() {
	urls := []string{
		"https://fake-json-api.mock.beeceptor.com/users/1",
		"https://fake-json-api.mock.beeceptor.com/users/2",
		"https://fake-json-api.mock.beeceptor.com/users/3",
	}
	var wg sync.WaitGroup
	ch := make(chan Result, len(urls))
	client := &http.Client{Timeout: 5 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	wg.Add(len(urls))
	for _, url := range urls {
		go fetch(ch, url, ctx, client, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for msg := range ch {
		fmt.Printf("%+v\n", msg)
	}
}
