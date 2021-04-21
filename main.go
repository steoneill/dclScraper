package main

import (
	"dclScraper/disney"
	"fmt"
	"log"
)

var (
	page       = 1
	totalPages = 100
)

func main() {
	fmt.Println("Starting app")

	cruiseClient := disney.NewDisneyClient()

	err := cruiseClient.Login()

	if err != nil {
		panic(err)
	}

	while page < totalPages {

	}

	req := cruiseClient.CreateRequest(page)

	cruises, err := cruiseClient.GetCruises(req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cruises)

}
