package main

import (
	"github.com/CatsMafia/ArchiNet/db/documents"
	"github.com/CatsMafia/ArchiNet/models"
	"github.com/CatsMafia/ArchiNet/utils"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var keksCollection *mgo.Collection
var usersCollection *mgo.Collection

const (
	COOKIE_NAME = "session"
)

func index(rend render.Render, r *http.Request) {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie == nil {
		rend.Redirect("/login")
	} else {
		id := cookie.Value[:32]
		pass := cookie.Value[32:]
		res := documents.UserDocument{}
		err := usersCollection.FindId(id).One(&res)
		if err != nil || res.Password != pass {
			rend.Redirect("/login")
		} else {
			rend.HTML(200, "index", "")
		}
	}
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
			cookie := &http.Cookie{Name: COOKIE_NAME, Value: res.Id + res.Password}
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

	keksCollection = sesion.DB("ArchiNet").C("keks")
	usersCollection = sesion.DB("ArchiNet").C("users")

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
	m.Get("/login", get_login_handler)
	m.Post("/login", post_login_handler)
	m.Get("/registration", get_registration_handler)
	m.Post("/registration", post_registration_handler)

	m.Run()
}
