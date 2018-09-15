package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func FizzBuzz(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	num, err := strconv.Atoi(p.ByName("num"))

	// 文字や 0 以下の値であれば 400:"Bad Request" を返却
	if err != nil || num < 1 {
		http.Error(w, fmt.Sprintf("%d:Bad Request", http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var dispOutput string
	for i := 1; i <= num; i++ {
		dispOutput += fmt.Sprintf("%3d is \"%s\"\n", i, checkFizzBuzz(i))
	}
	fmt.Fprintf(w, dispOutput)
}

// 入力値に対してFizz/Buzz/FizzBuzz!/数値を文字列で返却する
func checkFizzBuzz(num int) string {
	var result string

	switch {
	case num%3 == 0 && num%5 == 0:
		result = "FizzBuzz!"
	case num%3 == 0:
		result = "Fizz"
	case num%5 == 0:
		result = "Buzz"
	default:
		result = strconv.Itoa(num)
	}

	return result
}

func main() {
	router := httprouter.New()

	// FizzBuzz
	router.GET("/FizzBuzz/:num", FizzBuzz)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
