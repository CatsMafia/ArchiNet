package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"fmt"
	"net/http"
	"time"

	"github.com/Smeshnyavochki/models"
	"github.com/Smeshnyavochki/utils"
)

var shyteyki map[string]*models.Shyt

func index_handler(r render.Render) {
	r.HTML(200, "index", shyteyki)
}

func save_shyt_handler(rnd render.Render, r *http.Request) {
	id := utils.GenerateId()
	text := r.FormValue("text")
	date := time.Now()
	shyt := models.NewShyt(id, 0, text, date)
	shyteyki[id] = shyt
	rnd.Redirect("/")
}

func main() {
	m := martini.Classic()
	shyteyki = make(map[string]*models.Shyt, 0)
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))
	staticOptions := martini.StaticOptions{Prefix: "static"}
	m.Use(martini.Static("static", staticOptions))
	m.Get("/", index_handler)
	m.Post("/save-shyteyky", save_shyt_handler)
	m.Run()
}
