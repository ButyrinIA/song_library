package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"song_library/internal/models"
	"song_library/internal/services"
	"strconv"
)

type AddSongRequest struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}

// @Summary Получение данных песен
// @Description Получение данных библиотеки с фильтрацией и пагинацией
// @Tags Songs
// @Accept json
// @Produce json
// @Param group query string false "Название группы"
// @Param song query string false "Название песни"
// @Param release_date query string false "Дата релиза"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество записей на странице" default(10)
// @Success 200 {array} models.Song
// @Failure 500 {object} map[string]string
// @Router /api/songs/ [get]
func GetSongs(c *gin.Context) {
	filter := map[string]string{
		"group":        c.Query("group"),
		"song":         c.Query("song"),
		"release_date": c.Query("release_date"),
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	songs, err := services.GetAllSongs(filter, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, songs)
}

// @Summary Добавление новой песни
// @Description Добавление новой песни с информацией из API
// @Tags Songs
// @Accept json
// @Produce json
// @Param song body AddSongRequest true "Данные песни"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs [post]
func AddSong(c *gin.Context) {
	var request AddSongRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос: " + err.Error()})
		return
	}

	if err := services.AddSong(request.Group, request.Song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Песня успешно добавлена"})
}

// @Summary Изменение данных песни
// @Description Обновление информации о песне по ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.Song true "Обновленные данные песни"
// @Success 200 {object} models.Song
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs/{id} [put]
func UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var song models.Song

	if err := c.ShouldBindJSON(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	song.ID = uint(id)

	if err := services.UpdateSong(&song); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, song)
}

// @Summary Удаление песни
// @Description Удаление песни по ID
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs/{id} [delete]
func DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteSong(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Песня удалена"})
}

// @Summary Получение текста песни
// @Description Получение текста песни с пагинацией по куплетам
// @Tags Songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество куплетов на странице" default(1)
// @Success 200 {array} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/songs/{id}/text [get]
func GetSongText(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "1"))
	if err != nil || limit < 1 {
		limit = 1
	}

	verses, err := services.GetSongText(id, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, verses)
}
