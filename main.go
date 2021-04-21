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

	var Cruises []disney.Cruise

	cruiseClient := disney.NewDisneyClient()

	err := cruiseClient.Login()

	if err != nil {
		panic(err)
	}

	for page <= totalPages {
		req := cruiseClient.CreateRequest(page)

		response, err := cruiseClient.GetCruises(req)

		if err != nil {
			log.Fatalln(err)
		}

		if page == 2 {
			totalPages = response.TotalPages
		}

		for _, cruise := range response.Products {
			Cruises = append(Cruises, cruise)
		}

		fmt.Println(Cruises)

		page++
	}

}
