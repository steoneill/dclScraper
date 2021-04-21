package disney

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type Cruise struct {
	Name        string `json:"productName"`
	Available   bool   `json:"available"`
	Itineraries []struct {
		Sailings []struct {
			Available     bool `json:"available"`
			TravelParties struct {
				Party []struct {
					Price struct {
						Summary struct {
							AgeGroup    interface{} `json:"ageGroup"`
							Currency    string      `json:"currency"`
							Subtotal    float64     `json:"subtotal"`
							Tax         float64     `json:"tax"`
							TaxIncluded bool        `json:"taxIncluded"`
							Total       float64     `json:"total"`
						} `json:"summary"`
					} `json:"price"`
				} `json:"0"`
			} `json:"travelParties"`
		} `json:"sailings"`
		MinimumPriceSummary struct {
			AgeGroup    interface{} `json:"ageGroup"`
			Currency    string      `json:"currency"`
			Subtotal    float64     `json:"subtotal"`
			Tax         float64     `json:"tax"`
			TaxIncluded bool        `json:"taxIncluded"`
			Total       float64     `json:"total"`
		} `json:"minimumPriceSummary"`
	} `json:"itineraries"`
}

type Price struct {
	AgeGroup    interface{} `json:"ageGroup"`
	Currency    string      `json:"currency"`
	Subtotal    float64     `json:"subtotal"`
	Tax         float64     `json:"tax"`
	TaxIncluded bool        `json:"taxIncluded"`
	Total       float64     `json:"total"`
}

type Response struct {
	TotalPages int      `json:"totalPages"`
	Products   []Cruise `json:"products"`
}

func (c *CruiseClient) GetCruises(data CruiseData) (response Response, err error) {
	j, err := json.Marshal(data)

	var responseObject Response

	if err != nil {
		log.Fatalln(err)
	}

	res, err := c.httpClient.Post(disneyDataURL, "application/json", bytes.NewReader(j))
	if err != nil {
		log.Fatalln(err)
	}

	if res.Body == nil {
		return responseObject, errors.New("No response body")
	}
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return responseObject, err
	}

	err = json.Unmarshal(body, &responseObject)

	return responseObject, nil

}
