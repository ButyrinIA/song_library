package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "song_library/docs"
	"song_library/internal/database"
	"song_library/internal/handlers"
)

// @title Songs Library API
// @version 1.0
// @description API для управления библиотекой песен
// @host localhost:8080
// @BasePath /api/songs

func main() {

	database.InitDB()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/api/songs", handlers.GetSongs)
	r.POST("/api/songs", handlers.AddSong)
	r.PUT("/api/songs/:id", handlers.UpdateSong)
	r.DELETE("/api/songs/:id", handlers.DeleteSong)
	r.GET("/api/songs/:id/text", handlers.GetSongText)

	log.Println("Запуск сервера на порту 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера на порту 8080: %v", err)
	}

}
