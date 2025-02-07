package development

import (
	"main/internal/address"
	"main/internal/city"
	"main/internal/country"
	"math/rand"
	"strings"
	"time"
)

func NewCountry(countryName string) *country.Country {
	return &country.Country{
		Country: countryName,
	}
}

func NewCity(cityName string) *city.City {
	return &city.City{
		City: cityName,
	}
}

func NewAddress() *address.Address {
	return &address.Address{}
}

func generateRandomString(stringLength int) string {
	rand.NewSource(time.Now().UnixNano())

	var sb strings.Builder
	length := rand.Intn(stringLength)

	if length > 0 {
		firstChar := string(rune(rand.Intn(26) + 65))
		sb.WriteString(firstChar)

		for i := 1; i < length; i++ {
			randomChar := string(rune(rand.Intn(26) + 97))
			sb.WriteString(randomChar)
		}
	}

	return sb.String()
}
