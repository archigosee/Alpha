package main

type Doctor struct {
	Name       string  `json:"name"`
	Room       int     `json:"room"`
	Specialist bool    `json:"specialist"`
	Fee        float64 `json:"fee"`
}

type Patient struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

type Appointment struct {
	ID        int    `json:"id"`
	PatientID int    `json:"patient_id"`
	Doctor    string `json:"doctor"`
	Time      string `json:"time"`
}