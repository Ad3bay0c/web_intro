package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

const FILENAME = "blog.csv"

type Blog struct {
	ID			int64		`json:"id"`
	Title		string		`json:"title,omitempty"`
	Details		string		`json:"details,omitempty"`
	Date		time.Time	`json:"date"`
}

type Message struct {
	Message		string			`json:"message"`
	Data		Blogs			`json:"data"`
	Color		string			`json:"color"`
}

type Blogs struct {
	Blogs []Blog	`json:"blogs"`
}

func init() {
	loadFile()
}

func loadFile() {
	res, err := ioutil.ReadFile("blog.csv")

	if err != nil {
		log.Printf(err.Error())
	} else {
		_ = json.Unmarshal(res, &blogs)
	}
}
func (b Blogs) addToFile() error {
	res, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("blog.csv", res, 0666)

	return err
}