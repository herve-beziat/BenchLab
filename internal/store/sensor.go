package store

import (
	"database/sql"
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

// GetSensor récupère un capteur par son ID
func (s *Store) GetSensor(id string) (Sensor, error) {
	query := `
		SELECT id, name, type, location, unit, status, last_value, last_reading_at, created_at
		FROM sensors
		WHERE id = $1
	`

	row := s.db.QueryRow(query, id)

	var sensor Sensor
	err := row.Scan(
		&sensor.ID,
		&sensor.Name,
		&sensor.Type,
		&sensor.Location,
		&sensor.Unit,
		&sensor.Status,
		&sensor.LastValue,
		&sensor.LastReadingAt,
		&sensor.CreatedAt,
	)
	if err != nil {
		return Sensor{}, err
	}

	return sensor, nil
}

// ListSensors récupère tous les capteurs
func (s *Store) ListSensors() ([]Sensor, error) {
	query := `
		SELECT id, name, type, location, unit, status, last_value, last_reading_at, created_at
		FROM sensors
		ORDER BY created_at DESC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []Sensor
	for rows.Next() {
		var sensor Sensor
		err := rows.Scan(
			&sensor.ID,
			&sensor.Name,
			&sensor.Type,
			&sensor.Location,
			&sensor.Unit,
			&sensor.Status,
			&sensor.LastValue,
			&sensor.LastReadingAt,
			&sensor.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

// UpdateSensor met à jour un capteur existant
func (s *Store) UpdateSensor(id string, sensor Sensor) (Sensor, error) {
	query := `
		UPDATE sensors
		SET name = $1, type = $2, location = $3, unit = $4, status = $5, last_value = $6, last_reading_at = $7
		WHERE id = $8
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
		id,
	)

	var updated Sensor
	err := row.Scan(
		&updated.ID,
		&updated.Name,
		&updated.Type,
		&updated.Location,
		&updated.Unit,
		&updated.Status,
		&updated.LastValue,
		&updated.LastReadingAt,
		&updated.CreatedAt,
	)
	if err != nil {
		return Sensor{}, err
	}

	return updated, nil
}

// DeleteSensor supprime un capteur par son ID
func (s *Store) DeleteSensor(id string) error {
	query := `DELETE FROM sensors WHERE id = $1`

	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
