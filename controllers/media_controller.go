package controllers

import (
	"net/http"
	"strconv"

	"github.com/SergioML9/caliogo/services"
	"github.com/gorilla/mux"
)

func GetSeries(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	responseStatus, response := services.GetSeries()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(response)
}

func GetSerieSeasons(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	params := mux.Vars(r)
	serieID, err := strconv.Atoi(params["serie_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	responseStatus, response := services.GetSerieSeasons(serieID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(response)
}

func GetSeasonEpisodes(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	params := mux.Vars(r)
	seasonID, err := strconv.Atoi(params["season_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	responseStatus, response := services.GetSeasonEpisodes(seasonID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(response)
}

/*func CreateSupply(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestSupply := new(models.Supply)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestSupply)
	responseStatus, supply := services.CreateSupply(requestSupply)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(supply)
}

func ResetSupplies(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	params := mux.Vars(r)
	houseID, err := strconv.Atoi(params["house_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	responseStatus, supply := services.ResetSupplies(houseID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(supply)
}

func UpdateSupply(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestSupply := new(models.Supply)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestSupply)
	responseStatus, supply := services.UpdateSupply(requestSupply)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(supply)
}

func UpdateSupplyShowing(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	params := mux.Vars(r)
	requestHouseID, _ := strconv.Atoi(params["house_id"])
	requestShowing, err := strconv.Atoi(params["showing"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	responseStatus, supply := services.UpdateSupplyShowing(requestShowing, requestHouseID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(supply)
}

func DeleteSupply(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	params := mux.Vars(r)
	requestSupplyID, err := strconv.Atoi(params["supply_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	responseStatus, supply := services.DeleteSupply(requestSupplyID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(supply)
}
*/
