package main

import (
	"github.com/herve-beziat/BenchLab/grpc-service/proto"
	"github.com/herve-beziat/BenchLab/internal/store"
)

// server implémente l'interface SensorServiceServer générée par protoc
type server struct {
	proto.UnimplementedSensorServiceServer
	store *store.Store
}
