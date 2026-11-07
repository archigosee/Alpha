# Clinic Schema

### patients
- `patient_id` SERIAL PRIMARY KEY NOT NULL
- `name` TEXT NOT NULL
- `phone` TEXT 
- `date_of_birth` DATE NOT NULL

### doctors
- `doctor_id` SERIAL PRIMARY KEY NOT NULL
- `name` TEXT NOT NULL
- `specialty` TEXT NOT NULL
- `fee` NUMERIC(10,2) NOT NULL

### appointments
- `appointment_id` SERIAL PRIMARY KEY NOT NULL
- `patient_id` INTEGER NOT NULL REFERENCES patients(patient_id)
- `doctor_id` INTEGER NOT NULL REFERENCES doctors(doctor_id)
- `start_at` TIMESTAMP NOT NULL
- `status` TEXT NOT NULL
