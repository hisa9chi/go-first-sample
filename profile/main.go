package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserProfile struct {
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
	FavoriteFoods []string `json:"favaorite_foods"`
}

// 初期登録ユーザ
var users []UserProfile = []UserProfile{
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

// プロフィールを登録する
func PostProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reqProfile UserProfile
	err = json.Unmarshal(body, &reqProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if reqProfile.Name == "" {
		http.Error(w, "name is required.", http.StatusBadRequest)
		return
	}

	if reqProfile.Age < 0 {
		http.Error(w, "age greater than or equal to 0.", http.StatusBadRequest)
		return
	}

	if reqProfile.Gender == "" {
		http.Error(w, "gender is required.", http.StatusBadRequest)
		return
	}

	if checkUsers(reqProfile.Name) != nil {
		http.Error(w, "user is already registed.", http.StatusBadRequest)
		return
	}

	users = append(users, reqProfile)

	fmt.Fprintf(w, fmt.Sprintf("%d Created", http.StatusCreated))
}

// 指定されたname が登録されているユーザであればプロフィールを返す
func GetProfile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	name := p.ByName("name")

	resUserProfile := checkUsers(name)

	if resUserProfile == nil {
		http.Error(w, fmt.Sprintf("%d Not Found", http.StatusNotFound), http.StatusNotFound)
		return
	}

	// json オブジェクトへ変換
	bytes, err := json.Marshal(resUserProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	fmt.Fprintf(w, string(bytes))
}

// 指定されたnameと同じuser profile を取得する
func checkUsers(name string) *UserProfile {
	for _, user := range users {
		if name == user.Name {
			return &user
		}
	}

	return nil
}

func main() {
	router := httprouter.New()

	// 取得
	router.GET("/Profile/:name", GetProfile)
	// 登録
	router.POST("/Profile", PostProfile)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
