package main

import (
	"github.com/CatsMafia/ArchiNet/db/documents"
	"github.com/CatsMafia/ArchiNet/models"
	"github.com/CatsMafia/ArchiNet/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"

	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var keksCollection *mgo.Collection

func index(rend render.Render) {
	rend.HTML(200, "index", "")
}

func new_kek(r *http.Request) {
	id := utils.GenerateId(false)
	userId := r.FormValue("userId")
	text := r.FormValue("text")
	t := time.Now()
	kek := documents.KekDocument{id, userId, text, 0, t}
	keksCollection.Insert(kek)
}

func get_kek(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
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
		fmt.Println("AAAAAAAa1")
		keksDoc := []documents.KekDocument{}
		err := keksCollection.Find(nil).All(&keksDoc)
		if err != nil {
			ren.JSON(400, "haven't keks")
		} else {
			fmt.Println("AAAAAAAa2")
			var out string = ""
			for _, doc := range keksDoc {
				fmt.Println("AAAAAAAa3")
				kek := models.Kek{doc.Id, doc.UserId, doc.Text, doc.Rate, doc.Date}
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
