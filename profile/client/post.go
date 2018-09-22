package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../models"
)

func main() {
	profile := ParseArgs()

	// 登録対象データの出力
	fmt.Printf("%+v\n", profile)

	// json 形式に変換
	jsonStr, err := json.Marshal(profile)
	if err != nil {
		log.Fatal(err)
	}

	// POST リクエスト
	req, err := http.NewRequest("POST", "http://localhost:8080/Profile", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	resp.Body.Close()
}

// 引数で与えられたプロフィール情報を設定
func ParseArgs() models.UserProfile {
	// 引数の値を格納
	name := flag.String("name", "", "name is string")
	age := flag.Int("age", 0, "age is int")
	gender := flag.String("gender", "", "gender is 'Men' or 'Women'")
	foods := flag.String("favorite_foods", "", "Favorite foods is split ','")
	flag.Parse()

	return models.UserProfile{
		Name:          *name,
		Age:           *age,
		Gender:        *gender,
		FavoriteFoods: strings.Split(*foods, ","),
	}
}
