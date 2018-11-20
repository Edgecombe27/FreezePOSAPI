package main

import (
	"log"
	"net/http"

	"github.com/edgecombe27/freeze_pos_api/menu"
	"github.com/gorilla/mux"
)

func main() {
	menu.SetTestData()
	router := mux.NewRouter()
	router.HandleFunc("/menu", menu.GetMenu).Methods("GET")
	router.HandleFunc("menu/ingredients/{id}", menu.GetIngredient).Methods("GET")
	router.HandleFunc("menu/ingredients/{id}", menu.NewIngredient).Methods("POST")
	router.HandleFunc("menu/ingredients/{id}", menu.DeleteIngredient).Methods("DELETE")
	router.HandleFunc("menu/preparations/{id}", menu.GetPreparation).Methods("GET")
	router.HandleFunc("menu/preparations/{id}", menu.NewPreparation).Methods("POST")
	router.HandleFunc("menu/preparations/{id}", menu.DeletePreparation).Methods("DELETE")
	router.HandleFunc("menu/groups/{id}", menu.GetGroup).Methods("GET")
	router.HandleFunc("menu/groups/{id}", menu.NewGroup).Methods("POST")
	router.HandleFunc("menu/groups/{id}", menu.DeleteGroup).Methods("DELETE")
	router.HandleFunc("menu/menuitems/{id}", menu.GetMenuItem).Methods("GET")
	router.HandleFunc("menu/menuitems/{id}", menu.NewMenuItem).Methods("POST")
	router.HandleFunc("menu/menuitems/{id}", menu.DeleteMenuItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
