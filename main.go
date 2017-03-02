package main

import (
	"github.com/CatsMafia/ArchiNet/api"
	"github.com/CatsMafia/ArchiNet/db/documents"
	"github.com/CatsMafia/ArchiNet/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"errors"
	"fmt"
	"net/http"
)

const (
	COOKIE_NAME = "session"
)

var usersCollection *mgo.Collection
var sessionColletion *mgo.Collection

func index(rend render.Render, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie == nil {
		rend.Redirect("/login")
	} else {
		session_id := cookie.Value
		res := documents.SessionDocument{}
		err := sessionColletion.FindId(session_id).One(&res)
		if err != nil {
			rend.Redirect("/login")
		} else {
			rend.HTML(200, "index", "")
		}
	}
}

func get_login_handler(ren render.Render) {
	ren.HTML(200, "login", nil)
}

func get_registration_handler(ren render.Render) {
	ren.HTML(200, "registration", nil)
}

func post_registration_handler(ren render.Render, r *http.Request) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	if username == "" || password == "" {
		ren.Redirect("/registration")
	} else {
		var err error = errors.New("a")
		pass := utils.GetHash(password)
		for err != nil {
			user := documents.UserDocument{utils.GenerateId(), username, pass}
			err = usersCollection.Insert(user)
		}
		ren.Redirect("/login")
	}
}

func post_login_handler(ren render.Render, r *http.Request, w http.ResponseWriter) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
	fmt.Println(username)
	fmt.Println(password)
	res := documents.UserDocument{}
	err := usersCollection.Find(bson.M{"username": username}).One(&res)
	if err != nil {
		fmt.Println(err)
		ren.Redirect("/login")
	} else {
		if res.Password != utils.GetHash(password) {
			ren.Redirect("/login")
		} else {
			var err error = errors.New("a")
			session_id := ""
			for err != nil {
				session_id = utils.GenerateId()
				session := documents.SessionDocument{Id: session_id}
				err = sessionColletion.Insert(session)
			}
			cookie := &http.Cookie{Name: COOKIE_NAME, Value: session_id}
			http.SetCookie(w, cookie)
			ren.Redirect("/")
		}
	}

}

func main() {
	m := martini.Classic()
	fmt.Println("Hello")
	sesion, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}

	api.KeksCollection = sesion.DB("ArchiNet").C("keks")
	usersCollection = sesion.DB("ArchiNet").C("users")
	sessionColletion = sesion.DB("ArchiNet").C("session")

	m.Get("/", index)
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{AppHelpers}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))

	m.Post("/api/newkek", api.New_kek)
	m.Get("/api/getkek", api.Get_kek)
	m.Post("/api/putkek", api.Post_put_kek_handler)
	m.Get("/login", get_login_handler)
	m.Post("/login", post_login_handler)
	m.Get("/registration", get_registration_handler)
	m.Post("/registration", post_registration_handler)

	m.Run()
}
