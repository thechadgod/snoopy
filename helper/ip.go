package helper

import (
	"encoding/json"
	"net/http"
)

type yourFuckingIP struct {
	YourFuckingIPAddress   string `json:"YourFuckingIPAddress"`
	YourFuckingLocation    string `json:"YourFuckingLocation"`
	YourFuckingHostname    string `json:"YourFuckingHostname"`
	YourFuckingISP         string `json:"YourFuckingISP"`
	YourFuckingTorExit     bool   `json:"YourFuckingTorExit"`
	YourFuckingCountryCode string `json:"YourFuckingCountryCode"`
}

type PublicIP struct {
	Address  string `json:"ip"`
	Location string `json:"location"`
	ISP      string `json:"isp"`
}

func GetPublicIP() (PublicIP, error) {

	res, err := http.Get("http://wtfismyip.com/json")

	if err != nil {
		return PublicIP{}, err
	}

	defer res.Body.Close()

	var body yourFuckingIP

	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return PublicIP{}, err
	}

	publicIp := PublicIP{
		Address:  body.YourFuckingIPAddress,
		Location: body.YourFuckingLocation,
		ISP:      body.YourFuckingISP,
	}

	return publicIp, nil

}
