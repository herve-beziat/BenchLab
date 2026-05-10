-- Suppression si existant (utile pour rejouer la migration)
DROP TABLE IF EXISTS sensors;
DROP TYPE IF EXISTS sensor_status;
DROP TYPE IF EXISTS sensor_type;

-- Types ENUM
CREATE TYPE sensor_type AS ENUM ('TEMPERATURE', 'PRESSURE', 'VIBRATION');

CREATE TYPE sensor_status AS ENUM ('ACTIVE', 'INACTIVE', 'MAINTENANCE');

-- Table sensors
CREATE TABLE IF NOT EXISTS sensors (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name            VARCHAR(255) NOT NULL,
    type            sensor_type NOT NULL,
    location        VARCHAR(255) NOT NULL,
    unit            VARCHAR(50) NOT NULL,
    status          sensor_status NOT NULL DEFAULT 'ACTIVE',
    last_value      DOUBLE PRECISION,
    last_reading_at TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);