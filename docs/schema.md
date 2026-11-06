# Clinic Schema

### patients
- `patient_id` SERIAL PRIMARY KEY NOT NULL
- `name` TEXT
- `phone` TEXT
- `date_of_birth` TIMESTAMP

### doctors
- `doctor_id` SERIAL PRIMARY KEY NOT NULL
- `name` TEXT
- `specialist` TEXT
- `fee` NUMERIC

### appointments
- `appointment_id` SERIAL PRIMARY KEY NOT NULL
- `patient_id` FK
- `doctor_id` FK
- `start_at` TIMESTAMP
- `status` TEXT
