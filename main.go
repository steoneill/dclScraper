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

	cruiseClient := disney.NewCruiseClient()

	err := cruiseClient.Login()

	if err != nil {
		panic(err)
	}

	for page <= totalPages {
		req := cruiseClient.CreateRequest(page)
		response, err := cruiseClient.GetCruises(req)

		fmt.Println(page)

		if err != nil {
			log.Fatalln(err)
		}

		totalPages = response.TotalPages

		for _, cruise := range response.Products {
			fmt.Println(cruise)

			Cruises = append(Cruises, cruise)
		}

		page++
	}

	fmt.Println(len(Cruises))

	disney.FindPort(Cruises)

}
