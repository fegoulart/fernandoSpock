package main

import (
	"encoding/json"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"log"
	"bytes"
	"strings"
)

// URLs
var characterSearchUrl = "http://stapi.co/api/v1/rest/character/search?pageSize=2000"
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

type SourceStruct struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type TargetStruct struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type CharacterRelationsStruct struct {
	Type   string       `json:"type"`
	Source SourceStruct `json:"source"`
	Target TargetStruct `json:"target"`
}

type OrganizationStruct struct {
	UID                         string `json:"uid"`
	Name                        string `json:"name"`
	Government                  bool   `json:"government"`
	IntergovernmentOrganization bool   `json:"intergovernmentalOrganization"`
	ResearchOrganization        bool   `json:"researchOrganization"`
	SportOrganization           bool   `json:"sportOrganization"`
	MedicalOrganization         bool   `json:"medicalOrganization"`
	MilitaryOrganization        bool   `json:"militaryOrganization"`
	MilitaryUnit                bool   `json:"militaryUnit"`
	GovernmentAgency            bool   `json:"overnmentAgency"`
	LawEnforcementAgency        bool   `json:"lawEnforcementAgency"`
	PrisonOrPenalColony         bool   `json:"prisonOrPenalColony"`
	Mirror                      bool   `json:"mirror"`
	AlternateReality            bool   `json:"alternateReality"`
}

type Character2Struct struct {
	UID                    string                     `json:"uid"`
	Name                   string                     `json:"name"`
	Gender                 string                     `json:"gender"`
	YearOfBirth            int64                      `json:"yearOfBirth"`
	MonthOfBirth           int64                      `json:"monthOfBirth"`
	DayOfBirth             int64                      `json:"dayOfBirth"`
	PlaceOfBirth           string                     `json:"placeOfBirth"`
	YearOfDeath            int64                      `json:"yearOfDeath"`
	MonthOfDeath           int64                      `json:"monthOfDeath"`
	DayOfDeath             int64                      `json:"dayOfDeath"`
	PlaceOfDeath           string                     `json:"placeOfDeath"`
	Height                 float32                    `json:"height"`
	Weight                 float32                    `json:"weight"`
	Deceased               bool                       `json:"deceased"`
	BloodType              string                     `json:"bloodType"`
	MaritalStatus          string                     `json:"string"`
	SerialNumber           string                     `json:"serialNumber"`
	HologramActivationDate int32                      `json:"hologramActivationDate"`
	HologramStatus         string                     `json:"hologramStatus"`
	HologramDateStatus     string                     `json:"hologramDateStatus"`
	Hologram               bool                       `json:"hologram"`
	FictionalCharacter     bool                       `json:"fictionalCharacter"`
	Mirror                 bool                       `json:"mirror"`
	AlternateReality       bool                       `json:"alternateReality"`
	Performers             []PerformerStruct          `json:"performers"`
	Episodes               []EpisodeStruct            `json:"episodes"`
	Movies                 []MoviesStruct             `json:"movies"`
	CharacterSpecies       []CharacterSpeciesStruct   `json:"characterSpecies"`
	CharacterRelations     []CharacterRelationsStruct `json:"characterRelations"`
	Titles                 []TitlesStruct             `json:"titles"`
	Organizations          []OrganizationStruct       `json:"organizations"`
}

type CharacterFullStruct struct {
	Character Character2Struct `json:"character"`
}

func translate(englishInput string) string {
	charMap := map[string]string{
		"a": "0xF8D0",
		"A": "0xF8D0",
		"b": "0xF8D1",
		"B": "0xF8D1",
		"d": "0xF8D3",
		"D": "0xF8D3",
		"e": "0xF8D4",
		"E": "0xF8D4",
		"h": "0xF8D6",
		"H": "0xF8D6",
		"i": "0xF8D7",
		"I": "0xF8D7",
		"j": "0xF8D8",
		"J": "0xF8D8",
		"l": "0xF8D9",
		"L": "0xF8D9",
		"m": "0xF8DA",
		"M": "0xF8DA",
		"n": "0xF8DB",
		"N": "0xF8DB",
		"O": "0xF8DD",
		"o": "0xF8DD",
		"p": "0xF8DE",
		"P": "0xF8DE",
		"q": "0xF8DF",
		"Q": "0xF8E0",
		"r": "0xF8E1",
		"R": "0xF8E1",
		"s": "0xF8E2",
		"S": "0xF8E2",
		"t": "0xF8E3",
		"T": "0xF8E3",
		"u": "0xF8E5",
		"U": "0xF8E5",
		"v": "0xF8E6",
		"V": "0xF8E6",
		"w": "0xF8E7",
		"W": "0xF8E7",
		"y": "0xF8E8",
		"Y": "0xF8E8",
		"'": "0xF8E9",
		"0": "0xF8F0",
		"1": "0xF8F1",
		"2": "0xF8F2",
		"3": "0xF8F3",
		"4": "0xF8F4",
		"5": "0xF8F5",
		"6": "0xF8F6",
		"7": "0xF8F7",
		"8": "0xF8F8",
		"9": "0xF8F9",
		".": "0xF8FD",
		",": "0xF8FE",
		" ": "0x0020",
	}

	var translationResult = ""

	for i := 0; i < len(englishInput); i++ {
		result := charMap[string(englishInput[i])]
		if result == "" {
			return englishInput
		} else {
			//white space between codes
			if i > 0 {
				result = " " + result
			}
			translationResult = translationResult + result
		}

	}

	return translationResult

}

func checkIfWholeWordMatch(inputName string, apiName string) bool {

	var separatedNames []string

	if strings.ToLower(inputName) == strings.ToLower(apiName) {
		return true
	}

	separatedNames = strings.Split(apiName, " ")

	for i := 0; i < len(separatedNames); i++ {
		if strings.ToLower(separatedNames[i]) == strings.ToLower(inputName) {
			return true
		}
	}
	return false

}

func getSpecies(englishName string) (bool, string) {
	//Call to Character Endpoint

	data := url.Values{}
	data.Set("title", englishName)
	data.Add("name", englishName)

	// Call API to search character by name
	req, err := http.NewRequest("POST", characterSearchUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return false, "HTTP Request Encoding Failed - Character"
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return false, "HTTP Request Failed - Character"
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

			// Find a whole word match
			var characterUID string
			var fullMatchCharacterUID string
			var wordMatched = false

			characterUID = ""

			for i := elements; i > 0; i-- {

				//fmt.Println(inputJson.Characters[i-1].Name)
				wordMatched = checkIfWholeWordMatch(englishName, inputJson.Characters[i-1].Name)
				//fmt.Println(wordMatched)
				if wordMatched {
					if strings.ToLower(englishName) == strings.ToLower(inputJson.Characters[i-1].Name) {
						fullMatchCharacterUID = inputJson.Characters[i-1].UID
					}
					characterUID = inputJson.Characters[i-1].UID
				}
			}

			if characterUID == "" {
				return false, "Character not found"
			}
			// Make another call to get Character Details
			var newUrl string
			if fullMatchCharacterUID != "" {
				newUrl = characterUrl + fullMatchCharacterUID
			} else {
				newUrl = characterUrl + characterUID
			}
			req2, err := http.NewRequest("GET", newUrl, nil)
			if err != nil {
				log.Fatal("NewRequest: ", err)
				return false, "HTTP Request Encoding Failed - Species"
			}
			client2 := &http.Client{}

			resp2, err := client2.Do(req2)
			if err != nil {
				log.Fatal("Do: ", err)
				return false, "HTTP Request Failed - Species"
			}

			defer resp2.Body.Close()

			var characterJson CharacterFullStruct

			if err := json.NewDecoder(resp2.Body).Decode(&characterJson); err != nil {
				log.Println(err)
			}

			//Print Species
			if resp.StatusCode == 200 {
				if len(characterJson.Character.CharacterSpecies) == 0 {
					//fmt.Println("No species informed")
					return false, "No species informed"
				} else {

					//fmt.Println(characterJson.Character.CharacterSpecies[0].Name)
					return true, characterJson.Character.CharacterSpecies[0].Name
				}
			} else {
				//fmt.Println("Error retrieving character species")
				return false, "Error retrieving character species"
			}
		} else {
			//fmt.Println("Character not found")
			return false, "Character not found"
		}
	} else {
		//fmt.Println("Error searching for character")
		return false, "Error searching for character"
	}
	return false, "Undefined Error"
}

func main() {

	//Command line arguments

	argsWithoutProg := os.Args[1:]

	var englishName string
	englishName = ""

	//Concatenate all arguments
	for i := 0; i < len(argsWithoutProg); i++ {
		if i > 0 {
			englishName = englishName + " " + argsWithoutProg[i]
		} else {
			englishName = englishName + argsWithoutProg[i]
		}
	}

	translateResult := translate(englishName)

	fmt.Println(translateResult)


	if translateResult == englishName {

		// Translate Error
		fmt.Println("No translation found")

	} else {

		speciesResult, speciesMessage := getSpecies(englishName)
		if speciesResult == true {
			//ok
			fmt.Println(speciesMessage)
		} else {
			//error
			fmt.Println(speciesMessage)
		}

	}

}
