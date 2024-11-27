package repository

import (
	"gorm.io/gorm"
	"song_library/internal/models"
)

type RepositorySongs struct {
	DB *gorm.DB
}

func (r *RepositorySongs) GetSongs() ([]models.Song, error) {
	var songs []models.Song
	err := r.DB.Find(&songs).Error
	return songs, err
}
