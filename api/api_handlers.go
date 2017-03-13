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

var KeksCollection *mgo.Collection

func Get_kek(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	startS := r.FormValue("start")
	endS := r.FormValue("end")
	hashtag := r.FormValue("hashtag")
	linkpeople := r.FormValue("linkauthor")
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
		linkpeople = linkpeople[1:]
	}
	start, _ := strconv.ParseInt(startS, 10, 64)
	end, _ := strconv.ParseInt(endS, 10, 64)
	if id != "" {
		kekDoc := documents.KekDocument{}
		err := KeksCollection.FindId(id).One(&kekDoc)
		if err == nil {
			kek, _ := json.Marshal(kekDoc)
			ren.JSON(200, string(kek))
		} else {
			ren.JSON(400, "haven't Keks")
		}
	} else {
		KeksDoc := []documents.KekDocument{}
		err := KeksCollection.Find(nil).All(&KeksDoc)
		if err != nil {
			ren.JSON(400, "haven't Keks")
		} else {
			var out string = ""
			for i, doc := range KeksDoc {
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
				kek := models.Kek{doc.Id, doc.UserId, doc.Text, doc.Lols, doc.Date, "", doc.Hashtags, doc.LinksPeople, make([]string, 0)}
				kekJson, _ := json.Marshal(kek)
				out += string(kekJson) + "\n"
			}
			ren.JSON(200, out)
		}
	}
}

func New_kek(r *http.Request) {
	userId := r.FormValue("userId")
	text := r.FormValue("text")
	var err error = errors.New("a")
	for err != nil {
		id := utils.GenerateId()
		t := time.Now()
		kek := documents.KekDocument{id, userId, text, 0, t, "", utils.FindSubStr(text, "#", " "), utils.FindSubStr(text, "@", " "), make([]string, 0)}
		err = KeksCollection.Insert(kek)
		if err != nil {
			fmt.Println("Error add new kek : ", err)
		}

	}
}

func Post_put_lol_handler(ren render.Render, r *http.Request) {
	id := r.FormValue("id")
	deltaLol, _ := strconv.ParseInt(r.FormValue("lol"), 10, 64)
	user := r.FormValue("user")
	kekDoc := documents.KekDocument{}
	err := KeksCollection.FindId(string(id)).One(&kekDoc)
	if err != nil {
		fmt.Println(id)
		fmt.Println(err)
	}
	if !utils.IsIn(kekDoc.UserLols, user) && deltaLol > 0 {
		kekDoc.Lols += deltaLol
		kekDoc.UserLols = append(kekDoc.UserLols, user)
		err = KeksCollection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"lols": kekDoc.Lols, "userlols": kekDoc.UserLols}})
		if err != nil {
			fmt.Println(err)
		}
	} else if utils.IsIn(kekDoc.UserLols, user) && deltaLol < 0 {
		kekDoc.Lols += deltaLol
		kekDoc.UserLols = utils.RemoveElemString(kekDoc.UserLols, user)
		err = KeksCollection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"lols": kekDoc.Lols, "userlols": kekDoc.UserLols}})
		if err != nil {
			fmt.Println(err)
		}
	}

	ren.JSON(200, "Lols:"+fmt.Sprintf("%d", kekDoc.Lols))
}
