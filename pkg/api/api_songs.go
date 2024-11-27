package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"song_library/internal/models"
)

type SongDetail struct {
	Text        string `json:"text"`
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
}

func GettingSongDetails(group, song string) (*models.Song, error) {
	url := fmt.Sprintf("http://localhost:8080/info?group=%s&song=%s", group, song)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Ошибка запроса к API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка API: %s", resp.Status)
	}

	var details SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа API: %w", err)
	}
	return &models.Song{
		Group:       group,
		Song:        song,
		Text:        details.Text,
		Link:        details.Link,
		ReleaseDate: details.ReleaseDate,
		APIFetched:  true,
	}, nil
}
