package services

import (
	"fmt"
	"song_library/internal/database"
	"song_library/internal/models"
	"song_library/pkg/api"
	"strings"
)

func GetAllSongs(filter map[string]string, page, limit int) ([]models.Song, error) {
	var songs []models.Song

	query := "SELECT * FROM songs WHERE 1=1"
	args := []interface{}{}

	if group, ok := filter["group"]; ok && group != "" {
		query += " AND group ILIKE ?"
		args = append(args, "%"+group+"%")
	}
	if song, ok := filter["song"]; ok && song != "" {
		query += " AND song ILIKE ?"
		args = append(args, "%"+song+"%")
	}
	if date, ok := filter["releaseDate"]; ok && date != "" {
		query += " AND releaseDate = ?"
		args = append(args, date)
	}

	query += " ORDER BY id LIMIT ? OFFSET ?"
	args = append(args, limit, (page-1)*limit)

	if err := database.DB.Select(&songs, query, args...); err != nil {
		return nil, fmt.Errorf("Ошибка получения данных: %w", err)
	}
	return songs, nil
}

func AddSong(group, song string) error {
	existingSong := models.Song{}
	err := database.DB.Get(&existingSong, "SELECT * FROM songs WHERE group = $1 AND song = $2", group, song)
	if err == nil {
		return fmt.Errorf("Песня уже существует")
	}
	songApi, err := api.GettingSongDetails(group, song)
	if err != nil {
		return fmt.Errorf("Не удалось получить данные из API: %w", err)
	}

	_, err = database.DB.Exec("INSERT INTO songs (group, song, text, link, release_date, api_fetched) VALUES ($1, $2, $3, $4, $5, $6)", songApi.Group, songApi.Song, songApi.Text, songApi.Link, songApi.ReleaseDate, songApi.APIFetched)
	if err != nil {
		return fmt.Errorf("Ошибка сохранения песни в базе данных: %w", err)
	}
	return nil
}

func UpdateSong(song *models.Song) error {
	_, err := database.DB.NamedExec(
		"UPDATE songs SET group=:group, song=:song, text=:text, link=:link, releaseDate=:releaseDate WHERE id=:id",
		song)
	return err
}

func DeleteSong(id int) error {
	_, err := database.DB.Exec("DELETE FROM songs WHERE id = $1", id)
	return err
}

func GetSongText(id, page, limit int) ([]string, error) {
	var song models.Song
	if err := database.DB.Get(&song, "SELECT text FROM songs WHERE id = $1", id); err != nil {
		return nil, fmt.Errorf("Песня с ID %d не найдена: %w", id, err)
	}

	verses := strings.Split(song.Text, "\n\n")

	// Определяем границы текущей страницы
	start := (page - 1) * limit
	end := start + limit
	if start > len(verses) {
		return []string{}, nil
	}
	if end > len(verses) {
		end = len(verses)
	}

	return verses[start:end], nil

}
