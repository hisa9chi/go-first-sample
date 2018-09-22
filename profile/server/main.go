package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"../models"

	"github.com/julienschmidt/httprouter"
)

// html のテンプレート格納相対パス
const tplPath = "resources/templates"

// 初期登録のユーザプロフィール
var users []models.UserProfile = []models.UserProfile{
	{
		Name:          "Bob",
		Age:           25,
		Gender:        "Man",
		FavoriteFoods: []string{"Hanberger", "Cookie", "Chocolate"},
	},
	{
		Name:          "Alice",
		Age:           24,
		Gender:        "Woman",
		FavoriteFoods: []string{"Apple", "Orange", "Melon"},
	},
}

func Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t := template.Must(template.ParseFiles(tplPath + "/index.html.tpl"))
	err := t.ExecuteTemplate(w, "index.html.tpl", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 指定されたname が登録されているユーザであればプロフィールを返す
func GetProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	resUserProfile := CheckUsers(name)

	if resUserProfile == nil {
		http.Error(w, fmt.Sprintf("%d Not Found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	t := template.Must(template.ParseFiles(tplPath + "/profile.html.tpl"))
	err := t.ExecuteTemplate(w, "profile.html.tpl", resUserProfile)
	if err != nil {
		log.Fatal(err)
	}
}

// 指定されたnameと同じuser profile を取得する
func CheckUsers(name string) *models.UserProfile {
	for _, user := range users {
		if name == user.Name {
			return &user
		}
	}

	return nil
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/profile/:name", GetProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
