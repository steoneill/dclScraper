package twilio

import (
	"dclScraper/disney"
	"fmt"
	"github.com/sfreiberg/gotwilio"
	"log"
	"os"
)

var (
	accountSid = os.Getenv("accountSid")
	authToken  = os.Getenv("authToken")
)

func SendMessage(data *disney.Cruise) {

	twilio := gotwilio.NewTwilioClient(accountSid, authToken)

	from := os.Getenv("sendNumber")
	to := os.Getenv("receiveNumber")

	sailDate := data.Itineraries[0].Sailings.SailDateFrom
	cost := fmt.Sprintf("%f", data.Itineraries[0].Sailings.TravelParties.Detail[0].Price.Summary.Total)

	msgBody := "Cruise " + data.Name + " is showing as available, setting off on " + sailDate + ". It will currently cost £" + cost + "."

	r, _, err := twilio.SendSMS(from, to, msgBody, "", "")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(r)
}
