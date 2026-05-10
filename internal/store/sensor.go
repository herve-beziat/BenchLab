package store

import (
	"time"
)

// SensorType représente le type d'un capteur
type SensorType string

const (
	SensorTypeTemperature SensorType = "TEMPERATURE"
	SensorTypePressure    SensorType = "PRESSURE"
	SensorTypeVibration   SensorType = "VIBRATION"
)

// SensorStatus représente le statut d'un capteur
type SensorStatus string

const (
	SensorStatusActive      SensorStatus = "ACTIVE"
	SensorStatusInactive    SensorStatus = "INACTIVE"
	SensorStatusMaintenance SensorStatus = "MAINTENANCE"
)

// Sensor représente un capteur industriel
type Sensor struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	Type          SensorType   `json:"type"`
	Location      string       `json:"location"`
	Unit          string       `json:"unit"`
	Status        SensorStatus `json:"status"`
	LastValue     *float64     `json:"last_value"`
	LastReadingAt *time.Time   `json:"last_reading_at"`
	CreatedAt     time.Time    `json:"created_at"`
}

// CreateSensor insère un nouveau capteur en base et retourne le capteur créé
func (s *Store) CreateSensor(sensor Sensor) (Sensor, error) {
	query := `
		INSERT INTO sensors (name, type, location, unit, status, last_value, last_reading_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, type, location, unit, status, last_value, last_reading_at, created_at
	`

	row := s.db.QueryRow(
		query,
		sensor.Name,
		sensor.Type,
		sensor.Location,
		sensor.Unit,
		sensor.Status,
		sensor.LastValue,
		sensor.LastReadingAt,
	)

	var created Sensor
	err := row.Scan(
		&created.ID,
		&created.Name,
		&created.Type,
		&created.Location,
		&created.Unit,
		&created.Status,
		&created.LastValue,
		&created.LastReadingAt,
		&created.CreatedAt,
	)
	if err != nil {
		return Sensor{}, err
	}

	return created, nil
}
