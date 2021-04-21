package disney

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

type StateRoom struct {
	Accessible                 bool   `json:"accessible"`
	Available                  bool   `json:"available"`
	StateroomType              string `json:"stateroomType"`
	StateroomSubType           string `json:"stateroomSubType"`
	StateroomCategory          string `json:"stateroomCategory"`
	StateroomCategoryContentId string `json:"stateroomCategoryContentId"`
	StateroomSubTypeContentId  string `json:"stateroomSubTypeContentId"`
	Price                      Price  `json:"price"`
}

type Ship struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	SeawareId string `json:"seawareId"`
}

type Package struct {
	PackageID   string `json:"packageId"`
	PackageCode string `json:"packageCode"`
}

type Itinerary struct {
	Sailings    Sailing `json:"sailings"`
	ItineraryId string  `json:"itineraryId"`
}

type ItineraryResponse struct {
	Sailings    []Sailing `json:"sailings"`
	ItineraryId string    `json:"itineraryId"`
}

type Cruise struct {
	Name        string      `json:"productName"`
	ProductId   string      `json:"productId"`
	Itineraries []Itinerary `json:"itineraries"`
}

type Price struct {
	RateType         string         `json:"rateType"`
	PromoCode        string         `json:"promoCode"`
	BreakdownByGuest []PriceSummary `json:"breakdownByGuest"`
	Summary          PriceSummary   `json:"summary"`
}

type PriceSummary struct {
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

func (CruiseClient *CruiseClient) GetCruises(data CruiseData) (response Response, err error) {
	j, err := json.Marshal(data)

	var responseObject Response

	if err != nil {
		log.Fatalln(err)
	}

	res, err := CruiseClient.httpClient.Post(disneyDataURL, "application/json", bytes.NewReader(j))
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
