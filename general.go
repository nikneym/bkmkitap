package bkmkitap

import (
	"encoding/json"
	"net/http"
)

func GetOffers() (*Bundle, error) {
	resp, err := http.Get(getOffersEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bundle := new(Bundle)
	err = json.NewDecoder(resp.Body).Decode(bundle)
	if err != nil {
		return nil, err
	}

	return bundle, nil
}

func GetCountries() (*Region, error) {
	resp, err := http.Get(getCountriesEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	region := new(Region)
	err = json.NewDecoder(resp.Body).Decode(region)
	if err != nil {
		return nil, err
	}

	return region, nil
}
