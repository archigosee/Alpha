package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sort"
)

var patients = map[int]Patient{
	1: {ID: 1, Name: "Abebe Kebede", Phone: "0911223344", Age: 34},
	2: {ID: 2, Name: "Sara Tadesse", Phone: "0922334455", Age: 28},
	3: {ID: 3, Name: "Hana Girma", Phone: "0933445566", Age: 45},
}

var appointments = []Appointment{
	{ID: 1, PatientID: 1, Doctor: "Dr. Hana", Time: "09:00 AM"},
	{ID: 2, PatientID: 2, Doctor: "Dr. Abel", Time: "10:30 AM"},
	{ID: 3, PatientID: 3, Doctor: "Dr. Hana", Time: "02:00 PM"},
}

type FeeResponse struct {
    Specialist bool    `json:"specialist"`
    Fee        float64 `json:"fee"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the clinic! You asked for %s using %s", r.URL.Path, r.Method)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK - clinic server is healthy")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Alpha Clinic - Addis Ababa. Open Mon-Sat, 8am to 6pm.")
}

func consultationFee(isSpecialist bool) float64 {
	if isSpecialist {
		return 500
	}
	return 250
}

func findRoom(doctor string) (int, error) {
	if doctor == "Dr. Hana" {
		return 12, nil
	}
	return 0, errors.New("doctor not found")
}

func doctorHandler(w http.ResponseWriter, r *http.Request) {
	doctor := Doctor{Name: "Dr. Hana", Room: 12, Specialist: true, Fee: consultationFee(true)}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(doctor)
	if err != nil {
		http.Error(w, "could not encode doctor", http.StatusInternalServerError)
	}
}

func roomHandler(w http.ResponseWriter, r *http.Request) {
	doctor := r.URL.Query().Get("doctor")

	room, err := findRoom(doctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s is in room %d", doctor, room)
}

func feeHandler(w http.ResponseWriter, r *http.Request) {
    specialist := r.URL.Query().Get("specialist")

    var isSpecialist bool
    switch specialist {
    case "true":
        isSpecialist = true
    case "false":
        isSpecialist = false
    default:
        http.Error(w, "specialist must be true or false", http.StatusBadRequest)
        return
    }

     response := FeeResponse{Specialist: isSpecialist, Fee: consultationFee(isSpecialist)}

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(response)
    if err != nil {
        http.Error(w, "could not encode fee", http.StatusInternalServerError)
    }
}

func patientHandler(w http.ResponseWriter, r *http.Request) {
	idText := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "id must be a number", http.StatusBadRequest)
		return
	}

	p, found := patients[id]
	if !found {
		http.Error(w, "patient not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "could not encode patient", http.StatusInternalServerError)
	}
}

func patientsHandler(w http.ResponseWriter, r *http.Request) {
	list := []Patient{}
	for _, p := range patients {
		list = append(list, p)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, "could not encode patients", http.StatusInternalServerError)
	}
}

func appointmentsHandler(w http.ResponseWriter, r *http.Request) {
	doctor := r.URL.Query().Get("doctor")

	result := appointments
	if doctor != "" {
		filtered := []Appointment{}
		for _, appt := range appointments {
			if appt.Doctor == doctor {
				filtered = append(filtered, appt)
			}
		}
		result = filtered
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, "could not encode appointments", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/doctor", doctorHandler)
	http.HandleFunc("/room", roomHandler)
	http.HandleFunc("/fee", feeHandler)
	http.HandleFunc("/patient", patientHandler)
	http.HandleFunc("/patients", patientsHandler)
	http.HandleFunc("/appointments", appointmentsHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
