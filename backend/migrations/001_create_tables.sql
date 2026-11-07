
CREATE TABLE patients (
    id            SERIAL PRIMARY KEY,
    name          TEXT NOT NULL,
    phone         TEXT,
    date_of_birth DATE NOT NULL
);

CREATE TABLE doctors (
    id        SERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    specialty TEXT NOT NULL,
    fee       NUMERIC(10, 2) NOT NULL
);

CREATE TABLE appointments (
    id         SERIAL PRIMARY KEY,
    patient_id INTEGER NOT NULL REFERENCES patients(id),
    doctor_id  INTEGER NOT NULL REFERENCES doctors(id),
    start_at   TIMESTAMP NOT NULL,
    status     TEXT NOT NULL DEFAULT 'booked' CHECK (status IN ('booked', 'cancelled', 'completed'))
);

