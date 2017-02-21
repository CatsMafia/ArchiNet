package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/CatsMafia/ArchiNet/models"

	"encoding/json"
	"net/http"
)

var posts map[string]*models.Post

func index(rend render.Render) {
	rend.HTML(200, "index", "")
}

func newPost(r *http.Request) {
	text := r.FormValue("text")
	userId := r.FormValue("userId")
	post := models.NewPost(text, userId)
	posts[post.Id] = post
}

func getPost(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	post, _ := json.Marshal(posts[id])
	ren.JSON(200, string(post))
}

func main() {
	m := martini.Classic()
	posts = make(map[string]*models.Post, 0)
	m.Get("/", index)
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))
	m.Post("/api/newpost", newPost)
	m.Get("/api/getpost", getPost)
	m.Run()
}
