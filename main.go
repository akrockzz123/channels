package main

import (
	"fmt"
	"net/http"
	"time"
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

	/*for {

		go checkLink(<-c, c)
	}*/

	// alternative example for writing above code

	for l := range c { // waiting for channel to return link(l)

		//time.Sleep(5 * time.Second)

		//go checkLink(l, c)

		go func() {

			time.Sleep(5 * time.Second)

			checkLink(l, c)
		}()
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
