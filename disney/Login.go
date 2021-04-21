package disney

import (
	"errors"
	"fmt"
)

func (CruiseClient *CruiseClient) Login() error {

	res, err := CruiseClient.httpClient.Post(disneyLoginURL, "application/json", nil)

	if err != nil {
		return err
	}

	if res.Body == nil {
		return errors.New("no response body")
	}

	fmt.Println("Authed")

	return nil
}
