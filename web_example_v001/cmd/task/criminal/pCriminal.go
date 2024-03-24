package criminal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"tracker/cmd/task/errorhandler"
)

type TypeGender string

const (
	GenderMale   TypeGender = "male"
	GenderFemale TypeGender = "female"
)

type Man struct {
	Name     string     `json:"Name"`
	LastName string     `json:"LastName"`
	Age      uint       `json:"Age"`
	Gender   TypeGender `json:"Gender"`
	Crimes   uint       `json:"Crimes"`
	Photo    string     `json:"Photo"`
}

var people map[string]Man

func LoadFromJSONFile(fileName string) (map[string]Man, error) {

	jsonData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errorhandler.PrintError(fmt.Errorf("ошибка чтения JSON-файла:", err))
	}

	err = json.Unmarshal(jsonData, &people)
	if err != nil {
		return nil, errorhandler.PrintError(fmt.Errorf("ошибка десериализации JSON:", err))
	}

	return people, nil
}

func FindCriminalMan(suspectsStr string) (map[string]Man, string) {

	nameList := strings.Split(suspectsStr, "\n")

	// Удаление пустых строк и лишних пробелов и дублирующих записей
	var suspects []string
	seen := make(map[string]bool) // Используем map для отслеживания уникальных строк
	for _, name := range nameList {
		name = strings.TrimSpace(name)
		if name != "" && !seen[name] {
			seen[name] = true
			suspects = append(suspects, name)
		}
	}

	//возможно несколько записей с одинаковым значением Crimes
	var highestCrimes uint = 0
	highestCrimesPeople := make(map[string]Man)

	for _, suspect := range suspects {
		if person, ok := people[suspect]; ok {
			if person.Crimes > highestCrimes {
				highestCrimes = person.Crimes
				highestCrimesPeople = map[string]Man{suspect: person}
			} else if person.Crimes == highestCrimes {
				highestCrimesPeople[suspect] = person
			}
		}
	}

	response := fmt.Sprintln("The most criminal personalities")
	if len(highestCrimesPeople) == 0 {
		response = fmt.Sprintln("There is no information in the database for the requested suspects")
	}

	return highestCrimesPeople, response

}
