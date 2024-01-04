package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(artistID)
	if err != nil && (id > 1 || id < 52) {
		// Gestion de l'erreur (par exemple, afficher une page d'erreur)
		ErrorRender(w, http.StatusNotFound)
		return
	}

	// Récupérez les informations de l'artiste par ID (vous devez implémenter cette fonction)
	artist, err := GetArtistByID(id, w, r)
	if err != nil {
		ErrorRender(w, http.StatusNotFound)
		return
	}

	Location, err := GetlocateByID(id, w, r)
	if err != nil {
		ErrorRender(w, http.StatusNotFound)
		return
	}

	ConcertDates, err := GetConcertDatesByID(id, w, r)
	if err != nil {
		ErrorRender(w, http.StatusInternalServerError)
		return
	}

	RelationsData, err := GetRelationsByID(id, w, r)
	if err != nil {
		ErrorRender(w, http.StatusInternalServerError)
		return
	}

	// `finalData` is a variable of type `final` that is used to store the data of an artist. It contains
	// the following fields:
	finalData := OneInfos{
		Artists:      artist,
		Location:     Location,
		ConcertDates: ConcertDates,
		Relations:    RelationsData,
	}
	// Chargez le modèle pour afficher les détails de l'artiste
	renderTemplate(w, "artist_details", finalData)
}

func Home(w http.ResponseWriter, r *http.Request) {
	client = &http.Client{}
	if r.Method != http.MethodGet {
		ErrorRender(w, http.StatusMethodNotAllowed)
	}
	if r.URL.Path != "/" {
		// `ErrorRender(w, http.StatusNotFound)` is a function call that is used to render an error page with
		// the status code 404 (Not Found). It is typically called when a requested resource is not found.
		ErrorRender(w, http.StatusNotFound)
	} else {
		/*
			ici dans la page home on a besoin de de fetch l'API pour avoir la les Artists Et la Localisation
			comme vous pouvez le voir dans la sturcture AllInfos. Ainsi on peut avoir les donnees de tous les
			Artists et de leur Localisation charger dans la page home.
		*/
		urlart := "https://groupietrackers.herokuapp.com/api/artists"
		urlLoc := "https://groupietrackers.herokuapp.com/api/locations"

		Artists, err := GetArtistInfo(urlart)
		if err != nil {
			ErrorRender(w, http.StatusInternalServerError)
			return
		}
		location, err := GetLocationsInfo(urlLoc)
		if err != nil {
			ErrorRender(w, http.StatusInternalServerError)
			return
		}

		finalData2 := AllInfos{
			AllArtists:  Artists,
			AllLocation: location,
		}
		renderTemplate(w, "home", finalData2)
	}
}

func renderTemplate(w http.ResponseWriter, page string, dataArt interface{}) {
	t, err1 := template.ParseFiles("./Templates/" + page + ".html")
	if err1 != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	t, err := t.ParseFiles("./Static/style.css")
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, dataArt)

}
