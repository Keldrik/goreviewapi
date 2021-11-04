package lib

import (
	"net/http"
	"strconv"
)

func GetIntQuery(r *http.Request, key string) int {
	intString := r.URL.Query().Get(key)
	if intString == "" {
		return 0
	}
	intQuery, err := strconv.Atoi(intString)
	if err != nil {
		return 0
	}
	return intQuery
}

func GetLengthQuery(r *http.Request) int {
	lengthNumber := GetIntQuery(r, "length")
	if lengthNumber == 0 {
		return 6
	}
	return lengthNumber
}

func GetPageQuery(r *http.Request) int {
	pageNumber := GetIntQuery(r, "page")
	if pageNumber == 0 {
		return 1
	}
	return pageNumber
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
