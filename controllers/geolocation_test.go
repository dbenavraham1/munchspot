package controllers

import (
	"fmt"
	"github.com/dbenavraham1/munchspot/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	geolocationResponse = `{
   "results" : [
      {
         "address_components" : [
            {
               "long_name" : "16209",
               "short_name" : "16209",
               "types" : [ "street_number" ]
            },
            {
               "long_name" : "West 70th Place",
               "short_name" : "W 70th Pl",
               "types" : [ "route" ]
            },
            {
               "long_name" : "Meadows At Westwoods",
               "short_name" : "Meadows At Westwoods",
               "types" : [ "neighborhood", "political" ]
            },
            {
               "long_name" : "Arvada",
               "short_name" : "Arvada",
               "types" : [ "locality", "political" ]
            },
            {
               "long_name" : "Jefferson County",
               "short_name" : "Jefferson County",
               "types" : [ "administrative_area_level_2", "political" ]
            },
            {
               "long_name" : "Colorado",
               "short_name" : "CO",
               "types" : [ "administrative_area_level_1", "political" ]
            },
            {
               "long_name" : "United States",
               "short_name" : "US",
               "types" : [ "country", "political" ]
            },
            {
               "long_name" : "80007",
               "short_name" : "80007",
               "types" : [ "postal_code" ]
            },
            {
               "long_name" : "6968",
               "short_name" : "6968",
               "types" : [ "postal_code_suffix" ]
            }
         ],
         "formatted_address" : "16209 W 70th Pl, Arvada, CO 80007, USA",
         "geometry" : {
            "bounds" : {
               "northeast" : {
                  "lat" : 39.82526410000001,
                  "lng" : -105.1818807
               },
               "southwest" : {
                  "lat" : 39.8250955,
                  "lng" : -105.1821228
               }
            },
            "location" : {
               "lat" : 39.8251727,
               "lng" : -105.1819809
            },
            "location_type" : "ROOFTOP",
            "viewport" : {
               "northeast" : {
                  "lat" : 39.8265287802915,
                  "lng" : -105.1806527697085
               },
               "southwest" : {
                  "lat" : 39.8238308197085,
                  "lng" : -105.1833507302915
               }
            }
         },
         "place_id" : "ChIJh30Ar6-Pa4cR0Jj6hpZw-d4",
         "types" : [ "premise" ]
      }
   ],
   "status" : "OK"
}`
)

func TestGetGeocodeLocations(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testGeocodeAppAuthToken", r.Header.Get("X-Geocode-App-Token"))
		w.Write([]byte(geolocationResponse))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestGeoApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/geolocation/%s", test.BaseTestGeoApiUrl, "json"), nil)
	req.Header.Set("X-Geocode-App-Token", "testGeocodeAppAuthToken")

	geoLocation, err := getGeocodeLocations(client, req)

	assert.Nil(t, err)
	assert.NotNil(t, geoLocation)

	assert.Equal(t, 39.8251727, geoLocation.Lat)
	assert.Equal(t, -105.1819809, geoLocation.Lng)
}

func TestGetGeocodeLocationsEmptyGeoLocation(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testGeocodeAppAuthToken", r.Header.Get("X-Geocode-App-Token"))
		w.Write([]byte(""))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestGeoApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/geolocation/%s", test.BaseTestGeoApiUrl, "json"), nil)
	req.Header.Set("X-Geocode-App-Token", "testGeocodeAppAuthToken")

	geoLocation, err := getGeocodeLocations(client, req)

	assert.NotNil(t, err)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
	assert.NotNil(t, geoLocation)
	assert.Equal(t, 0.0, geoLocation.Lng)
	assert.Equal(t, 0.0, geoLocation.Lat)
}