package repository

import (
	"fmt"

	"github.com/SergioML9/caliogo/services/models"
)

// SelectSeries ...
func SelectSeries() (models.Series, error) {
	statement := fmt.Sprintf("SELECT media.id, media.name, media.nicename FROM media INNER JOIN serie ON media.id = serie.media_id WHERE media.type = 1")
	rows, err := DBCon.Query(statement)
	fmt.Println(statement)

	var series models.Series

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var serie models.Serie
		if err := rows.Scan(&serie.ID, &serie.Name, &serie.Nicename); err != nil {
			return nil, err
		}

		series = append(series, serie)
	}

	return series, nil
}

// SelectSerieSeasons ...
func SelectSerieSeasons(serieID int) (models.Seasons, error) {
	statement := fmt.Sprintf("SELECT * FROM season WHERE season.serie_id = %d", serieID)
	rows, err := DBCon.Query(statement)
	fmt.Println(statement)

	var seasons models.Seasons

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var season models.Season
		if err := rows.Scan(&season.ID, &season.Number, &season.SerieID); err != nil {
			return nil, err
		}

		seasons = append(seasons, season)
	}

	return seasons, nil
}

// SelectSeasonEpisodes ...
func SelectSeasonEpisodes(seasonID int) (models.Episodes, error) {
	statement := fmt.Sprintf("SELECT * FROM episode WHERE season_id = %d", seasonID)
	rows, err := DBCon.Query(statement)
	fmt.Println(statement)

	var episodes models.Episodes

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var episode models.Episode
		if err := rows.Scan(&episode.ID, &episode.Name, &episode.Nicename, &episode.Number, &episode.SeasonID, &episode.FormatID); err != nil {
			return nil, err
		}

		episodes = append(episodes, episode)
	}

	return episodes, nil
}
