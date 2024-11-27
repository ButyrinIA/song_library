package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Подключение драйвера PostgreSQL
	"log"
	"os"
	"time"
)

var DB *sqlx.DB

// Функция для инициализации базы данных
func InitDB() {
	err := godotenv.Load("song_library/.env")
	if err != nil {
		log.Fatalf("Ошибка загрузки данных из .env: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Ошибка подключения к Базе Данных: %v", err)
	}

	// Настройка пула соединений
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	DB = db
	log.Println("Успешное подключение к Базе Данных")

	migrateDatabase(db)
}

// Функция для выполнения миграций при старте
func migrateDatabase(db *sqlx.DB) {
	m, err := migrate.New(
		"file:///app/internal/database/migrations", // Путь к миграциям внутри контейнера
		fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")),
	)
	if err != nil {
		log.Fatalf("Ошибка при создании миграций: %v", err)
	}

	err = m.Up()
	if err != nil && err.Error() != "no change" {
		log.Fatalf("Ошибка при применении миграций: %v", err)
	}
	log.Println("Миграции успешно применены")
}
