package params

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/olshevskiy87/goradiooooo/utils"
	"github.com/parnurzeal/gorequest"
)

const (
	URL_COUNTRY = "https://radiooooo.app/country/mood"
)

func canonicalizeCountries(pCountries []string) ([]string, error) {
	var (
		countriesMap = make(map[string]bool)
		countries    []string
	)
	for _, country := range pCountries {
		var countryUpper = strings.ToUpper(country)
		if len(countryUpper) != 3 {
			return nil, fmt.Errorf("country iso-code \"%s\" is not a 3-letters string", country)
		}
		for _, c := range countryUpper {
			if c < 'A' || c > 'Z' {
				return nil, fmt.Errorf("country iso-code \"%s\" contains invalid characters", country)
			}
		}
		if _, ok := countriesMap[countryUpper]; !ok {
			countriesMap[countryUpper] = true
			countries = append(countries, countryUpper)
		}
	}
	return countries, nil
}

func GetAvailableCountries(pMoods []string, pDecade Decade) ([]string, error) {
	var moods map[string]bool
	if len(pMoods) == 0 {
		moods = map[string]bool{MOOD_SLOW: true, MOOD_FAST: true, MOOD_WEIRD: true}
	} else {
		canonMoods, err := canonicalizeMoods(pMoods)
		if err != nil {
			return nil, fmt.Errorf("could not canonicalize moods %v: %v", pMoods, err)
		}
		moods = make(map[string]bool, len(canonMoods))
		for _, m := range canonMoods {
			moods[m] = true
		}
	}

	requestAgent := gorequest.New().
		Timeout(10*time.Second).
		Set("User-Agent", utils.UserAgent)
	resp, respBody, errs := requestAgent.
		Get(URL_COUNTRY).
		Query(fmt.Sprintf("decade=%d", pDecade)).
		End()
	if errs != nil {
		return nil, fmt.Errorf("could not perform request with url \"%s\": %v", URL_COUNTRY, errs)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"could not perform request with url \"%s\" (status %d): %v",
			URL_COUNTRY,
			resp.StatusCode,
			errs,
		)
	}
	var respStruct map[string]interface{}
	err := json.Unmarshal([]byte(respBody), &respStruct)
	if err != nil {
		return nil, fmt.Errorf("could not parse json with countries: %v", err)
	}

	countriesMap := make(map[string]bool)

	for keyMood, valCountries := range respStruct {
		if _, ok := moods[keyMood]; !ok {
			continue
		}
		for _, c := range valCountries.([]interface{}) {
			cStr := c.(string)
			countriesMap[cStr] = true
		}
	}
	countries := make([]string, 0, len(countriesMap))
	for key := range countriesMap {
		countries = append(countries, key)
	}
	return countries, nil
}
