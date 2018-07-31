package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Feedback struct {
	Text  string `json:"text"`
	Grade int    `json:"grade"`
	Card  Card   `json:"card"`
}

type Card int

const (
	NOCARD Card = iota + 1
	YELLOWCARD
	REDCARD
)

func NewFeedback(text string, grade int, card Card) Feedback {
	return Feedback{text, grade, card}
}

//Stores Feedback struct to PostNumber in User Folder
func StoreFeedback(matrikel int, postNr int, fb Feedback) error {
	jsondata, err := json.Marshal(fb)
	if err != nil {
		return err
	}
	path := fmt.Sprintf("./Userdata/Students/%v/post_%v_feedback.json", matrikel, postNr)
	err = ioutil.WriteFile(path, jsondata, 0777)
	if err != nil {
		return err
	}
	return nil
}
