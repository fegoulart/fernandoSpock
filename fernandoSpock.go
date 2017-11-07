package main

import (
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"log"
	"bytes"
)

// URLs
var characterSearchUrl = "http://stapi.co/api/v1/rest/character/search"
var characterUrl = "http://stapi.co/api/v1/rest/character?uid="

// Structs

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
	PageNumber       int64 `json:"pageNumber"`
	PageSize         int64 `json:"pageSize"`
	NumberOfElements int64 `json:"numberOfElements"`
	TotalElements    int64 `json:"totalElements"`
	TotalPages       int64 `json:"totalPages"`
	FirstPage        bool  `json:"firstPage"`
	LastPage         bool  `json:"lastPage"`
}

type SortStruct struct {
	Clauses []string `json:"clauses"`
}

type PageStruct struct {
	Page       PageJson          `json:"page"`
	Sort       SortStruct        `json:"sort"`
	Characters []CharacterStruct `json:"characters"`
}

type PerformerStruct struct {
	UID                string `json:"uid"`
	Name               string `json:"name"`
	BirthName          string `json:"birthName"`
	Gender             string `json:"gender"`
	DateOfBirth        string `json:"dateOfBirth"`
	PlaceOfBirth       string `json:"placeOfBirth"`
	DateOfDeath        string `json:"dateOfDeath"`
	PlaceOfDeath       string `json:"placeOfDeath"`
	AnimalPerformer    bool   `json:"animalPerformer"`
	DisPerformer       bool   `json:"disPerformer"`
	Ds9Performer       bool   `json:"ds9Performer"`
	EntPerformer       bool   `json:"entPerformer"`
	FilmPerformer      bool   `json:"filmPerformer"`
	StandInPerformer   bool   `json:"standInPerformer"`
	StuntPerformer     bool   `json:"stuntPerformer"`
	TasPerformer       bool   `json:"tasPerformer"`
	TngPerformer       bool   `json:"tngPerformer"`
	TosPerformer       bool   `json:"tosPerformer"`
	VideoGamePerformer bool   `json:"videoGamePerformer"`
	VoicePerformer     bool   `json:"voicePerformer"`
	VoyPerformer       bool   `json:"voyPerformer"`
}

type SeasonStruct struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type SeriesStruct struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type EpisodeStruct struct {
	UID                    string       `json:"uid"`
	Title                  string       `json:"title"`
	TitleGerman            string       `json:"titleGerman"`
	TitleItalian           string       `json:"titleItalian"`
	TitleJapanese          string       `json:"titleJapanese"`
	Series                 SeriesStruct `json:"series"`
	Season                 SeasonStruct `json:"season"`
	SeasonNumber           int64        `json:"seasonNumber"`
	EpisodeNumber          int64        `json:"episodeNumber"`
	ProductionSerialNumber string       `json:"productionSerialNumber"`
	FeatureLength          bool         `json:"featureLength"`
	StardateFrom           float64      `json:"stardateFrom"`
	StarDateTo             float64      `json:"starDateTo"`
	YearFrom               int64        `json:"yearFrom"`
	YearTo                 int64        `json:"yearTo"`
	UsAirDate              string       `json:"usAirDate"`
	FinalScriptDate        string       `json:"finalScriptDate"`
}

type MainDirectorStruct struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type MoviesStruct struct {
	UID                     string             `json:"uid"`
	Title                   string             `json:"title"`
	MainDirector            MainDirectorStruct `json:"mainDirector"`
	TitleBulgarian          string             `json:"titleBulgarian"`
	TitleCatalan            string             `json:"titleCatalan"`
	TitleChineseTraditional string             `json:"titleChineseTraditional"`
	TitleGerman             string             `json:"titleGerman"`
	TitleItalian            string             `json:"titleItalian"`
	TitleJapanese           string             `json:"titleJapanese"`
	TitlePolish             string             `json:"titlePolish"`
	TitleRussian            string             `json:"titleRussian"`
	TitleSerbian            string             `json:"titleSerbian"`
	TitleSpanish            string             `json:"titleSpanish"`
	StardateFrom            float64            `json:"stardateFrom"`
	StarDateTo              float64            `json:"starDateTo"`
	YearFrom                int64              `json:"yearFrom"`
	YearTo                  int64              `json:"yearTo"`
	UsReleaseDate           string             `json:"usReleaseDate"`
}

type CharacterSpeciesStruct struct {
	UID         string `json:"uid"`
	Name        string `json:"name"`
	Numerator   int64  `json:"numerator"`
	Denominator int64  `json:"denominator"`
}

type TitlesStruct struct {
	UID            string `json:"uid"`
	Name           string `json:"name"`
	MilitaryRank   bool   `json:"militaryRank"`
	FleetRank      bool   `json:"fleetRank"`
	ReligiousTitle bool   `json:"religiousTitle"`
	Position       bool   `json:"position"`
	Mirror         bool   `json:"mirror"`
}

type Character2Struct struct {
	UID                    string                   `json:"uid"`
	Name                   string                   `json:"name"`
	Gender                 string                   `json:"gender"`
	YearOfBirth            int64                    `json:"yearOfBirth"`
	MonthOfBirth           int64                    `json:"monthOfBirth"`
	DayOfBirth             int64                    `json:"dayOfBirth"`
	PlaceOfBirth           string                   `json:"placeOfBirth"`
	YearOfDeath            int64                    `json:"yearOfDeath"`
	MonthOfDeath           int64                    `json:"monthOfDeath"`
	DayOfDeath             int64                    `json:"dayOfDeath"`
	PlaceOfDeath           string                   `json:"placeOfDeath"`
	Height                 float32                  `json:"height"`
	Weight                 float32                  `json:"weight"`
	Deceased               bool                     `json:"deceased"`
	BloodType              string                   `json:"bloodType"`
	MaritalStatus          string                   `json:"string"`
	SerialNumber           string                   `json:"serialNumber"`
	HologramActivationDate int32                    `json:"hologramActivationDate"`
	HologramStatus         string                   `json:"hologramStatus"`
	HologramDateStatus     string                   `json:"hologramDateStatus"`
	Hologram               bool                     `json:"hologram"`
	FictionalCharacter     bool                     `json:"fictionalCharacter"`
	Mirror                 bool                     `json:"mirror"`
	AlternateReality       bool                     `json:"alternateReality"`
	Performers             []PerformerStruct        `json:"performers"`
	Episodes               []EpisodeStruct          `json:"episodes"`
	Movies                 []MoviesStruct           `json:"movies"`
	CharacterSpecies       []CharacterSpeciesStruct `json:"characterSpecies"`
	CharacterRelations     []string                 `json:"characterRelations"`
	Titles                 []TitlesStruct           `json:"titles"`
	Organizations          []string                 `json:"organizations"`
}

type CharacterFullStruct struct {
	Character Character2Struct `json:"character"`
}

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

	// Call API to search character by name
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

	// Check API result
	if resp.StatusCode == 200 {

		var elements int64
		elements = inputJson.Page.NumberOfElements

		// Check if character exists
		if elements > 0 {
			// Character key is inputJson.Characters[0].UID

			// Make another call to get Character Details

			newUrl := characterUrl+inputJson.Characters[0].UID

			req2, err := http.NewRequest("GET", newUrl, nil)
			if err != nil {
				log.Fatal("NewRequest: ", err)
				return
			}
			client2 := &http.Client{}

			resp2, err := client2.Do(req2)
			if err != nil {
				log.Fatal("Do: ", err)
				return
			}

			defer resp2.Body.Close()

			var characterJson CharacterFullStruct

			if err := json.NewDecoder(resp2.Body).Decode(&characterJson); err != nil {
				log.Println(err)
			}

			//Print Species
			if resp.StatusCode == 200 {
				fmt.Println(characterJson.Character.CharacterSpecies[0].Name)
			} else {
				fmt.Println("Error retrieving character species")
			}
		} else {
			fmt.Println("Character not found")
		}
	} else {
		fmt.Println("Error searching for character")
	}
}
