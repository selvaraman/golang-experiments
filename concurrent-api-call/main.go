package main

import (
	"fmt"
	"net/http"
	"context"
	"time"
	"io"
)

func fetch(url string, ctx context.Context, client *http.Client) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}

func main() {
	urls := []string{
		"https://fake-json-api.mock.beeceptor.com/users/1",
		"https://fake-json-api.mock.beeceptor.com/users/2",
		"https://fake-json-api.mock.beeceptor.com/users/3",
	}
	client := &http.Client{Timeout: 5 * time.Second}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	for _, url := range urls {
		fetch(url, ctx, client)
	}
}
