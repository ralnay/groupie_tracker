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
	// L Location
	// D Date
}

type Artist struct {
	Id         uint     `json:"id"`
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Members    []string `json:"members"`
	Creation   uint     `json:"creationDate"`
	FirstAlbum string   `json:"firstAlbum"`
}

// type Location struct {
// 	Locations []string `json:"locations"`
// }

// type Date struct {
// 	Dates []string `json:"dates"`
// }

type Relation struct {
	Relation map[string][]string `json:"datesLocations"`
}

var (
	a   []Artist
	// lMap  map[string]json.RawMessage
	// l []Location
	// dMap     map[string]json.RawMessage
	// d    []Date
	rMap  map[string]json.RawMessage
	r []Relation
)

func ArtistData() []Artist {
	art, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal()
	}
	artistData, err := ioutil.ReadAll(art.Body)
	if err != nil {
		log.Fatal()
	}
	json.Unmarshal(artistData, &a)
	return a
}

// func Loc() []Location {
// 	var bytes []byte
// 	location, err2 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
// 	if err2 != nil {
// 		log.Fatal()
// 	}
// 	locationData, err3 := ioutil.ReadAll(location.Body)
// 	if err3 != nil {
// 		log.Fatal()
// 	}
// 	err := json.Unmarshal(locationData, &lMap)
// 	if err != nil {
// 		fmt.Println("error :", err)
// 	}
// 	for _, m := range lMap {
// 		for _, v := range m {
// 			bytes = append(bytes, v)
// 		}
// 	}
// 	err = json.Unmarshal(bytes, &l)
// 	if err != nil {
// 		fmt.Println("error :", err)
// 	}
// 	return l
// }

// func Dates() []Date {
// 	var bytes []byte
// 	dates, err2 := http.Get("https://groupietrackers.herokuapp.com/api/dates")
// 	if err2 != nil {
// 		log.Fatal()
// 	}
// 	datesData, err3 := ioutil.ReadAll(dates.Body)
// 	if err3 != nil {
// 		log.Fatal()
// 	}
// 	err := json.Unmarshal(datesData, &dMap)
// 	if err != nil {
// 		fmt.Println("error :", err)
// 	}
// 	for _, m := range dMap {
// 		for _, v := range m {
// 			bytes = append(bytes, v)
// 		}
// 	}
// 	err = json.Unmarshal(bytes, &d)
// 	if err != nil {
// 		fmt.Println("error :", err)
// 	}
// 	return d
// }

func RelationData() []Relation {
	var bytes []byte
	relation, err2 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err2 != nil {
		log.Fatal()
	}
	relationData, err3 := ioutil.ReadAll(relation.Body)
	if err3 != nil {
		log.Fatal()
	}
	err := json.Unmarshal(relationData, &rMap)
	if err != nil {
		fmt.Println("error :", err)
	}

	for _, m := range rMap {
		for _, v := range m {
			bytes = append(bytes, v)
		}
	}

	err = json.Unmarshal(bytes, &r)
	if err != nil {
		fmt.Println("error :", err)
	}
	return r
}

func AllData() []Data {
	ArtistData()
	RelationData()
	// Loc()
	// Dates()
	data := make([]Data, len(a))
	for i := 0; i < len(a); i++ {
		data[i].A = a[i]
		data[i].R = r[i]
		// data[i].L = l[i]
		// data[i].D = d[i]
	}
	return data
}
