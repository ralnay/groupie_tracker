package piscine

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	A Artist
	R Relation
	L Location
	D Date
}

type Artist struct {
	Id         uint     `json:"id"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Members    []string `json:"members"`
	Creation   uint     `json:"creationDate"`
	FirstAlbum string   `json:"firstAlbum"`
}

type Relation struct {
	Relation map[string][]string `json:"datesLocations"`
}

type Location struct {
	Locations map[string][]string `json:"locations"`
}

type Date struct {
	Dates map[string][]string `json:"dates"`
}

var (
	a    []Artist
	rMap map[string]json.RawMessage
	r    []Relation
	loc  []Location
	dat  []Date
)

func ArtistData() []Artist {
	art, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	artistData, err := ioutil.ReadAll(art.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(artistData, &a)
	return a
}

func RelationData() []Relation {
	var bytes []byte
	relation, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatal(err)
	}
	relationData, err := ioutil.ReadAll(relation.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(relationData, &rMap)
	if err != nil {
		fmt.Println("error:", err)
	}

	for _, m := range rMap {
		bytes = append(bytes, m...) // Appending multiple bytes (fix)
	}

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		fmt.Println("error:", err)
	}
	return r
}

func LocationData() []Location {
	locData, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatal(err)
	}
	locationBody, err := ioutil.ReadAll(locData.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(locationBody, &loc)
	return loc
}

func DateData() []Date {
	dateData, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatal(err)
	}
	dateBody, err := ioutil.ReadAll(dateData.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(dateBody, &dat)
	return dat
}

func AllData() []Data {
	ArtistData()
	RelationData()
	LocationData()
	DateData()
	data := make([]Data, len(a))
	for i := 0; i < len(a); i++ {
		data[i].A = a[i]
		data[i].R = r[i]
		data[i].L = loc[i]
		data[i].D = dat[i]
	}
	return data
}
