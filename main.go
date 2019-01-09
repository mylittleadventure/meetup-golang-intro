package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/mylittleadventure/meetup-golang-intro/scraper"
	"github.com/mylittleadventure/meetup-golang-intro/server"
)

func main() {
	pages, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing pages number: " + err.Error())
		pages = 3
	}

	client := &http.Client{}

	server.Serve(scraper.ScrapeBooking(client, pages))
}
