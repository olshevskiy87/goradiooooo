package params

import (
	"fmt"
	"strings"
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
