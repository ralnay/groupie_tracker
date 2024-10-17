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

var (
	a   []Artist
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

func RelationData() []Relation {
	var bytes []byte
	relation, err2 := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err2 != nil {
		log.Fatal()
	}
	relationData, err2 := ioutil.ReadAll(relation.Body)
	if err2 != nil {
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
	data := make([]Data, len(a))
	for i := 0; i < len(a); i++ {
		data[i].A = a[i]
		data[i].R = r[i]
	}
	return data
}
