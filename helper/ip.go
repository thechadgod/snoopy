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

type IP struct {
	Address  string `json:"ip"`
	Location string `json:"location"`
	ISP      string `json:"isp"`
}

type IpApiData struct {
	Status string `json:"status"`
	Query  string `json:"query"`

	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
	ISP     string `json:"isp"`
	Org     string `json:"org"`
	AS      string `json:"as"`

	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`

	Timezone    string `json:"timezone"`
	ZIP         string `json:"zip"`
	CountryCode string `json:"countryCode"`
	RegionName  string `json:"regionName"`
}

func GetPublicIP() (IP, error) {

	res, err := http.Get("http://wtfismyip.com/json")

	if err != nil {
		return IP{}, err
	}

	defer res.Body.Close()

	var body yourFuckingIP

	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return IP{}, err
	}

	publicIp := IP{
		Address:  body.YourFuckingIPAddress,
		Location: body.YourFuckingLocation,
		ISP:      body.YourFuckingISP,
	}

	return publicIp, nil

}

func TraceIP(ip string) (IP, error) {

	url := "http://ip-api.com/json/" + ip

	res, err := http.Get(url)

	if err != nil {
		return IP{}, err
	}
	defer res.Body.Close()

	var body IpApiData

	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return IP{}, err
	}

	publicIp := IP{
		Address:  body.Query,
		Location: body.City + ", " + body.Region + ", " + body.Country,
		ISP:      body.ISP,
	}

	return publicIp, nil

}
