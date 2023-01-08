package api

import (
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/yosa12978/MyShirts/internal/api/endpoints/v1"
	"github.com/yosa12978/MyShirts/internal/api/midware"
)

func Init() http.Handler {
	apiv1 := v1.NewAPIv1()

	router := mux.NewRouter().StrictSlash(true)
	router.Use(midware.JsonAPI)
	v1router := router.PathPrefix("/v1").Subrouter()
	shirtsv1 := v1router.PathPrefix("/shirts").Subrouter()
	shirtsv1.HandleFunc("/", apiv1.GetAll).Methods(http.MethodGet)
	shirtsv1.HandleFunc("/{id}", apiv1.GetByID).Methods(http.MethodGet)
	shirtsv1.HandleFunc("/", apiv1.Create).Methods(http.MethodPost)
	shirtsv1.HandleFunc("/{id}", apiv1.Update).Methods(http.MethodPut, http.MethodPatch)
	shirtsv1.HandleFunc("/", apiv1.Delete).Methods(http.MethodDelete)

	return router
}
