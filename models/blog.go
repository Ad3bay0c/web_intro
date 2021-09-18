package models

import (
	"encoding/json"
	"io/ioutil"
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
	Message		string	`json:"message"`
	Blogs		Blog	`json:"blogs"`
	Color		string	`json:"color"`
}

type Blogs struct {
	Blogs []Blog	`json:"blogs"`
}

func (b Blogs) addToFile() error {
	res, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("blog.csv", res, 0666)

	return err
}