# Alpha Clinic — Database Schema Design

This document outlines the relational database schema for the Alpha Clinic booking system using **PostgreSQL**.

---

## 1. Entity Relationship Overview

The clinic database centers around appointment scheduling:
- One **patient** can have multiple **appointments** (1-to-Many).
- One **doctor** can have multiple **appointments** (1-to-Many).
- An **appointment** connects a single patient with a single doctor at a scheduled time.

```
+---------------+              +--------------------+              +---------------+
|   patients    | 1          N |    appointments    | N          1 |    doctors    |
+---------------+--------------+--------------------+--------------+---------------+
| patient_id PK |              | appointment_id  PK |              | doctor_id  PK |
| name          |              | patient_id      FK |              | name          |
| phone         |              | doctor_id       FK |              | specialist    |
| date_of_birth |              | start_at           |              | fee           |
+---------------+              | status             |              +---------------+
                               +--------------------+
```

---

## 2. Table Specifications

### Table: `patients`
Stores registered patient demographic and contact details.

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `patient_id` | `SERIAL` (or `INT GENERATED ALWAYS AS IDENTITY`) | `PRIMARY KEY`, `NOT NULL` | Unique identifier for each patient |
| `name` | `VARCHAR(100)` | `NOT NULL` | Full name of the patient |
| `phone` | `VARCHAR(20)` | `NOT NULL`, `UNIQUE` | Contact phone number |
| `date_of_birth` | `DATE` | `NOT NULL` | Patient birth date |

#### Sample Data
| patient_id | name | phone | date_of_birth |
|---|---|---|---|
| 1 | sara | 0988776666 | 2000-10-20 |
| 2 | kal | 0987654322 | 2000-12-10 |

---

### Table: `doctors`
Stores doctor profiles, medical specialties, and consultation rates.

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `doctor_id` | `SERIAL` (or `INT GENERATED ALWAYS AS IDENTITY`) | `PRIMARY KEY`, `NOT NULL` | Unique identifier for each doctor |
| `name` | `VARCHAR(100)` | `NOT NULL` | Doctor's name (e.g., "Dr. Hana") |
| `specialist` | `VARCHAR(100)` | `NOT NULL` | Area of medical specialization |
| `fee` | `NUMERIC(10, 2)` | `NOT NULL` | Consultation fee in ETB / Birr |

#### Sample Data
| doctor_id | name | specialist | fee |
|---|---|---|---|
| 1 | Dr hana | cardiologist | 250.00 |
| 2 | Dr abel | dermatology | 250.00 |

---

### Table: `appointments`
Tracks scheduled visits between patients and doctors.

| Column | Data Type | Constraints | Description |
|---|---|---|---|
| `appointment_id` | `SERIAL` (or `INT GENERATED ALWAYS AS IDENTITY`) | `PRIMARY KEY`, `NOT NULL` | Unique identifier for each appointment |
| `patient_id` | `INTEGER` | `NOT NULL`, `FOREIGN KEY REFERENCES patients(patient_id) ON DELETE CASCADE` | The patient who booked |
| `doctor_id` | `INTEGER` | `NOT NULL`, `FOREIGN KEY REFERENCES doctors(doctor_id) ON DELETE RESTRICT` | The attending doctor |
| `start_at` | `TIMESTAMPTZ` | `NOT NULL` | Appointment date and start time |
| `status` | `VARCHAR(20)` | `NOT NULL`, `DEFAULT 'booked'` | Status (`booked`, `canceled`, `completed`) |

#### Sample Data
| appointment_id | patient_id | doctor_id | start_at | status |
|---|---|---|---|---|
| 1 | 2 | 1 | 2026-11-10 10:00:00+03 | Canceled |
| 2 | 1 | 2 | 2026-11-10 12:00:00+03 | booked |

---

## 3. Key Relationships & Integrity Rules

1. **Foreign Key `appointments.patient_id` → `patients.patient_id`**
   - Ensures an appointment cannot be created for a nonexistent patient.
   - `ON DELETE CASCADE`: If a test patient record is deleted, associated appointments are removed.
2. **Foreign Key `appointments.doctor_id` → `doctors.doctor_id`**
   - Ensures appointments are only booked with valid registered doctors.
   - `ON DELETE RESTRICT`: Prevents accidentally deleting a doctor who has scheduled appointments.
3. **`NOT NULL` across required attributes**
   - Enforces complete data integrity — an appointment must always have an assigned patient, doctor, date/time, and status.

