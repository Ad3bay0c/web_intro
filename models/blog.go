package models

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"time"
)

const FILENAME = "blog.csv"

type Blog struct {
	ID			int64		`json:"id"`
	Title		string		`json:"title,omitempty"`
	Details		string		`json:"details,omitempty"`
	Comment		int			`json:"comment"`
	View		int			`json:"view"`
	Date		time.Time	`json:"date"`
}

type Message struct {
	Message		string			`json:"message"`
	Data		Blogs			`json:"data"`
	Color		string			`json:"color"`
	Post		Blog			`json:"post"`
}

type Blogs struct {
	Blogs []Blog	`json:"blogs"`
}
var t *template.Template

func init() {
	loadFile()
	t = template.Must(template.ParseGlob("frontend/*"))
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