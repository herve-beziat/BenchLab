package main

import (
	"log"
	"net"
	"os"

	"github.com/herve-beziat/BenchLab/grpc-service/proto"
	"github.com/herve-beziat/BenchLab/internal/store"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// Créer le serveur gRPC
	grpcServer := grpc.NewServer()

	// Enregistrer le service
	proto.RegisterSensorServiceServer(grpcServer, &server{store: s})

	// Activer la reflection (pour les outils de debug comme grpcurl)
	reflection.Register(grpcServer)

	// Démarrer le serveur
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("erreur écoute port : %v", err)
	}

	log.Printf("gRPC service démarré sur le port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("erreur démarrage serveur : %v", err)
	}
}
