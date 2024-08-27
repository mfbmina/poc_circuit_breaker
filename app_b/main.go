package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sony/gobreaker/v2"
)

func main() {
	var st gobreaker.Settings
	st.Name = "Circuit Breaker PoC"
	st.Timeout = time.Second * 5
	st.MaxRequests = 2
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		return counts.ConsecutiveFailures >= 1
	}

	cb := gobreaker.NewCircuitBreaker[int](st)

	url := "http://localhost:8080/success"
	cb.Execute(func() (int, error) { return Get(url) })
	fmt.Println("Circuit Breaker state:", cb.State(), cb.Counts().Requests) // closed!

	url = "http://localhost:8080/failure"
	for range 2 {
		cb.Execute(func() (int, error) { return Get(url) })
		fmt.Println("Circuit Breaker state:", cb.State(), cb.Counts().Requests) // still closed!
	}

	time.Sleep(time.Second * 6)
	url = "http://localhost:8080/success"
	cb.Execute(func() (int, error) { return Get(url) })
	fmt.Println("Circuit Breaker state:", cb.State()) // half-open!

	url = "http://localhost:8080/success"
	cb.Execute(func() (int, error) { return Get(url) })
	fmt.Println("Circuit Breaker state:", cb.State()) // closed!
}

func Get(url string) (int, error) {
	r, _ := http.Get(url)

	if r.StatusCode != http.StatusOK {
		return r.StatusCode, fmt.Errorf("failed to get %s", url)
	}

	return r.StatusCode, nil
}
