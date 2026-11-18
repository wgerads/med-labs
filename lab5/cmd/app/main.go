package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "gitlab.com/med-group5591241/med/docs"
	"gitlab.com/med-group5591241/med/internal/app/handler"
	"gitlab.com/med-group5591241/med/internal/app/repository"
	"gitlab.com/med-group5591241/med/internal/app/service"
)

// @title Med API
// @version 1.0
func main() {
	db, err := sqlx.Connect("postgres", "host=localhost port=5433 user=dev dbname=med sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	patientRepository := repository.NewPatientRepository(db)
	wardRepository := repository.NewWardRepository(db)
	medicalCardRepository := repository.NewMedicalCardRepository(db)
	medicationRepository := repository.NewMedicationRepository(db)
	prescriptionRepository := repository.NewPrescriptionRepository(db)

	patientService := service.NewPatientService(patientRepository)
	wardService := service.NewWardService(wardRepository)
	medicalCardService := service.NewMedicalCardService(medicalCardRepository)
	medicationService := service.NewMedicationService(medicationRepository)
	prescriptionService := service.NewPrescriptionService(prescriptionRepository)

	patientHandler := handler.NewPatientHandler(patientService)
	wardHandler := handler.NewWardHandler(wardService)
	medicalCardHandler := handler.NewMedicalCardHandler(medicalCardService)
	medicationHandler := handler.NewMedicationHandler(medicationService)
	prescriptionHandler := handler.NewPrescriptionHandler(prescriptionService)

	r := mux.NewRouter()

	r.HandleFunc("/patients", patientHandler.GetPatients).Methods(http.MethodGet)
	r.HandleFunc("/patients", patientHandler.CreatePatient).Methods(http.MethodPost)
	r.HandleFunc("/patients/{id}", patientHandler.GetPatient).Methods(http.MethodGet)

	r.HandleFunc("/wards", wardHandler.GetWards).Methods(http.MethodGet)

	r.HandleFunc("/medical-cards", medicalCardHandler.CreateMedicalCard).Methods(http.MethodPost)
	r.HandleFunc("/medical-cards/{id}", medicalCardHandler.GetMedicalCard).Methods(http.MethodGet)
	r.HandleFunc("/medical-cards/{id}", medicalCardHandler.UpdateMedicalCard).Methods(http.MethodPatch)

	r.HandleFunc("/medications", medicationHandler.GetMedications).Methods(http.MethodGet)

	r.HandleFunc("/prescriptions", prescriptionHandler.CreatePrescription).Methods(http.MethodPost)
	r.HandleFunc("/prescriptions/{id}", prescriptionHandler.DeletePrescription).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
