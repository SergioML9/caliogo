package services

import (
	"encoding/json"
	"net/http"

	"github.com/SergioML9/caliogo/repository"
	"github.com/SergioML9/caliogo/services/models"
)

func GetSeries() (int, []byte) {

	repository.InitDB()
	series, err := repository.SelectSeries()
	repository.CloseDB()

	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	if series == nil {
		series = []models.Serie{}
	}
	response, _ := json.Marshal(series)
	return http.StatusOK, response

}

func GetSerieSeasons(serieID int) (int, []byte) {

	repository.InitDB()
	seasons, err := repository.SelectSerieSeasons(serieID)
	repository.CloseDB()

	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	if seasons == nil {
		seasons = []models.Season{}
	}
	response, _ := json.Marshal(seasons)
	return http.StatusOK, response

}

func GetSeasonEpisodes(seasonID int) (int, []byte) {

	repository.InitDB()
	episodes, err := repository.SelectSeasonEpisodes(seasonID)
	repository.CloseDB()

	if err != nil {
		return http.StatusInternalServerError, []byte("")
	}

	if episodes == nil {
		episodes = []models.Episode{}
	}
	response, _ := json.Marshal(episodes)
	return http.StatusOK, response

}

/*
func CreateSupply(requestSupply *models.Supply) (int, []byte) {

	repository.InitDB()
	supplyID, err := repository.InsertSupply(requestSupply)
	repository.CloseDB()
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		requestSupply.ID = supplyID
		response, _ := json.Marshal(requestSupply)
		return http.StatusOK, response
	}
}

func ResetSupplies(houseID int) (int, []byte) {

	repository.InitDB()
	err := repository.InsertBaseSupply(houseID)
	repository.CloseDB()
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		response, _ := json.Marshal("[]")
		return http.StatusOK, response
	}
}

func UpdateSupply(requestSupply *models.Supply) (int, []byte) {

	repository.InitDB()
	err := repository.UpdateSupply(requestSupply)
	repository.CloseDB()
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		response, _ := json.Marshal(requestSupply)
		return http.StatusOK, response
	}
}

func UpdateSupplyShowing(showing int, houseID int) (int, []byte) {

	repository.InitDB()
	err := repository.UpdateSupplyShowing(showing, houseID)
	repository.CloseDB()
	if err != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		var responseSupply models.Supply
		responseSupply.HouseID = houseID
		responseSupply.Showing = showing
		response, _ := json.Marshal(responseSupply)
		return http.StatusOK, response
	}
}

func DeleteSupply(requestSupplyID int) (int, []byte) {

	repository.InitDB()
	err := repository.DeleteSupply(requestSupplyID)

	repository.CloseDB()

	if err != nil {
		return http.StatusInternalServerError, []byte("")
	} else {
		response, _ := json.Marshal("[]")
		return http.StatusOK, response
	}
}
*/
