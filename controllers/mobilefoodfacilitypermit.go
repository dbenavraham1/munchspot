package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type MobileFoodFacilityPermit struct {
	Objectid            string   `json:"objectid"`
	Applicant           string   `json:"applicant"`
	Facilitytype        string   `json:"facilitytype"`
	Cnn                 string   `json:"cnn"`
	Locationdescription string   `json:"locationdescription"`
	Address             string   `json:"address"`
	Blocklot            string   `json:"blocklot"`
	Block               string   `json:"block"`
	Lot                 string   `json:"lot"`
	Permit              string   `json:"permit"`
	Status              string   `json:"status"`
	Fooditems           string   `json:"fooditems"`
	X                   string   `json:"x"`
	Y                   string   `json:"y"`
	Latitude            string   `json:"latitude"`
	Longitude           string   `json:"longitude"`
	Schedule            string   `json:"schedule"`
	Received            string   `json:"received"`
	Priorpermit         string   `json:"priorpermit"`
	Expirationdate      string   `json:"expirationdate"`
	Location            Location `json:"location,omitempty"`
}

type Location struct {
	Type        string  `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func FoodResourceHandler(w http.ResponseWriter, r *http.Request) {
	appToken := r.Header.Get("X-App-Token")

	id := mux.Vars(r)["id"]
	format := mux.Vars(r)["format"]
	resourceApi := fmt.Sprintf("%s/resource/%s.%s", BaseApiUrl, id, format)
	req, err := http.NewRequest(http.MethodGet, resourceApi, r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Pass along all query parameters
	req.URL.RawQuery = r.URL.Query().Encode()

	req.Header.Set("X-App-Token", appToken);
	req.Header.Set("Content-Type", "application/json");

	client := NewClient(BaseApiUrl)
	mobileFoodFacilityPermits, err := getFoodLocations(client, req)
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
		err = encoder.Encode(mobileFoodFacilityPermits)
		if err != nil {
			panic(err)
		}
	}
}

func getFoodLocations(client *Client, r *http.Request) ([]MobileFoodFacilityPermit, error) {
	resp, err := client.httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	mobileFoodFacilityPermits := []MobileFoodFacilityPermit{}
	err = json.Unmarshal(body, &mobileFoodFacilityPermits)
	if err != nil {
		return nil, err
	} else {
		return mobileFoodFacilityPermits, nil
	}
}
