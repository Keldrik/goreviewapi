package lib

import (
	"net/http"
	"strconv"
)

func GetLengthQuery(r *http.Request) int {
	lengthString := r.URL.Query().Get("length")
	if lengthString == "" {
		return 6
	}
	lengthNumber, err := strconv.Atoi(lengthString)
	if err != nil {
		return 6
	}
	return lengthNumber
}

func GetListingQuery(r *http.Request) int {
	listingString := r.URL.Query().Get("listing")
	if listingString == "" {
		return 0
	}
	listingNumber, err := strconv.Atoi(listingString)
	if err != nil {
		return 0
	}
	return listingNumber
}
