package disney

func (CruiseClient *CruiseClient) CreateRequest(page int, itineraryId string, productId string) CruiseData {

	req := CruiseData{
		Currency:     "GBP",
		Filters:      make([]string, 0),
		PartyMix:     []PartyMix{{false, 2, 1, []NonAdultAges{{2, "YEAR"}}, "0", "Child", true}},
		Region:       "INTL",
		StoreId:      "DCL",
		Affiliations: make([]string, 0),
		Page:         page,
		PageHistory:  false,
		ItineraryId:  itineraryId,
		ProductId:    productId,
	}

	return req

}
