package main

import (
	"context"
	"time"

	"github.com/herve-beziat/BenchLab/grpc-service/proto"
	"github.com/herve-beziat/BenchLab/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateSensor crée un nouveau capteur
func (s *server) CreateSensor(ctx context.Context, req *proto.SensorRequest) (*proto.Sensor, error) {
	sensor := store.Sensor{
		Name:     req.Name,
		Type:     store.SensorType(req.Type.String()),
		Location: req.Location,
		Unit:     req.Unit,
		Status:   store.SensorStatus(req.Status.String()),
	}

	if req.LastValue != 0 {
		sensor.LastValue = &req.LastValue
	}

	created, err := s.store.CreateSensor(sensor)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "erreur création capteur : %v", err)
	}

	return sensorToProto(created), nil
}

// GetSensor récupère un capteur par son ID
func (s *server) GetSensor(ctx context.Context, req *proto.SensorId) (*proto.Sensor, error) {
	sensor, err := s.store.GetSensor(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "capteur non trouvé : %v", err)
	}

	return sensorToProto(sensor), nil
}

// sensorToProto convertit un store.Sensor en proto.Sensor
func sensorToProto(s store.Sensor) *proto.Sensor {
	p := &proto.Sensor{
		Id:        s.ID,
		Name:      s.Name,
		Location:  s.Location,
		Unit:      s.Unit,
		CreatedAt: s.CreatedAt.Format(time.RFC3339),
	}

	if s.LastValue != nil {
		p.LastValue = *s.LastValue
	}

	if s.LastReadingAt != nil {
		p.LastReadingAt = s.LastReadingAt.Format(time.RFC3339)
	}

	// Convertir les enums
	switch s.Type {
	case store.SensorTypeTemperature:
		p.Type = proto.SensorType_TEMPERATURE
	case store.SensorTypePressure:
		p.Type = proto.SensorType_PRESSURE
	case store.SensorTypeVibration:
		p.Type = proto.SensorType_VIBRATION
	}

	switch s.Status {
	case store.SensorStatusActive:
		p.Status = proto.SensorStatus_ACTIVE
	case store.SensorStatusInactive:
		p.Status = proto.SensorStatus_INACTIVE
	case store.SensorStatusMaintenance:
		p.Status = proto.SensorStatus_MAINTENANCE
	}

	return p
}
