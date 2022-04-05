package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func ping(url string, c chan string) {
	_, err := http.Get(url)

	if err == nil {
		c <- url + " it's OK!"
	} else {
		c <- ""
	}
}

func main() {
	for {
		c := make(chan string)

		urls := [2]string{"http://www.error.error.kl", "http://twitter.com"}

		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(2)

		go ping("http://www.google.com", c)
		go ping("http://www.facebook.com", c)
		go ping(urls[i], c)

		google, facebook, br := <-c, <-c, <-c

		if (google != "") && (facebook != "") && (br != "") {
			fmt.Println("GOOD TO GO!")
		} else {
			fmt.Println("Something went wrong")
		}

		fmt.Println("")

		time.Sleep(100 * time.Millisecond)
	}
}
