package disney

func (c *CruiseClient) CreateRequest(page int) CruiseData {

	req := CruiseData{
		Currency:     "GBP",
		Filters:      make([]string, 0),
		PartyMix:     []PartyMix{{false, 2, 1, []NonAdultAges{{2, "YEAR"}}, "0", "Child", true}},
		Region:       "INTL",
		StoreId:      "DCL",
		Affiliations: make([]string, 0),
		Page:         page,
		PageHistory:  false,
	}

	return req

}
