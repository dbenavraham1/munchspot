### MunchSpot

#### Purpose
This project provides API services for finding food truck locations around San Francisco.
It serves any frontend platform (iOS, Android, webapp) looking to integrate food truck locations
into their app.

#### Design
The overall design provides proxy endpoints to targeted Open Data
([Mobile Food Facility Permit](https://dev.socrata.com/foundry/data.sfgov.org/rqzj-sfat)) and
Google [Geocoding API](https://developers.google.com/maps/documentation/geocoding/overview)
endpoint API services that retrieve a user's location and finds the nearest food trucks.
- First, a user's location is determined from the device they're on or entered as an address.
- Second, the user's location is used to find nearby available food trucks.

I was able to leverage the existing Open Data/Google geocoding API services above to locate and find
all nearby food trucks around San Francisco.

**Note:**
This is my second golang project and I'm looking to gain more experience with it.

#### Dependencies
- Generate an Open Data app token [here](https://data.sfgov.org/profile/edit/developer_settings)
or use an existing one.
- Generate a Google Geocoding API key [here](https://developers.google.com/maps/documentation/geocoding/get-api-key)
or use an existing one.
- Install Golang [here](https://golang.org/doc/install)
- Install Docker [here](https://docs.docker.com/v17.09/engine/installation/)

#### Endpoints

- `/geolocation/{format:"json/xml"}?address=street%20zipcode&key=apikey` **GET**

This endpoint provides the ability to lookup any free formed address to retrieve the lat/lng
coordinates to help point a marker on a map. The coordinates are used as input to the 
food truck location lookup service.

**Required:** X-Geocode-App-Token header
              
`X-Geocode-App-Toke: {apikey}`

##### Example request
```bash
curl -H 'X-Geocode-App-Token: {apikey}' 'http://localhost:8080/geolocation/json?address=16209%20W%2070th%20Pl%20Arvada%20CO'
```

##### Query parameters

| Parameter    | Default   | Required  |
| ------------- |:-------------:| -----:|
| address      | N/A | yes |

**Example Return Response**
```json
{
 "lat": 39.8251727,
 "lng": -105.1819809
}
```

__________________________________________________________________

- `/food/resource/{id}/{format:"json/xml/csv"}` **GET**

This endpoint is used to lookup food trucks around San Francisco, based on provided lat/lng coordinates.

**Required:** Authorization header

`X-App-Token: {appToken}`

##### Example request
```bash
curl -H 'X-App-Token: {appToken}' 'http://localhost:8080/food/resource/rqzj-sfat/json?$where=within_circle(location,37.7708922310318,-122.389169231483,500)'
```

```bash
curl -H 'X-App-Token: {appToken}' 'http://localhost:8080/food/resource/rqzj-sfat/json?$where=within_circle(location,37.7708922310318,-122.389169231483,500)&$limit=5&$offset=0'
```

| Parameter    | Default   | Required  |
| ------------- |:-------------:| -----:|
| $where      | N/A | no |
| $limit      | N/A | no |
| $offset     | N/A | no |

**Example Return Response**
```json
[{
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
 }]
```

#### Running
- Clone
```
git clone https://github.com/dbenavraham1/munchspot.git
cd munchspot
```
- Build
```
docker build -t munchspot .
```
- Run
```
docker run -p 8080:8080 -d --name munchspot --rm munchspot
```
- Stop
```
docker stop munchspot
```
- Remove
```
docker rmi munchspot
```

#### Testing
```bash
go test -v *.go
```
