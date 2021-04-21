package main

import (
	"dclScraper/disney"
	"dclScraper/twilio"
	"fmt"
	"log"
	"strings"
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
		req := cruiseClient.CreateRequest(page, "", "")
		response, err := cruiseClient.GetCruises(req)
		if err != nil {
			log.Fatalln(err)
		}

		totalPages = response.TotalPages

		for _, cruise := range response.Products {
			itinerary, err := cruiseClient.GetSailings(&cruise)

			if err != nil {
				log.Println(err)
			}

			for _, sailing := range itinerary.Sailings {
				cruise.Itineraries[0].Sailings = sailing
				Cruises = append(Cruises, cruise)
			}
		}

		page++
	}

	for i, cruise := range Cruises {
		if strings.Contains(cruise.Name, "Liverpool") || strings.Contains(cruise.Name, "Newcastle") {
			twilio.SendMessage(&Cruises[i])
			return
		}
	}
}
