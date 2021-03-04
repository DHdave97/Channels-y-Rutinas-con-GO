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

	/* for i := 0; i < len(links); i++ { //limitado al tamaño del slice
		fmt.Println(<-c)
	} */

	/* for { //bucle infinito
		go checkLink(<-c, c)
	} */

	/* for l := range c { //bucle segun el channel
		go checkLink(l, c)
	} */

	for l := range c { //bucle segun el channel
		go func(link string) {
			time.Sleep(2 * time.Second) //esperar 2 segundos
			checkLink(link, c)
		}(l) // funcion literal
	}

}
func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "algo anda mal ...!!!")
		c <- link
		return
	}

	fmt.Println(link, "Todo chevere. :D")
	c <- link
}
