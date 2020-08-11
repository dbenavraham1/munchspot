package controllers

import (
	"fmt"
	"github.com/dbenavraham1/munchspot/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	mobileFoodFacilityPermitsResponse = `[
 {
  "objectid": "1337920",
  "applicant": "La Jefa",
  "facilitytype": "",
  "cnn": "19987000",
  "locationdescription": "04TH ST: NELSON RISING LN to GENE FRIEND WAY (1500 - 1599)",
  "address": "1550 04TH ST",
  "blocklot": "8711007",
  "block": "8711",
  "lot": "007",
  "permit": "19MFF-00018",
  "status": "REQUESTED",
  "fooditems": "Tacos: burritos: quesadillas: tortas: nachos (refried beans: cheese sauce: salsa fresca): carnes (beef: chicken: marinated pork: fried pork): canned beans: rice: sodas: horchata drinks.",
  "x": "6014983.88328",
  "y": "2108042.28534",
  "latitude": "37.7691244121681",
  "longitude": "-122.391474911246",
  "schedule": "http://bsm.sfdpw.org/PermitsTracker/reports/report.aspx?title=schedule\u0026report=rptSchedule\u0026params=permit=19MFF-00018\u0026ExportPDF=1\u0026Filename=19MFF-00018_schedule.pdf",
  "received": "2019-05-22",
  "priorpermit": "0",
  "expirationdate": "",
  "location": {
   "type": "Point",
   "coordinates": [
    -122.391474911246,
    37.7691244121681
   ]
  }
 }
	]`
)

func TestGetFoodLocations(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testAppAuthToken", r.Header.Get("X-App-Token"))
		w.Write([]byte(mobileFoodFacilityPermitsResponse))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food/resource/%s/%s", test.BaseTestApiUrl, "testId", "json"), nil)
	req.Header.Set("X-App-Token", "testAppAuthToken")

	mobileFoodFacilityPermits, err := getFoodLocations(client, req)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(mobileFoodFacilityPermits))

	assert.Equal(t, "1337920", mobileFoodFacilityPermits[0].Objectid)
	assert.Equal(t, "37.7691244121681", mobileFoodFacilityPermits[0].Latitude)
	assert.Equal(t, "-122.391474911246", mobileFoodFacilityPermits[0].Longitude)
}

func TestGetFoodLocationsEmptyMobileFoodFacilityPermits(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testAppAuthToken", r.Header.Get("X-App-Token"))
		w.Write([]byte(""))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food/resource/%s/%s", test.BaseTestApiUrl, "testId", "json"), nil)
	req.Header.Set("X-App-Token", "testAppAuthToken")

	mobileFoodFacilityPermits, err := getFoodLocations(client, req)

	assert.Nil(t, mobileFoodFacilityPermits)
	assert.NotNil(t, err)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}
