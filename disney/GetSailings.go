package disney

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

func (CruiseClient *CruiseClient) GetSailings(cruise *Cruise) (response ItineraryResponse, err error) {

	req := CruiseClient.CreateRequest(1, cruise.Itineraries[0].ItineraryId, cruise.ProductId)

	j, err := json.Marshal(req)

	var responseObject ItineraryResponse

	if err != nil {
		log.Fatalln(err)
	}

	res, err := CruiseClient.httpClient.Post(disneySailingsURL, "application/json", bytes.NewReader(j))

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
