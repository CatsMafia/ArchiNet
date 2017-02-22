package main

import (
	"github.com/CatsMafia/ArchiNet/db/documents"
	"github.com/CatsMafia/ArchiNet/models"
	"github.com/CatsMafia/ArchiNet/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"

	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var keksCollection *mgo.Collection

func index(rend render.Render) {
	rend.HTML(200, "index", "")
}

func new_kek(r *http.Request) {
	userId := r.FormValue("userId")
	text := r.FormValue("text")
	t := time.Now()
	var err error = errors.New("a")
	for err != nil {
		id := utils.GenerateId()
		kek := documents.KekDocument{id, userId, text, 0, t, utils.FindSubStr(text, "#", " "), utils.FindSubStr(text, "@", " ")}
		err = keksCollection.Insert(kek)
	}
}

func get_kek(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	startS := r.FormValue("start")
	endS := r.FormValue("end")
	hashtag := r.FormValue("hashtag")
	linkpeople := r.FormValue("LinkAuthor")
	if startS == "" {
		startS = "0"
	}
	if endS == "" {
		endS = "10"
	}
	if hashtag != "" {
		hashtag = hashtag[1:]
	}
	if linkpeople != "" {
		linkpeople = hashtag[1:]
	}
	start, _ := strconv.ParseInt(startS, 10, 64)
	end, _ := strconv.ParseInt(endS, 10, 64)
	if id != "" {
		kekDoc := documents.KekDocument{}
		err := keksCollection.FindId(id).One(&kekDoc)
		if err == nil {
			kek, _ := json.Marshal(kekDoc)
			ren.JSON(200, string(kek))
		} else {
			ren.JSON(400, "haven't keks")
		}
	} else {
		keksDoc := []documents.KekDocument{}
		err := keksCollection.Find(nil).All(&keksDoc)
		if err != nil {
			ren.JSON(400, "haven't keks")
		} else {
			var out string = ""
			for i, doc := range keksDoc {
				if int64(i) < start {
					continue
				}
				if int64(i) > end {
					break
				}

				hashtags := strings.Split(doc.Hashtags, "#")
				linkspeople := strings.Split(doc.LinksPeople, "@")
				flag, flag2 := utils.IsIn(hashtags, hashtag), utils.IsIn(linkspeople, linkpeople)
				if !flag || !flag2 {
					end++
					continue
				}
				kek := models.Kek{doc.Id, doc.UserId, doc.Text, doc.Rate, doc.Date, doc.Hashtags, doc.LinksPeople}
				kekJson, _ := json.Marshal(kek)
				out += string(kekJson) + "\n"
			}
			ren.JSON(200, out)
		}
	}
}

func main() {
	m := martini.Classic()

	sesion, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	keksCollection = sesion.DB("ArchiNet").C("keks")

	m.Get("/", index)
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))
	m.Post("/api/newkek", new_kek)
	m.Get("/api/getkek", get_kek)
	m.Run()
}
