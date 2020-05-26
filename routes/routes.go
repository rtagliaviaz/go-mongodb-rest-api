package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"../model"
	"github.com/gorilla/mux"
)

func GetAllChars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := model.GetAll()
	json.NewEncoder(w).Encode(payload)
}

func GetChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	payload := model.GetOne(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func CreateNewChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	var char model.Char
	_ = json.NewDecoder(r.Body).Decode(&char)
	model.InsertOne(char)
	json.NewEncoder(w).Encode(char)
}

func DeleteChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	model.DeleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func EditChar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	var res model.Char
	params := mux.Vars(r)
	_ = json.NewDecoder(r.Body).Decode(&res)
	payload, err := model.Edit(params["id"], res)
	if err != nil {
		log.Println("error")
	}
	log.Println(payload)
	json.NewEncoder(w).Encode(payload)
}
