package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var apikey = "YOUR_API_KEY"
var latitude = 42.36399
var longitude = -71.05493
var address string

type output struct {
	Items []struct {
		Title string `json:"title"`
	} `json:"items"`
}

type MapValues struct {
	API       string
	Latitude  float64
	Longitude float64
	Address   string
}

func main() {
	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey=" + apikey + "&at=" + fmt.Sprint(latitude) + "," + fmt.Sprint(longitude)

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var data output
	json.Unmarshal(body, &data)

	for _, add := range data.Items {
		address = add.Title
		fmt.Printf(address)
	}

	http.HandleFunc("/", MapPage)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func MapPage(w http.ResponseWriter, r *http.Request) {
	Values := MapValues{
		API:       apikey,
		Latitude:  latitude,
		Longitude: longitude,
		Address:   address,
	}
	t, err := template.ParseFiles("map.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, Values)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
