package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Results struct {
	Results []Result `json:"results"`
}

type Result struct {
	Formatted_address string `json:"formatted_address"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Location_type string `json:"location_type"`
	Location GeoLocation `json:"location,omitempty"`
}

type GeoLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func GeocodeLocationHandler(w http.ResponseWriter, r *http.Request) {
	geocodeAppToken := r.Header.Get("X-Geocode-App-Token")

	format := mux.Vars(r)["format"]
	resourceApi := fmt.Sprintf("%s/maps/api/geocode/%s", BaseGeoApiUrl, format)
	req, err := http.NewRequest(http.MethodGet, resourceApi, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Pass along all query parameters
	req.URL.RawQuery = r.URL.Query().Encode()
	req.URL.RawQuery += fmt.Sprintf("&key=%s", geocodeAppToken)

	req.Header.Set("Content-Type", "application/json");

	client := NewClient(BaseGeoApiUrl)
	geoLocation, err := getGeocodeLocations(client, req)
	if err != nil {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", " ")
		err = encoder.Encode(err)
		if err != nil {
			panic(err)
		}

		panic(err)
	} else {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", " ")
		err = encoder.Encode(geoLocation)
		if err != nil {
			panic(err)
		}
	}
}

func getGeocodeLocations(client *Client, r *http.Request) (GeoLocation, error) {
	resp, err := client.httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	results := Results{}
	err = json.Unmarshal(body, &results)

	if err != nil {
		return GeoLocation{}, err
	} else {
		if (results.Results != nil && len(results.Results) > 0) {
			return results.Results[0].Geometry.Location, nil
		}

		return GeoLocation{}, nil
	}
}
