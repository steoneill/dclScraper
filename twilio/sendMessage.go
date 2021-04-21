package twilio

import (
	"dclScraper/disney"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var (
	accountSid = os.Getenv("accountSid")
	authToken  = os.Getenv("authToken")
	urlStr     = "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
)

func SendMessage(data *disney.Cruise) {
	msgData := url.Values{}

	msgData.Set("To", os.Getenv("sendNumber"))
	msgData.Set("From", os.Getenv("receiveNumber"))

	sailDate := data.Itineraries[0].Sailings.SailDateFrom
	cost := fmt.Sprintf("%f", data.Itineraries[0].Sailings.TravelParties.Detail[0].Price.Summary.Total)

	msgBody := "Cruise " + data.Name + " is showing as available, setting off on " + sailDate + ". It will currently cost £" + cost + "."

	msgData.Set("Body", msgBody)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}

}
