package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 引数の取得
	name := flag.String("name", "", "mame is string")
	flag.Parse()

	// Get リクエスト発行
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/Profile/%s", *name))
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
	resp.Body.Close()
}
