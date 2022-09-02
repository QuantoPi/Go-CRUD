package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/QuantoPi/Go-CRUD/db"
	"github.com/QuantoPi/Go-CRUD/models"
	"github.com/gorilla/mux"
)

func GetCostumersHandler(w http.ResponseWriter, r *http.Request) {
	var costumers []models.Costumer
	db.DB.Find(&costumers)
	json.NewEncoder(w).Encode(&costumers)
}
func GetCostumerHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var costumer models.Costumer
	db.DB.First(&costumer, params["id"])
	if costumer.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&costumer)
}
func PostCostumerHandler(w http.ResponseWriter, r *http.Request) {
	var costumer models.Costumer

	json.NewDecoder(r.Body).Decode(&costumer)

	newCostumer := db.DB.Create(&costumer)
	err := newCostumer.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&costumer)

	w.Write([]byte("PostCostumerHandler"))
}
func DeleteCostumerHandler(w http.ResponseWriter, r *http.Request) {
	var costumer models.Costumer
	params := mux.Vars(r)
	db.DB.First(&costumer, params["id"])

	newCostumer := db.DB.Create(&costumer)
	err := newCostumer.Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
	}
	db.DB.Unscoped().Delete(&costumer)
	w.WriteHeader(http.StatusOK)
}
func UpdateCostumerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id64, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := uint(id64)
	var costumer models.Costumer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&costumer); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	if db.DB.Where("id = ?", id).Updates(&costumer).RowsAffected == 0 {
		w.Write([]byte("Unable to update"))
		return
	}

	costumer.ID = id
	w.WriteHeader(http.StatusOK)
}
