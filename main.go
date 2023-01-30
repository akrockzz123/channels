package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {

		go checkLink(link, c)
	}

	for {

		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {

		fmt.Println(link, "might not be active now")

		c <- link

		return
	}

	c <- link

	fmt.Println(link, "is up")
}
