package disney

type StateRoom struct {
	Accessible                 bool    `json:"accessible"`
	Available                  bool    `json:"available"`
	StateroomType              string  `json:"stateroomType"`
	StateroomSubType           string  `json:"stateroomSubType"`
	StateroomCategory          string  `json:"stateroomCategory"`
	StateroomCategoryContentId string  `json:"stateroomCategoryContentId"`
	StateroomSubTypeContentId  string  `json:"stateroomSubTypeContentId"`
	Price                      []Price `json:"price"`
}

type Ship struct {
	Id             string      `json:"id"`
	Name           string      `json:"name"`
	SeawareId      string      `json:"seawareId"`
	StateRoomTypes []StateRoom `json:"stateroomTypes"`
}

type Package struct {
	PackageID   string `json:"packageId"`
	PackageCode string `json:"packageCode"`
}

type Sailing struct {
	SailingId      string          `json:"sailingId"`
	ProductId      string          `json:"productId"`
	ItineraryId    string          `json:"itineraryId"`
	Ship           Ship            `json:"ship"`
	SailDateFrom   string          `json:"sailDateFrom"`
	SailDateTo     string          `json:"sailDateTo"`
	Available      bool            `json:"available"`
	IsEarlyBooking bool            `json:"isEarlyBooking"`
	TravelParties  []TravelParties `json:"travelParties"`
	Package        Package         `json:"package"`
}

type TravelParties struct {
	TravelParties struct {
		Field1 []struct {
			Accessible                 bool   `json:"accessible"`
			Available                  bool   `json:"available"`
			StateroomType              string `json:"stateroomType"`
			StateroomSubType           string `json:"stateroomSubType"`
			StateroomCategory          string `json:"stateroomCategory"`
			StateroomCategoryContentId string `json:"stateroomCategoryContentId"`
			StateroomSubTypeContentId  string `json:"stateroomSubTypeContentId"`
			Price                      struct {
				RateType  string `json:"rateType"`
				PromoCode string `json:"promoCode"`
				Summary   struct {
					AgeGroup    interface{} `json:"ageGroup"`
					Currency    string      `json:"currency"`
					Subtotal    float64     `json:"subtotal"`
					Tax         float64     `json:"tax"`
					TaxIncluded bool        `json:"taxIncluded"`
					Total       float64     `json:"total"`
				} `json:"summary"`
				BreakdownByGuest struct {
					Field1 struct {
						AgeGroup    string  `json:"ageGroup"`
						Currency    string  `json:"currency"`
						Subtotal    float64 `json:"subtotal"`
						Tax         float64 `json:"tax"`
						TaxIncluded bool    `json:"taxIncluded"`
						Total       float64 `json:"total"`
					} `json:"1"`
					Field2 struct {
						AgeGroup    string  `json:"ageGroup"`
						Currency    string  `json:"currency"`
						Subtotal    float64 `json:"subtotal"`
						Tax         float64 `json:"tax"`
						TaxIncluded bool    `json:"taxIncluded"`
						Total       float64 `json:"total"`
					} `json:"2"`
					Field3 struct {
						AgeGroup    string  `json:"ageGroup"`
						Currency    string  `json:"currency"`
						Subtotal    float64 `json:"subtotal"`
						Tax         float64 `json:"tax"`
						TaxIncluded bool    `json:"taxIncluded"`
						Total       float64 `json:"total"`
					} `json:"3"`
				} `json:"breakdownByGuest"`
			} `json:"price"`
		} `json:"0"`
	} `json:"travelParties"`
}

func (c *CruiseClient) GetSailings(cruise *Cruise) {

}
