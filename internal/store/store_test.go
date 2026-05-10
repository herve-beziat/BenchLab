package store

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestConnection(t *testing.T) {
	godotenv.Load("../../.env")

	s, err := New()
	if err != nil {
		t.Fatalf("impossible de se connecter à la BDD : %v", err)
	}
	defer s.Close()

	t.Log("connexion PostgreSQL OK")
}

func TestCreateSensor(t *testing.T) {
	godotenv.Load("../../.env")

	s, err := New()
	if err != nil {
		t.Fatalf("connexion échouée : %v", err)
	}
	defer s.Close()

	lastValue := 42.5
	sensor, err := s.CreateSensor(Sensor{
		Name:      "Turbine-A3-Temp",
		Type:      SensorTypeTemperature,
		Location:  "Bâtiment C - Salle 12",
		Unit:      "°C",
		Status:    SensorStatusActive,
		LastValue: &lastValue,
	})
	if err != nil {
		t.Fatalf("CreateSensor échoué : %v", err)
	}

	t.Logf("capteur créé : %s", sensor.ID)
}
