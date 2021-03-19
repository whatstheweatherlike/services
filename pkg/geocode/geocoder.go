package geocode

import (
	"fmt"
	"github.com/codingsince1985/geo-golang"
	"github.com/codingsince1985/geo-golang/opencage"
	"os"
	"strconv"
	"strings"
)

var GeoCoder geo.Geocoder

func init() {
	GeoCoder = opencage.Geocoder(os.Getenv("OPENCAGEDATA_APIKEY"))
}

func Reverse(coords string) (string, error) {
	latLon := strings.Split(coords, ",")
	lat, err := strconv.ParseFloat(latLon[0], 64)
	if err != nil {
		return "", fmt.Errorf("error reverse geocoding coords %s: %v", coords, err)
	}
	lng, err := strconv.ParseFloat(latLon[1], 64)
	if err != nil {
		return "", fmt.Errorf("error reverse geocoding coords %s: %v", coords, err)
	}
	address, err := GeoCoder.ReverseGeocode(lat, lng)
	if err != nil {
		return "", fmt.Errorf("error reverse geocoding coords %s: %v", coords, err)
	}
	return address.FormattedAddress, nil
}
