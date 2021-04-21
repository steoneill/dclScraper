package disney

type Sailing struct {
	SailingId string `json:"sailingId"`
	Ship      struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		SeawareId string `json:"seawareId"`
	} `json:"ship"`
	SailDateFrom  string `json:"sailDateFrom"`
	SailDateTo    string `json:"sailDateTo"`
	Available     bool   `json:"available"`
	TravelParties struct {
		Detail []StateRoom `json:"0"`
	} `json:"travelParties"`
}
