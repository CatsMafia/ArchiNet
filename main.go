package main

import (
	"github.com/CatsMafia/LolScroll/api"
	"github.com/CatsMafia/LolScroll/db/documents"
	"github.com/CatsMafia/LolScroll/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
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
		var c documents.SessionDocument
		in := strings.Replace(cookie.Value, "'", "\"", -1)
		json.Unmarshal([]byte(in), &c)
		session_id := c.Id
		res := documents.SessionDocument{}
		err := sessionColletion.FindId(session_id).One(&res)
		if err != nil || res.Name != c.Name {
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
			user := documents.UserDocument{utils.GenerateUserId(), username, pass}
			err = usersCollection.Insert(user)
		}
		ren.Redirect("/login")
	}
}

func post_login_handler(ren render.Render, r *http.Request, w http.ResponseWriter) {
	username := r.FormValue("Username")
	password := r.FormValue("Password")
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
			var session documents.SessionDocument
			for err != nil {
				session_id = utils.GenerateId()
				session = documents.SessionDocument{Id: session_id, Name: username}
				err = sessionColletion.Insert(session)
			}
			jsonB, err := json.Marshal(session)
			if err != nil {
				fmt.Println(err)
			}
			out := strings.Replace(string(jsonB), "\"", "'", -1)
			cookie := &http.Cookie{Name: COOKIE_NAME, Value: out}
			http.SetCookie(w, cookie)
			ren.Redirect("/")
		}
	}

}

func main() {
	m := martini.Classic()
	fmt.Println("Run")
	sesion, err := mgo.Dial("localhost")

	if err != nil {
		fmt.Println("Please install MongoDB")
		panic(err)
	}

	api.LolsCollection = sesion.DB("LolScroll").C("lols")
	usersCollection = sesion.DB("LolScroll").C("users")
	sessionColletion = sesion.DB("LolScroll").C("session")

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "layout",                   // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                       // Output human readable JSON
	}))

	m.Get("/", index)
	m.Post("/api/newlol", api.New_lol)
	m.Get("/api/getlols", api.Get_lol)
	m.Post("/api/putkek", api.Post_put_kek_handler)
	m.Get("/login", get_login_handler)
	m.Post("/login", post_login_handler)
	m.Get("/registration", get_registration_handler)
	m.Post("/registration", post_registration_handler)

	m.Run()
}
