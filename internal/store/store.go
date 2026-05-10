package store

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Store contient la connexion à la base de données
type Store struct {
	db *sql.DB
}

// New crée une nouvelle instance de Store et ouvre la connexion à PostgreSQL
func New() (*Store, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("erreur ouverture BDD : %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erreur connexion BDD : %w", err)
	}

	return &Store{db: db}, nil
}

// Close ferme la connexion à la base de données
func (s *Store) Close() error {
	return s.db.Close()
}
