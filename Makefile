
migrate:
	psql -d songs -f internal/database/migrations/001_create_songs_table.sql
