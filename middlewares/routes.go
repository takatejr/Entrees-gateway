package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Route struct {
	Url        string
	Auth       bool
	proxy      map[string]string
	Controller http.HandlerFunc
}

var ROUTES = []Route{
	{
		Url:        "/",
		Auth:       false,
		proxy:      nil,
		Controller: final,
	},
	{
		Url:        "/hello",
		Auth:       true,
		proxy:      nil,
		Controller: final2,
	},
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("Executing finalHandler")
	arr := map[string]string{
		"hehe": "lece z koxem",
	}
	fmt.Println(arr)
	jData, _ := json.Marshal(arr)
	w.WriteHeader(200)
	w.Write(jData)
}

func final2(w http.ResponseWriter, r *http.Request) {
	//log.Print("Executing finalHandler")
	arr := map[string]string{
		"hehe2": "lece z koxem2",
	}
	jData, _ := json.Marshal(arr)
	w.WriteHeader(200)
	w.Write(jData)
}
