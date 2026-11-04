package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Doctor struct {
	Name       string  `json:"name"`
	Room       int     `json:"room"`
	Specialist bool    `json:"specialist"`
	Fee        float64 `json:"fee"`
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

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/doctor", doctorHandler)
	http.HandleFunc("/room", roomHandler)
	http.HandleFunc("/fee", feeHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
