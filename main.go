package main

import (
	"fmt"
	"net/http"
)

func ping(url string, c chan string) {
	_, err := http.Get(url)

	if err == nil {
		fmt.Println(url + " it's OK!")
	} else {
		fmt.Println(url + " it's not working")
	}
	close(c)
}

func main() {
	c := make(chan string)
	go ping("http://www.google.com", c)

	result, _ := <-c
	fmt.Println(result)
}
