package config

import "database/sql"

type Config struct {
	MemoryStorage   bool
	DatabaseConnect *sql.DB
}
