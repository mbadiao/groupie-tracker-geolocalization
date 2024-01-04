package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var client2 = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{},) error {
	resp, err := client2.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func ErrorRender(w http.ResponseWriter, codeError int) {
	w.WriteHeader(codeError)
	errorData := Error{
		Erreur: http.StatusText(codeError),
	}
	renderTemplate(w, "error", errorData)
	return
}

func GetArtistInfo(url string) ([]ArtistResult, error) {
	var artist []ArtistResult
	err := GetJson(url, &artist)
	return artist, err
}

func GetLocationsInfo(url string) (AllLocalisation, error) {
	var location AllLocalisation
	err := GetJson(url, &location)
	return location, err
}

func GetConcertDatesInfo(url string) ([]ConcertDatesData, error) {
	var concertDates []ConcertDatesData
	err := GetJson(url, &concertDates)
	return concertDates, err
}

func GetRelationsInfo(url string) ([]RelationsData, error) {
	var relationData []RelationsData
	err := GetJson(url, &relationData)
	return relationData, err
}

// Ajoutez une nouvelle fonction pour récupérer un artiste par ID

func GetArtistByID(id int, w http.ResponseWriter,r *http.Request) (ArtistResult, error) {
	MethodGet(w,r)
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", id)
	var artist ArtistResult
	err := GetJson(url, &artist)
	// Renvoie l'artiste trouvé
	return artist, err
}

func GetlocateByID(id int, w http.ResponseWriter,r *http.Request) (LocationsData, error) {
	MethodGet(w,r)
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)

	var locationData LocationsData
	err := GetJson(url, &locationData)
	// Renvoie l'artiste trouvé
	return locationData, err
}

func GetConcertDatesByID(id int, w http.ResponseWriter,r *http.Request) (ConcertDatesData, error) {
	MethodGet(w,r)
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	var concertDates ConcertDatesData
	err := GetJson(url, &concertDates)
	concertDates.Dates = Format(concertDates.Dates) 
	return concertDates, err
}

func GetRelationsByID(id int, w http.ResponseWriter,r *http.Request) (RelationsData, error) {
	MethodGet(w,r)
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	var relationData RelationsData
	err := GetJson(url, &relationData)
	return relationData, err
}

func MethodGet(w http.ResponseWriter,r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorRender(w, http.StatusMethodNotAllowed)
		return
	}
}
//cette fonction a pour but d'enlevee * dans une chaine de caractere
func Format(Date []string) []string {
	Tab := []string{}
	for _,val := range Date {
		Tab  = append(Tab,strings.Trim(val,"*") ) 
	}
	return Tab
}