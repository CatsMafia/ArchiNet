package api

import (
	"github.com/CatsMafia/LolScroll/db/documents"
	"github.com/CatsMafia/LolScroll/models"
	"github.com/CatsMafia/LolScroll/utils"

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

var LolsCollection *mgo.Collection

func Get_lol(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	startS := r.FormValue("start")
	countS := r.FormValue("count")
	hashtag := r.FormValue("hashtag")
	linkpeople := r.FormValue("linkauthor")
	if startS == "" {
		startS = "0"
	}
	if countS == "" {
		countS = "10"
	}
	if hashtag != "" {
		hashtag = hashtag[1:]
	}
	if linkpeople != "" {
		linkpeople = linkpeople[1:]
	}
	start, _ := strconv.ParseInt(startS, 10, 64)
	count, _ := strconv.ParseInt(countS, 10, 64)
	end := start + count
	if id != "" {
		lolDoc := documents.LolDocument{}
		err := LolsCollection.FindId(id).One(&lolDoc)
		if err == nil {
			lol, _ := json.Marshal(lolDoc)
			ren.JSON(200, string(lol))
		} else {
			ren.JSON(400, "haven't Lols")
		}
	} else {
		LolsDoc := []documents.LolDocument{}
		err := LolsCollection.Find(nil).All(&LolsDoc)
		if err != nil {
			ren.JSON(400, "haven't Lols")
		} else {
			var out string = ""
			for i, doc := range LolsDoc {
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
				lol := models.Lol{doc.Id, doc.UserId, doc.Text, doc.Keks, doc.Date, "", doc.Hashtags, doc.LinksPeople, make([]string, 0)}
				lolJson, _ := json.Marshal(lol)
				out += string(lolJson) + "\n"
			}
			ren.JSON(200, out)
		}
	}
}

func New_lol(r *http.Request) {
	userId := r.FormValue("userId")
	text := r.FormValue("text")
	var err error = errors.New("a")
	for err != nil {
		id := utils.GenerateId()
		t := time.Now()
		lol := documents.LolDocument{id, userId, text, 0, t, "", utils.FindSubStr(text, "#", " "), utils.FindSubStr(text, "@", " "), make([]string, 0)}
		err = LolsCollection.Insert(lol)
	}
}

func Post_put_kek_handler(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	deltaKek, _ := strconv.ParseInt(r.FormValue("kek"), 10, 64)
	user := r.FormValue("user")
	lolDoc := documents.LolDocument{}
	err := LolsCollection.FindId(string(id)).One(&lolDoc)
	if err != nil {
		ren.JSON(300, id)
		ren.JSON(300, err)
	}
	if !utils.IsIn(lolDoc.UserKeks, user) && deltaKek > 0 {
		lolDoc.Keks += deltaKek
		lolDoc.UserKeks = append(lolDoc.UserKeks, user)
		err = LolsCollection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"lols": lolDoc.Keks, "userlols": lolDoc.UserKeks}})
		if err != nil {
			ren.JSON(300, err)
		}
	} else if utils.IsIn(lolDoc.UserKeks, user) && deltaKek < 0 {
		lolDoc.Keks += deltaKek
		lolDoc.UserKeks = utils.RemoveElemString(lolDoc.UserKeks, user)
		err = LolsCollection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"lols": lolDoc.Keks, "userlols": lolDoc.UserKeks}})
		if err != nil {
			ren.JSON(300, err)
		}
	}

	ren.JSON(200, "Lols:"+fmt.Sprintf("%d", lolDoc.Keks))
}
