package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var apiKey = os.Getenv("API_KEY")

func main() {
	http.HandleFunc("/api/route.aspx", handler)
	http.HandleFunc("/api/etd.aspx", handler)

	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	href := ""
	var value interface{} = nil

	switch req.Form["cmd"][0] {
	case "routeinfo", "routes":
		href = "http://api.bart.gov/api/route.aspx"
		value = &RouteInfo{}

	case "etd":
		href = "http://api.bart.gov/api/etd.aspx"
		value = &Station{}
	}

	query := req.RequestURI
	if offset := strings.Index(query, "?"); offset > 0 {
		href = href + query[offset:]
	}
	fmt.Println(href)
	if req.Form["key"] == nil || len(req.Form["key"]) == 0 {
		href = href + "&key=" + apiKey
	}

	res, err := http.Get(href)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	if value == nil {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusOK)

	} else {
		err = xml.NewDecoder(res.Body).Decode(value)
		if err != nil {
			handleErr(w, err)
			return
		}

		bytes, err := json.Marshal(value)
		if err != nil {
			handleErr(w, err)
			return
		}

		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(bytes)))
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
	}
}

func handleErr(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
