package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herve-beziat/BenchLab/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	// Charger les variables d'environnement
	if err := godotenv.Load(); err != nil {
		log.Fatal("erreur chargement .env")
	}

	// Connexion à la base de données
	s, err := store.New()
	if err != nil {
		log.Fatalf("erreur connexion BDD : %v", err)
	}
	defer s.Close()

	// Initialiser le router Gin
	router := gin.Default()

	// Routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Démarrer le serveur
	port := os.Getenv("REST_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("REST service démarré sur le port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("erreur démarrage serveur : %v", err)
	}
}
