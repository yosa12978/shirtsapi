package v1

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yosa12978/MyShirts/internal/api/dtos"
	"github.com/yosa12978/MyShirts/internal/repos"
	"github.com/yosa12978/MyShirts/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func (a *APIv1) GetShirts(w http.ResponseWriter, r *http.Request) {
	shirtService := services.NewShirtService(repos.NewShirtRepoMongo())
	shirts, err := shirtService.GetShirts()
	if err != nil {
		log.Printf("%s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("{\"status_code\": \"500\"}"))
		return
	}
	log.Printf("Returning all shirts")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(shirts)
}

func (a *APIv1) GetShirtByID(w http.ResponseWriter, r *http.Request) {
	shirtService := services.NewShirtService(repos.NewShirtRepoMongo())
	vars := mux.Vars(r)
	shirt, err := shirtService.GetShirtByID(vars["id"])
	if err != nil {
		log.Printf("%s", err.Error())
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(404)
			w.Write([]byte("{\"status_code\": \"404\"}"))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte("{\"status_code\": \"500\"}"))
		return
	}
	log.Printf("Returning shirt with id %s", vars["id"])
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(shirt)
}

func (a *APIv1) CreateShirt(w http.ResponseWriter, r *http.Request) {
	shirtService := services.NewShirtService(repos.NewShirtRepoMongo())
	bodyb, _ := io.ReadAll(r.Body)
	var body dtos.ShirtCreateDTO
	json.Unmarshal(bodyb, &body)
	shirt, err := body.Map()
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"status_code\": \"400\"}"))
		return
	}
	err = shirtService.AddShirt(shirt)
	if err != nil {
		log.Printf("%s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("{\"status_code\": \"500\"}"))
		return
	}
	log.Printf("Creating new shirt")
	w.WriteHeader(201)
	w.Write([]byte("{\"status_code\": \"201\"}"))
}

func (a *APIv1) UpdateShirt(w http.ResponseWriter, r *http.Request) {
	shirtService := services.NewShirtService(repos.NewShirtRepoMongo())
	vars := mux.Vars(r)

	oldShirt, err := shirtService.GetShirtByID(vars["id"])
	if err != nil {
		log.Printf("%s", err.Error())
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(404)
			w.Write([]byte("{\"status_code\": \"404\"}"))
			return
		}
		w.WriteHeader(500)
		w.Write([]byte("{\"status_code\": \"500\"}"))
		return
	}

	bodyb, _ := io.ReadAll(r.Body)
	var body dtos.ShirtUpdateDTO
	json.Unmarshal(bodyb, &body)

	err = shirtService.UpdateShirt(*body.Map(&oldShirt))
	if err != nil {
		log.Printf("%s", err.Error())
		w.WriteHeader(500)
		w.Write([]byte("{\"status_code\": \"500\"}"))
		return
	}
	log.Printf("Updating shirt with id %s", vars["id"])
	w.WriteHeader(201)
	w.Write([]byte("{\"status_code\": \"201\"}"))
}

func (a *APIv1) DeleteShirt(w http.ResponseWriter, r *http.Request) {
	shirtService := services.NewShirtService(repos.NewShirtRepoMongo())
	vars := mux.Vars(r)
	err := shirtService.DeleteShirt(vars["id"])
	if err != nil {
		log.Printf("%s", err.Error())
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(404)
			w.Write([]byte("{\"status_code\": \"404\"}"))
		}
		w.WriteHeader(500)
		w.Write([]byte("{\"status_code\": \"500\"}"))
		return
	}
	log.Printf("Deleting shirt with id %s", vars["id"])
	w.WriteHeader(200)
	w.Write([]byte("{\"status_code\": \"200\"}"))
}
