package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const defaultParallelValue = 10

// makeParallelRequests sends requests in a concurrent manner up to the capacity of the limitingChan
// After Golang 1.15, goroutines are inherently run in parallel
// and the number of parallel goroutines is determined by the value of
// GOMAXPROCS variable. The value of GOMAXPROCS is set to be the same as the
// number of processors of the machine. After the requests are made up to the
// parallel limit, they are always concurrent.

func makeHttpRequest(url string, limitingChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Sending signal to the limitingChannel to add one to the limit.
	// It will block calls until there is more room for the requests
	limitingChan <- 1

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error while making HTTP GET Request: ", err.Error())
		// In case of error, read from limitingChannel to make room for blocked
		// goroutines to start,
		<-limitingChan
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		<-limitingChan
		return
	}

	md5HashOfResponse := md5.Sum(bodyBytes)
	fmt.Println(fmt.Sprintf("%s %x", url, md5HashOfResponse))

	// Read from limitingChannel to make room for blocked
	// goroutines to start,
	<-limitingChan

}
func makeParallelRequests(urls []string, concurrencyLimit int) {

	// This is a bufferred channel which will block the number of requests to
	// the concurrency limit.
	limitingChannel := make(chan int, concurrencyLimit)
	defer func() {
		close(limitingChannel)
	}()

	var wg sync.WaitGroup
	for _, url := range urls {
		// For each url, add 1 to the waitGroup inorder to wait for it to complete
		wg.Add(1)
		// Calling makeHttpRequests funstion as a goroutine to make the HTTP call
		go makeHttpRequest(url, limitingChannel, &wg)
	}
	wg.Wait()
}
func main() {

	// Defining flag for parallelFlag value. If it is not provided as a flag argument, defaultParallelValue is set as
	// the value for parallelFlag
	parallelFlag := flag.Int("parallel", defaultParallelValue, "number of parallel threads that can be initiated")

	// Parsing the flags provided in the command
	flag.Parse()

	// Urls are provided as the tail of unnamed arguments and therefore is read through the flag.Args() function
	urls := flag.Args()

	if len(urls) == 0 {
		log.Println("No url provided")
		return
	}
	makeParallelRequests(urls, *parallelFlag)
}
