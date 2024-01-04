package handlers

import "net/http"

var client *http.Client

/*
	La structure OneInfos permette d'avoir les Information, La Localisation ,La date ou Les dates de concert
	et les relations  d'un Artist ou d'un Groupe d'Artists
*/
type OneInfos struct {
	Artists      ArtistResult
	Location     LocationsData
	ConcertDates ConcertDatesData
	Relations    RelationsData
}

/*
	La structure AllInfos permette d'avoir les Informations et toutes les localisations
	de tous les Artists ou tous les Groupes d'Artists
*/
type AllInfos struct {
	AllArtists  []ArtistResult
	AllLocation AllLocalisation
}

/*
	La structure ArtistResult permette de recevoir les informations d'un Artist ou d'un Groupe d'Artists
*/
type ArtistResult struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      []string
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
}

/*
	La structure RelationsData permette d'avoir les relations  d'un Artist ou d'un Groupe d'Artists
*/
type RelationsData struct {
	Id             int
	DatesLocations map[string][]string
}

/*
	La structure LocationsData permette d'avoir  La Localisation   d'un Artist ou d'un Groupe d'Artists
*/
type LocationsData struct {
	Id        int
	Locations []string
}

/*
	La structure ConcertDatesData permette d'avoir La date ou Les dates de concert
	d'un Artist ou d'un Groupe d'Artists
*/
type ConcertDatesData struct {
	Id    int
	Dates []string
}

/*
	La structure AllLocalisationpermette d'avoir toutes les localisations
	de tous les Artists ou tous les Groupes d'Artists
*/
type AllLocalisation struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

/*
	La structure Error nous permette d'aficher le type d'erreur
	sans pour autant avoir a creer plusieur pages d'erreur
*/
type Error struct {
	Erreur string
}
