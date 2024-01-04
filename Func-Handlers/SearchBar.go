package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
)

//SEARCH BAR DEBUT
// cette fonction parcours les structures afin de verifier si la valeur saisie correspond a
//un artist/band, member, FirstAlbum, CreationDate, locations et l'ajouter dans un tableau d'ID
func TabIdSearch(w http.ResponseWriter, num int, wordToSearch, Groupe string, AllData AllInfos, r *http.Request) []int {
	var TabId []int
	if num == 2 || num == 1 {
		for _, a := range AllData.AllArtists {
			if Groupe == "artist/band" || num == 1 {
				if matchString(a.Name, wordToSearch) {
					if !Inter(TabId, a.Id) {
						TabId = append(TabId, a.Id)
					}
				}
			}

			if Groupe == "member" || num == 1 {
				for _, j := range a.Members {
					if matchString(j, wordToSearch) {
						if !Inter(TabId, a.Id) {
							TabId = append(TabId, a.Id)
						}
					}
				}
			}
			if Groupe == "FirstAlbum" || num == 1 {
				if matchString(a.FirstAlbum, wordToSearch) {
					if !Inter(TabId, a.Id) {
						TabId = append(TabId, a.Id)
					}
				}
			}
			if Groupe == "CreationDate" || num == 1 {
				if matchString(strconv.Itoa(a.CreationDate), wordToSearch) {
					if !Inter(TabId, a.Id) {
						TabId = append(TabId, a.Id)
					}
				}

			}
		}
		if Groupe == "locations" || num == 1 {
			for _, a := range AllData.AllLocation.Index {
				for _, i := range a.Locations {
					if matchString(i, wordToSearch) {
						if !Inter(TabId, a.ID) {
							TabId = append(TabId, a.ID)
						}
					}
				}
			}
		}
	}
	return TabId
}
//cette fonction verifie la presence ou l'absence de doublon dans un tableau

func Inter(tab []int, j int) bool {
	for i := range tab {
		if tab[i] == j {
			return true
		}
	}
	return false
}

func SplitField(s string) (int, string, string) {
	//spliter la valeur saisie en deux si un # et retouner la valeur ou les valeur selon la presence ou l'absence de # 
	SplitS := strings.Split(s, "#")
	if len(SplitS) == 1 {
		SplitS[0] = strings.TrimSpace(SplitS[0])
		return 1, SplitS[0], ""
	}
	if len(SplitS) == 2 {
		SplitS[0] = strings.TrimSpace(SplitS[0])
		SplitS[1] = strings.TrimSpace(SplitS[1])
		if SplitS[1] == "artist/band" || SplitS[1] == "member" || SplitS[1] == "FirstAlbum" || SplitS[1] == "CreationDate" || SplitS[1] == "locations" {
			return 2, SplitS[0], SplitS[1]
		}
	}
	return 0, "", ""
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorRender(w, http.StatusMethodNotAllowed)
		return
	}
	Field := r.FormValue("searchinput")
	if Field == "" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}

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
	num, WordToSearch, Groupe := SplitField(Field)
	Final := TabIdSearch(w, num, WordToSearch, Groupe, finalData2, r)
	var Charge AllInfos
	//utilisation du go routine pour la vitesse du fetch par id
	var wg sync.WaitGroup
	for j := 0; j < len(Final); j++ {
		wg.Add(1)
		go func(artistId int) {
			defer wg.Done()
			artist, err := GetArtistByID(artistId, w, r)
			if err != nil {
				return
			}
			Charge.AllArtists = append(Charge.AllArtists, artist)
		}(Final[j])
	}
	wg.Wait()
	message := Error{
		Erreur: "Not Match Search",
	}
	if len(Final) == 0 {
		renderTemplate(w, "error", message)
		return
	}
	renderTemplate(w, "search", Charge)
}


//verifiation de la saisie avec to lower pour eviter la casse
func matchString(FirstString, SecondString string) bool {
	return strings.Contains(strings.ToLower(FirstString), strings.ToLower(SecondString))
}
