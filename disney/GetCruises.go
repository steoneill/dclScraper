package disney

import (
	"bytes"
	json2 "encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type Cruise struct {
	Name      string  `json:"productName"`
	Available bool    `json:"available"`
	Price     []Price `json:"minimumPriceSummary"`
}

type Price struct {
	AgeGroup    interface{} `json:"ageGroup"`
	Currency    string      `json:"currency"`
	Subtotal    float64     `json:"subtotal"`
	Tax         float64     `json:"tax"`
	TaxIncluded bool        `json:"taxIncluded"`
	Total       float64     `json:"total"`
}

func (c *CruiseClient) GetCruises(data CruiseData) (cruises []Cruise, err error) {
	json, err := json2.Marshal(data)

	var Cruises []Cruise

	if err != nil {
		log.Fatalln(err)
	}

	res, err := c.httpClient.Post(disneyDataURL, "application/json", bytes.NewReader(json))
	if err != nil {
		log.Fatalln(err)
	}

	if res.Body == nil {
		return nil, errors.New("No response body")
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	fmt.Println("BODY")
	fmt.Println(string(body))
	return Cruises, nil

}
