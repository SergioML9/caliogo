package routers

import (
	"github.com/SergioML9/caliogo/controllers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func SetMediaRoutes(router *mux.Router) *mux.Router {
	router.Handle("/media/series",
		negroni.New(
			negroni.HandlerFunc(controllers.GetSeries),
		)).Methods("GET")
	router.Handle("/media/serie/seasons/{serie_id}",
		negroni.New(
			negroni.HandlerFunc(controllers.GetSerieSeasons),
		)).Methods("GET")
	router.Handle("/media/serie/episodes/{season_id}",
		negroni.New(
			negroni.HandlerFunc(controllers.GetSeasonEpisodes),
		)).Methods("GET")
	/*router.Handle("/reset/supplies/{house_id}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.ResetSupplies),
		)).Methods("GET")
	router.Handle("/new/supply",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.CreateSupply),
		)).Methods("POST")
	router.Handle("/update/supply",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.UpdateSupply),
		)).Methods("PUT")
	router.Handle("/update/supply/showing/{house_id}/{showing}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.UpdateSupplyShowing),
		)).Methods("GET")
	router.Handle("/delete/supply/{supply_id}",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.DeleteSupply),
		)).Methods("DELETE")*/

	return router
}
