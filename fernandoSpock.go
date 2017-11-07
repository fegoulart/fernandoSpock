package fernandoSpock

import (
"encoding/json"
"fmt"
"os"
"net/http"
"net/url"
"log"
"bytes"
)

type CharacterStruct struct {
	UID                    string  `json:"uid"`
	Name                   string  `json:"name"`
	Gender                 string  `json:"gender"`
	YearOfBirth            int64   `json:"yearOfBirth"`
	MonthOfBirth           int64   `json:"monthOfBirth"`
	PlaceOfBirth           string  `json:"placeOfBirth"`
	YearOfDeath            int64   `json:"yearOfDeath"`
	MonthOfDeath           int64   `json:"monthOfDeath"`
	DayOfDeath             int64   `json:"dayOfDeath"`
	PlaceOfDeath           string  `json:"placeOfDeath"`
	Height                 float32 `json:"height"`
	Weight                 float32 `json:"weight"`
	Deceased               bool    `json:"deceased"`
	BloodType              string  `json:"bloodType"`
	MaritalStatus          string  `json:"string"`
	SerialNumber           string  `json:"serialNumber"`
	HologramActivationDate int32   `json:"hologramActivationDate"`
	HologramStatus         string  `json:"hologramStatus"`
	HologramDateStatus     string  `json:"hologramDateStatus"`
	Hologram               bool    `json:"hologram"`
	FictionalCharacter     bool    `json:"fictionalCharacter"`
	Mirror                 bool    `json:"mirror"`
	AlternateReality       bool    `json:"alternateReality"`

}

type PageJson struct {
	PageNumber      int64 `json:"pageNumber"`
	PageSize        int64 `json:"pageSize"`
	NumberOfElements int64 `json:"numberOfElements"`
	TotalElements   int64 `json:"totalElements"`
	TotalPages      int64 `json:"totalPages"`
	FirstPage       bool  `json:"firstPage"`
	LastPage        bool  `json:"lastPage"`
}

type SortStruct struct {
	Clauses []string `json:"clauses"`
}

type PageStruct struct {
	Page       PageJson          `json:"page"`
	Sort       SortStruct        `json:"sort"`
	Characters []CharacterStruct `json:"characters"`
}


var characterSearchUrl = "http://stapi.co/api/v1/rest/character/search"
//TODO: Uncomment characterUrl
//var characterUrl  = "http://stapi.co/api/v1/rest/character?uid="
var specieUrl = "http://stapi.co/api/v1/rest/species?uid="


type LocationStruct struct {
	UID string `json:"uid"`
	Name string `json:"name"`
}

type HomeWorldStruct struct {
	UID string `json:"uid"`
	Name string `json:"name"`
	AstronomicalObjectType string `json:"astronomicalObjectType"`
	Location LocationStruct `json:"location"`
}

type QuadrantStruct struct {
	UID string `json:"uid"`
	Name string `json:"name"`
	AstronomicalObjectType string `json:"astronomicalObjectType"`
	Location LocationStruct `json:"location"`

}

type SpeciesStruct struct {
	UID string `json:"uid"`
	Name string `json:"name"`
	HomeWorld HomeWorldStruct `json:"homeWorldStruct"`
	Quadrant QuadrantStruct `json:"quadrant"`
	ExtinctSpecies bool `json:"extinctSpecies"`
	WarpCapableSpecies bool `json:"warpCapableSpecies"`
	ExtraGalacticSpecies bool `json:"extraGalacticSpecies"`
	HumanoidSpecies bool `json:"humanoidSpecies"`
	ReptilianSpecies bool `json:"reptilianSpecies"`
	NonCorporealSpecies bool `json:"nonCorporealSpecies"`
	ShapeshiftingSpecies bool `json:"shapeshiftingSpecies"`
	SpaceborneSpecies bool `json:"spaceborneSpecies"`
	TelepathicSpecies bool `json:"telepathicSpecies"`
	TransDimensionalSpecies bool `json:"transDimensionalSpecies"`
	UnnamedSpecies bool `json:"unnamedSpecies"`
	AlternateReality bool `json:"alternateReality"`
	Characters []CharacterStruct `json:"characters"`
}

//TODO: Create Character Structs

func main() {


	//Command line arguments

	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	//TODO: Translate to Klingon


	//Call to Character Endpoint

	//FIXME: Change to arguments
	data := url.Values{}
	data.Set("title", "Uhura")
	data.Add("name", "Uhura")

	req, err := http.NewRequest("POST", characterSearchUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var inputJson PageStruct


	if err := json.NewDecoder(resp.Body).Decode(&inputJson); err != nil {
		log.Println(err)
	}

	var elements int64
	elements = inputJson.Page.NumberOfElements

	if elements > 0 {
		fmt.Println(inputJson.Characters[0].UID)

		// Make another call to get Character Details

		//TODO: Make call to get character details


		// Make another call to get Species
		//FIXME: Change url
		req3, err := http.NewRequest("GET", specieUrl + inputJson.Characters[0].UID, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return
		}
		client3 := &http.Client{}

		resp3, err := client3.Do(req3)
		if err != nil {
			log.Fatal("Do: ", err)
			return
		}

		defer resp3.Body.Close()

		var speciesJson SpeciesStruct


		if err := json.NewDecoder(resp3.Body).Decode(&speciesJson); err != nil {
			log.Println(err)
		}

		//TODO: Erase this
		fmt.Println(resp.Status)

		fmt.Println(speciesJson)
		fmt.Println(speciesJson.Name)



	} else {
		fmt.Println("Character not found")
	}

}
