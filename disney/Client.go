package disney

import (
	"log"
	"net/http"
	"net/http/cookiejar"
)

type CruiseClient struct {
	hpts       string
	hptsh      string
	httpClient *http.Client
}

type CruiseData struct {
	Currency     string     `json:"currency"`
	Filters      []string   `json:"filters"`
	PartyMix     []PartyMix `json:"partyMix"`
	Region       string     `json:"region"`
	StoreId      string     `json:"storeId"`
	Affiliations []string   `json:"affiliations"`
	Page         int        `json:"page"`
	PageHistory  bool       `json:"pageHistory"`
	ItineraryId  string     `json:"itineraryId"`
	ProductId    string     `json:"productId"`
}

type NonAdultAges struct {
	Age     int    `json:"age"`
	AgeUnit string `json:"ageUnit"`
}

type PartyMix struct {
	Accessible      bool           `json:"accessible"`
	AdultCount      int            `json:"adultCount"`
	ChildCount      int            `json:"childCount"`
	NonAdultAges    []NonAdultAges `json:"nonAdultAges"`
	PartyMixId      string         `json:"partyMixId"`
	StateroomAction string         `json:"stateroomAction"`
	IsAddGuest      bool           `json:"isAddGuest"`
}

const (
	disneyLoginURL    = "https://disneycruise.disney.go.com/dcl-cruise-101-webapi/product-availability/authz/private"
	disneyDataURL     = "https://disneycruise.disney.go.com/dcl-cruise-101-webapi/product-availability/available-products/"
	disneySailingsURL = "https://disneycruise.disney.go.com/dcl-cruise-101-webapi/product-availability/available-sailings/"
)

func NewCruiseClient() *CruiseClient {
	jar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}

	c := &CruiseClient{}

	c.httpClient = &http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Jar:           jar,
		Timeout:       http.DefaultClient.Timeout,
	}

	return c
}
