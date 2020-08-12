package controllers

import (
	"fmt"
	"github.com/dbenavraham1/munchspot/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	mobileFoodFacilityPermitsJsonResponse = `[
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
	mobileFoodFacilityPermitsXmlResponse = `
<?xml version="1.0" encoding="UTF-8"?>
<response>
   <row>
      <row _id="row-vsry_nu8i_9kqi" _uuid="00000000-0000-0000-44CF-DFC7E145D135" _position="0" _address="https://data.sfgov.org/resource/rqzj-sfat/row-vsry_nu8i_9kqi">
         <objectid>934609</objectid>
         <applicant>May Catering</applicant>
         <facilitytype>Truck</facilitytype>
         <cnn>3887000</cnn>
         <locationdescription>CHANNEL ST: 03RD ST to 04TH ST (0 - 0)</locationdescription>
         <address>Assessors Block 8711/Lot023</address>
         <blocklot>8711023</blocklot>
         <block>8711</block>
         <lot>023</lot>
         <permit>17MFF-0110</permit>
         <status>EXPIRED</status>
         <fooditems>Cold Truck: Sandwiches: fruit: snacks: candy: hot and cold drinks</fooditems>
         <x>6015253.3503</x>
         <y>2109839.76272</y>
         <latitude>37.7740748410583</latitude>
         <longitude>-122.390668458146</longitude>
         <schedule>http://bsm.sfdpw.org/PermitsTracker/reports/report.aspx?title=schedule&amp;report=rptSchedule&amp;params=permit=17MFF-0110&amp;ExportPDF=1&amp;Filename=17MFF-0110_schedule.pdf</schedule>
         <dayshours>Mo/Mo/Mo/Mo/Mo:9AM-10AM/10AM-11AM/12PM-1PM/1PM-2PM</dayshours>
         <approved>2017-12-28T00:00:00</approved>
         <received>2017-02-13</received>
         <priorpermit>1</priorpermit>
         <expirationdate>2018-07-15T00:00:00</expirationdate>
         <location>POINT (-122.390668458146 37.7740748410583)</location>
      </row>
   </row>
</response>`
)

func TestGetFoodLocationsJson(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testAppAuthToken", r.Header.Get("X-App-Token"))
		w.Write([]byte(mobileFoodFacilityPermitsJsonResponse))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food/resource/%s/%s", test.BaseTestApiUrl, "testId", "json"), nil)
	req.Header.Set("X-App-Token", "testAppAuthToken")

	mobileFoodFacilityPermits, err := getFoodLocationsJson(client, req)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(mobileFoodFacilityPermits))

	assert.Equal(t, "1337920", mobileFoodFacilityPermits[0].Objectid)
	assert.Equal(t, "37.7691244121681", mobileFoodFacilityPermits[0].Latitude)
	assert.Equal(t, "-122.391474911246", mobileFoodFacilityPermits[0].Longitude)
}

func TestGetFoodLocationsJsonEmptyMobileFoodFacilityPermits(t *testing.T) {
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

	mobileFoodFacilityPermits, err := getFoodLocationsJson(client, req)

	assert.Nil(t, mobileFoodFacilityPermits)
	assert.NotNil(t, err)
	assert.Equal(t, "unexpected end of JSON input", err.Error())
}

func TestGetFoodLocationsXml(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testAppAuthToken", r.Header.Get("X-App-Token"))
		w.Write([]byte(mobileFoodFacilityPermitsXmlResponse))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food/resource/%s/%s", test.BaseTestApiUrl, "testId", "xml"), nil)
	req.Header.Set("X-App-Token", "testAppAuthToken")

	mobileFoodFacilityPermits, err := getFoodLocationsXml(client, req)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(mobileFoodFacilityPermits.Row.Row))

	assert.Equal(t, "934609", mobileFoodFacilityPermits.Row.Row[0].Objectid)
	assert.Equal(t, "37.7740748410583", mobileFoodFacilityPermits.Row.Row[0].Latitude)
	assert.Equal(t, "-122.390668458146", mobileFoodFacilityPermits.Row.Row[0].Longitude)
}

func TestGetFoodLocationsXmlEmptyMobileFoodFacilityPermits(t *testing.T) {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "testAppAuthToken", r.Header.Get("X-App-Token"))
		w.Write([]byte(""))
	})
	httpClient, teardown := test.TestingHTTPClient(h)
	defer teardown()

	client := NewClient(test.BaseTestApiUrl)
	client.httpClient = httpClient

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/food/resource/%s/%s", test.BaseTestApiUrl, "testId", "xml"), nil)
	req.Header.Set("X-App-Token", "testAppAuthToken")

	mobileFoodFacilityPermits, err := getFoodLocationsXml(client, req)

	assert.NotNil(t, mobileFoodFacilityPermits)
	assert.Equal(t, 0, len(mobileFoodFacilityPermits.Row.Row))
	assert.NotNil(t, err)
	assert.Equal(t, "EOF", err.Error())
}
