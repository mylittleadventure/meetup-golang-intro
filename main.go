package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/mylittleadventure/example-golang/scraper"
	"bitbucket.org/mylittleadventure/example-golang/server"
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
