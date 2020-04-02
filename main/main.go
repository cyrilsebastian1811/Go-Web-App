package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getCurrentTime).Methods("GET")
	router.HandleFunc("/time", getTime).Methods("GET")

	http.ListenAndServe("0.0.0.0:8000", router)
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	t := time.Now()
	json.NewEncoder(w).Encode(t)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	timeInfo := make(map[string]interface{})
	v := r.URL.Query()

	for _, value := range v {
		loc, err := time.LoadLocation(value[0])
		if err != nil {
			fmt.Println(err)
			timeInfo[value[0]] = nil
			break
		}
		timeInfo[value[0]] = time.Now().In(loc)
	}

	json.NewEncoder(w).Encode(timeInfo)
}
