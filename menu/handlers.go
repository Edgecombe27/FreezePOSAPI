package menu

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var testmenu Menu
var ingredients []Ingredient
var menuItems []MenuItem
var preparations []Preparation
var groups []Group

func SetTestData() {
	ingredients = append(ingredients, Ingredient{ID: 0, Title: "Chocolate", Price: 1.0})
	ingredients = append(ingredients, Ingredient{ID: 1, Title: "Sugar", Price: 0.5})
	ingredients = append(ingredients, Ingredient{ID: 2, Title: "Cinnamon", Price: 2.0})

	preparations = append(preparations, Preparation{ID: 0, Title: "Shake"})
	preparations = append(preparations, Preparation{ID: 1, Title: "Stir"})

	ing1 := []int{0, 1}
	prep1 := []int{0}
	ing2 := []int{1, 2}
	prep2 := []int{1}
	ing3 := []int{0}
	prep3 := []int{}
	group3 := []int{0}
	ing4 := []int{2}
	prep4 := []int{1}
	group4 := []int{1}

	groups = append(groups, Group{ID: 0, Title: "Chocolatey", Ingredients: ing1, Preparations: prep1})
	groups = append(groups, Group{ID: 1, Title: "Cinnamony", Ingredients: ing2, Preparations: prep2})

	menuItems = append(menuItems, MenuItem{ID: 0, Title: "Chocolate Bar", Price: 2.0, Ingredients: ing3, Preparations: prep3, Groups: group3})
	menuItems = append(menuItems, MenuItem{ID: 1, Title: "Cinnamon Bar", Price: 3.0, Ingredients: ing4, Preparations: prep4, Groups: group4})

	testmenu = Menu{Ingredients: ingredients, MenuItems: menuItems, Preparations: preparations, Groups: groups}
}

/* Menu */
func GetMenu(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testmenu)
}

/* Ingredients */

func GetIngredients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testmenu.Ingredients)
}
func GetIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	} else if id > len(testmenu.Ingredients) {
	} else {
		json.NewEncoder(w).Encode(testmenu.Ingredients[id])
	}
}
func NewIngredient(w http.ResponseWriter, r *http.Request) {

}
func DeleteIngredient(w http.ResponseWriter, r *http.Request) {

}

/* Preparations */
func GetPreparations(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testmenu.Preparations)
}
func GetPreparation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	} else if id > len(testmenu.Preparations) {
	} else {
		json.NewEncoder(w).Encode(testmenu.Preparations[id])
	}
}
func NewPreparation(w http.ResponseWriter, r *http.Request) {

}
func DeletePreparation(w http.ResponseWriter, r *http.Request) {

}

/* Groups */
func GetGroups(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testmenu.Groups)
}
func GetGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	} else if id > len(testmenu.Groups) {
	} else {
		json.NewEncoder(w).Encode(testmenu.Groups[id])
	}

}
func NewGroup(w http.ResponseWriter, r *http.Request) {

}
func DeleteGroup(w http.ResponseWriter, r *http.Request) {

}

/* MenuItems */
func GetMenuItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(testmenu.MenuItems)
}
func GetMenuItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
	} else if id > len(testmenu.MenuItems) {
	} else {
		json.NewEncoder(w).Encode(testmenu.MenuItems[id])
	}
}
func NewMenuItem(w http.ResponseWriter, r *http.Request) {

}
func DeleteMenuItem(w http.ResponseWriter, r *http.Request) {

}
