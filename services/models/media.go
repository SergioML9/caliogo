package models

type Media struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nicename string `json:"nicename"`
	Type     int    `json:"type"`
}

type Serie struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nicename string `json:"nicename"`
}

type Season struct {
	ID      int `json:"id"`
	Number  int `json:"number"`
	SerieID int `json:"serie_id"`
}

type Episode struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nicename string `json:"nicename"`
	Number   int    `json:"number"`
	SeasonID int    `json:"season_id"`
	FormatID int    `json:"format_id"`
}

type Medias []Media

type Series []Serie

type Seasons []Season

type Episodes []Episode
