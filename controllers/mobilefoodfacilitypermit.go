package controllers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	jsonFormat = "json"
	xmlFormat = "xml"
	csvFormat = "csv"
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

type MobileFoodFacilityPermitsResponse struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Row     struct {
		Text string `xml:",chardata"`
		Row  []struct {
			Text                string `xml:",chardata"`
			ID                  string `xml:"_id,attr"`
			Uuid                string `xml:"_uuid,attr"`
			Position            string `xml:"_position,attr"`
			AttrAddress         string `xml:"_address,attr"`
			Objectid            string `xml:"objectid"`
			Applicant           string `xml:"applicant"`
			Facilitytype        string `xml:"facilitytype"`
			Cnn                 string `xml:"cnn"`
			Locationdescription string `xml:"locationdescription"`
			Address             string `xml:"address"`
			Blocklot            string `xml:"blocklot"`
			Block               string `xml:"block"`
			Lot                 string `xml:"lot"`
			Permit              string `xml:"permit"`
			Status              string `xml:"status"`
			Fooditems           string `xml:"fooditems"`
			X                   string `xml:"x"`
			Y                   string `xml:"y"`
			Latitude            string `xml:"latitude"`
			Longitude           string `xml:"longitude"`
			Schedule            string `xml:"schedule"`
			Dayshours           string `xml:"dayshours"`
			Approved            string `xml:"approved"`
			Received            string `xml:"received"`
			Priorpermit         string `xml:"priorpermit"`
			Expirationdate      string `xml:"expirationdate"`
			Location            string `xml:"location"`
		} `xml:"row"`
	} `xml:"row"`
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

	if format == jsonFormat {
		mobileFoodFacilityPermitsJson, err := getFoodLocationsJson(client, req)
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
			err = encoder.Encode(mobileFoodFacilityPermitsJson)
			if err != nil {
				panic(err)
			}
		}
	} else if format == xmlFormat {
		mobileFoodFacilityPermitsCsv, err := getFoodLocationsCsv(client, req)
		if err != nil {
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", " ")
			err = encoder.Encode(err)
			if err != nil {
				panic(err)
			}

			panic(err)
		} else {
			w.Header().Set("Content-Disposition", "attachment; filename=mobilefoodfacilitypermits.csv")
			w.Header().Set("Content-Type", "text/csv; charset=UTF-8")
			w.Header().Set("Content-Length", strconv.Itoa(len(mobileFoodFacilityPermitsCsv)))
			if _, err := w.Write(mobileFoodFacilityPermitsCsv); err != nil {
				panic(err)
			}
		}
	} else if format == csvFormat {
		mobileFoodFacilityPermitsCsv, err := getFoodLocationsCsv(client, req)
		if err != nil {
			encoder := json.NewEncoder(w)
			encoder.SetIndent("", " ")
			err = encoder.Encode(err)
			if err != nil {
				panic(err)
			}

			panic(err)
		} else {
			w.Header().Set("Content-Disposition", "attachment; filename=mobilefoodfacilitypermits.csv")
			w.Header().Set("Content-Type", "text/csv; charset=UTF-8")
			w.Header().Set("Content-Length", strconv.Itoa(len(mobileFoodFacilityPermitsCsv)))
			if _, err := w.Write(mobileFoodFacilityPermitsCsv); err != nil {
				panic(err)
			}
		}
	}
}

func getFoodLocationsJson(client *Client, r *http.Request) ([]MobileFoodFacilityPermit, error) {
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

func getFoodLocationsXml(client *Client, r *http.Request) (MobileFoodFacilityPermitsResponse, error) {
	resp, err := client.httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	mobileFoodFacilityPermitsResponse := MobileFoodFacilityPermitsResponse{}
	err = xml.Unmarshal(body, &mobileFoodFacilityPermitsResponse)
	if err != nil {
		return mobileFoodFacilityPermitsResponse, err
	} else {
		return mobileFoodFacilityPermitsResponse, nil
	}
}

func getFoodLocationsCsv(client *Client, r *http.Request) ([]byte, error) {
	resp, err := client.httpClient.Do(r)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body, err
}
