package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func TabIdFilter(MinCreationDateN, MaxCreationDateN, MinDateN, MaxDateN int, NumOfMember []string, location string, AllData AllInfos, w http.ResponseWriter, r *http.Request) []int {
    // Déclaration de tranches pour stocker les ID correspondant à chaque critère
    var TabCreationDate []int
    var TabFirstAlbum []int
    var TabMembers []int
	var TabLocation []int

    // Parcours de la liste des artistes
    for _, a := range AllData.AllArtists {
        // Filtrage par date de création
        if a.CreationDate >= MinCreationDateN && a.CreationDate <= MaxCreationDateN {
            TabCreationDate = appendIfNotExists(TabCreationDate, a.Id)
        }

        // Filtrage par année du premier album
        TabfirstAlumSplit := strings.Split(string(a.FirstAlbum), "-")
        firstAlbumYear, _ := strconv.Atoi(TabfirstAlumSplit[len(TabfirstAlumSplit)-1])
        if firstAlbumYear >= MinDateN && firstAlbumYear <= MaxDateN {
            TabFirstAlbum = appendIfNotExists(TabFirstAlbum, a.Id)
        }

        // Filtrage par nombre de membres
        var MaxNumToSearch int
		var MinNumToSearch int
        MaxNumToSearch, _ = strconv.Atoi(NumOfMember[0])
        if len(NumOfMember) > 1 {
			MinNumToSearch, _ = strconv.Atoi(NumOfMember[0])
            MaxNumToSearch, _ = strconv.Atoi(NumOfMember[len(NumOfMember)-1])
            if len(a.Members) >= MinNumToSearch && len(a.Members) <= MaxNumToSearch {
                TabMembers = appendIfNotExists(TabMembers, a.Id)
            }
        } else {
            if len(a.Members) == MaxNumToSearch {
                TabMembers = appendIfNotExists(TabMembers, a.Id)
            }
        }

		for _, v := range AllData.AllLocation.Index {
			for _, locationcity := range v.Locations {
				if strings.ToLower(location) == "all" {
					TabLocation = appendIfNotExists(TabLocation, v.ID)
				}else if strings.ToLower(location) == strings.ToLower(locationcity) {
					TabLocation = appendIfNotExists(TabLocation, v.ID)
				}
				
			}
		}
		

    }

	//fmt.Println("TabLocation -> ",TabLocation)
	//fmt.Println("TabMembers -> ",TabMembers)
    // Trouver l'FoundMatchesId des trois tranches
    TabId := FoundMatchesId(TabCreationDate, TabFirstAlbum, TabMembers, TabLocation)
	fmt.Println("TabId -> ",TabId)

    return TabId
}


// Fonction pour etourner une nouvelle tranche qui contient uniquement les éléments communs à toutes les tranches
func FoundMatchesId(Tab ...[]int) []int {
	// Si aucune tranche n'est fournie, retourne une tranche vide
    if len(Tab) == 0 {
        return []int{}
    }
	// Initialise le résultat avec la première tranche

    result := Tab[0]
	// Parcours les tranches à partir de la deuxième

    for _, slice := range Tab[1:] {
		// Initialise une tranche temporaire pour stocker les id communs
        MatchesId := []int{}

		// Parcours les éléments de la tranche actuelle
        for _, element := range slice {

			 // Vérifie si l'élément existe déjà dans le résultat
            if VerifyExistingElement(result, element) {

				// Si oui, l'ajoute à la tranche temporaire
                MatchesId = append(MatchesId, element)
            }
        }
		 // Le résultat devient la tranche temporaire
        result = MatchesId
    }

    return result
}

// Fonction pour vérifier si un élément existe dans une tranche
func VerifyExistingElement(Tab []int, Id int) bool {
    for _, v := range Tab {
        if v == Id {
            return true
        }
    }
    return false
}

// Fonction pour ajouter un élément à une tranche uniquement s'il n'existe pas déjà
func appendIfNotExists(Tab []int, Id int) []int {
    for _, v := range Tab {
        if v == Id {
            return Tab
        }
    }
    return append(Tab, Id)
}

func convertFormValueToInt(value string) (int, error) {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}
	return num, nil
 }

// //cette fonction verifie la presence ou l'absence de doublon dans un tableau

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Println("Error parsing form:", err)
		ErrorRender(w, http.StatusBadRequest)
		return
	}


	//date de creation
	MinCreationDate := r.FormValue("mincreationdate")
	MinCreationDateN, err := convertFormValueToInt(MinCreationDate)
	if err != nil {
		ErrorRender(w, http.StatusBadRequest)
		return
	}
	MaxCreationDate := r.FormValue("maxcreationdate")
	MaxCreationDateN, err := convertFormValueToInt(MaxCreationDate)
	if err != nil {
		ErrorRender(w, http.StatusBadRequest)
		return
	}
	fmt.Println("creation date -> ",MinCreationDate, MaxCreationDateN)

	
	// premier album
	MinDate := r.FormValue("minfirstalbum")
	MinDateN, err := convertFormValueToInt(MinDate)
	if err != nil {
		ErrorRender(w, http.StatusBadRequest)
		return
	}
	
	MaxDate := r.FormValue("maxfirstalbum")
	MaxDateN, err := convertFormValueToInt(MaxDate)
	if err != nil {
		ErrorRender(w, http.StatusBadRequest)
		return
	}
	fmt.Println("First Album -> ",MinDateN, MaxDateN)

	// nombre de membre
	 var members []string
	 members = r.Form["members"]
	 fmt.Println(members)
	
	
	 if len(members) == 0 {
	 	members = []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	 }
	 //location
	location := r.FormValue("location")

	urlart := "https://groupietrackers.herokuapp.com/api/artists"
	urlLoc := "https://groupietrackers.herokuapp.com/api/locations"

	Artists, err := GetArtistInfo(urlart)
	if err != nil {
		ErrorRender(w, http.StatusInternalServerError)
		return
	}

	Location, err := GetLocationsInfo(urlLoc)
	if err != nil {
		ErrorRender(w, http.StatusInternalServerError)
		return
	}

	finalData2 := AllInfos{
		AllArtists:  Artists,
		AllLocation: Location,
	}
	var Final []int
	//Final = TabCreationDateN(MinCreationDateN, MaxCreationDateN, w, finalData2, r)
	Final = TabIdFilter(MinCreationDateN,MaxCreationDateN,MinDateN,MaxDateN,members,location,finalData2,w,r)
   var Charge AllInfos
   var wg sync.WaitGroup
   wg.Add(len(Final)) // Add the number of goroutines to the WaitGroup
   for _, artistId := range Final {
	   go func(id int) {
		   defer wg.Done()
		   artist, err := GetArtistByID(id, w, r)
		   if err != nil {
			   return
		   }
		   Charge.AllArtists = append(Charge.AllArtists, artist) // Append the artist to the slice
	   }(artistId)
   }
   wg.Wait() // Wait for all goroutines to finish
	message := Error{
		Erreur: "Search Not Match",
	}
	if len(Final) == 0 {
		renderTemplate(w, "error", message)
		return
	}
	renderTemplate(w, "filter", Charge)
}