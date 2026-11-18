package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/med-group5591241/med/internal/app/model"
	"gitlab.com/med-group5591241/med/internal/app/repository"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func getPatients(patientRepository *repository.PatientRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patients := patientRepository.GetPatients()
		if len(patients) == 0 {
			w.Write([]byte("Пациентов нет"))
			return
		}
		for _, patient := range patients {
			w.Write([]byte(fmt.Sprintf("%d. %s, полис %s\n", patient.ID, patient.FullName, patient.Policy)))
		}
	}
}

func createPatient(patientRepository *repository.PatientRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		patient := model.Patient{
			FullName:  "Петров Пётр Петрович",
			BirthDate: time.Date(1990, 5, 20, 0, 0, 0, 0, time.UTC),
			Policy:    "9876543210987654",
		}

		err := patientRepository.CreatePatient(patient)
		if err != nil {
			w.Write([]byte("Не удалось создать пациента"))
			return
		}

		w.Write([]byte("Пациент создан"))
	}
}

func getPatient(patientRepository *repository.PatientRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		patientID, _ := strconv.Atoi(vars["id"])

		patient := patientRepository.GetPatient(patientID)
		if patient == nil {
			w.Write([]byte("Пациент не найден"))
			return
		}

		w.Write([]byte(fmt.Sprintf("%s, дата рождения %s, полис %s", patient.FullName, patient.BirthDate.Format("02.01.2006"), patient.Policy)))
		if len(patient.Cards) > 0 {
			w.Write([]byte("\nМедицинские карты:"))
			for _, card := range patient.Cards {
				w.Write([]byte(fmt.Sprintf("\nКарта %d, открыта %s", card.ID, card.OpenedAt.Format("02.01.2006"))))
			}
		}
	}
}

func getWards(wardRepository *repository.WardRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wards := wardRepository.GetWards()
		if len(wards) == 0 {
			w.Write([]byte("Палат нет"))
			return
		}
		for _, ward := range wards {
			w.Write([]byte(fmt.Sprintf("Палата %d: коек %d, свободно %d\n", ward.Number, ward.Beds, ward.FreeBeds)))
		}
	}
}

func createMedicalCard(medicalCardRepository *repository.MedicalCardRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		card := model.MedicalCard{
			PatientID: 1,
			DoctorID:  1,
			WardID:    1,
		}

		err := medicalCardRepository.CreateMedicalCard(card)
		if err != nil {
			w.Write([]byte("Не удалось открыть медицинскую карту: " + err.Error()))
			return
		}

		w.Write([]byte("Медицинская карта открыта"))
	}
}

func getMedicalCard(medicalCardRepository *repository.MedicalCardRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardID, _ := strconv.Atoi(vars["id"])

		card := medicalCardRepository.GetMedicalCard(cardID)
		if card == nil {
			w.Write([]byte("Медицинская карта не найдена"))
			return
		}

		w.Write([]byte(fmt.Sprintf("Карта %d: пациент %d, врач %d, палата %d, открыта %s", card.ID, card.PatientID, card.DoctorID, card.WardID, card.OpenedAt.Format("02.01.2006"))))
		if card.DiagnosisID != nil {
			w.Write([]byte(fmt.Sprintf("\nДиагноз: %d", *card.DiagnosisID)))
		}
		if card.ClosedAt != nil {
			w.Write([]byte(fmt.Sprintf("\nЗакрыта %s", card.ClosedAt.Format("02.01.2006"))))
		}
		if len(card.Prescriptions) > 0 {
			w.Write([]byte("\nНазначения:"))
			for _, prescription := range card.Prescriptions {
				w.Write([]byte(fmt.Sprintf("\n%d. Медикамент %d, %s, %s", prescription.ID, prescription.MedicationID, prescription.Frequency, prescription.Duration)))
			}
		}
	}
}

func updateMedicalCard(medicalCardRepository *repository.MedicalCardRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cardID, _ := strconv.Atoi(vars["id"])

		err := medicalCardRepository.UpdateMedicalCard(cardID, 1, "Покой, обильное питьё, повторный осмотр через неделю")
		if err != nil {
			w.Write([]byte("Не удалось изменить медицинскую карту: " + err.Error()))
			return
		}

		w.Write([]byte("Медицинская карта изменена: указан диагноз, карта закрыта"))
	}
}

func getMedications(medicationRepository *repository.MedicationRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		medications := medicationRepository.GetMedications()
		if len(medications) == 0 {
			w.Write([]byte("Медикаментов нет"))
			return
		}
		for _, medication := range medications {
			w.Write([]byte(fmt.Sprintf("%d. %s, %s\n", medication.ID, medication.Name, medication.Form)))
		}
	}
}

func createPrescription(prescriptionRepository *repository.PrescriptionRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prescription := model.Prescription{
			MedicalCardID: 1,
			MedicationID:  1,
			Frequency:     "3 раза в день",
			Duration:      "7 дней",
		}

		err := prescriptionRepository.CreatePrescription(prescription)
		if err != nil {
			w.Write([]byte("Не удалось назначить медикамент"))
			return
		}

		w.Write([]byte("Медикамент назначен"))
	}
}

func deletePrescription(prescriptionRepository *repository.PrescriptionRepository) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		prescriptionID, _ := strconv.Atoi(vars["id"])

		err := prescriptionRepository.DeletePrescription(prescriptionID)
		if err != nil {
			w.Write([]byte("Не удалось отменить назначение: " + err.Error()))
			return
		}

		w.Write([]byte("Назначение отменено"))
	}
}

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

	r := mux.NewRouter()

	r.HandleFunc("/patients", getPatients(patientRepository)).Methods(http.MethodGet)
	r.HandleFunc("/patients", createPatient(patientRepository)).Methods(http.MethodPost)
	r.HandleFunc("/patients/{id}", getPatient(patientRepository)).Methods(http.MethodGet)

	r.HandleFunc("/wards", getWards(wardRepository)).Methods(http.MethodGet)

	r.HandleFunc("/medical-cards", createMedicalCard(medicalCardRepository)).Methods(http.MethodPost)
	r.HandleFunc("/medical-cards/{id}", getMedicalCard(medicalCardRepository)).Methods(http.MethodGet)
	r.HandleFunc("/medical-cards/{id}", updateMedicalCard(medicalCardRepository)).Methods(http.MethodPatch)

	r.HandleFunc("/medications", getMedications(medicationRepository)).Methods(http.MethodGet)

	r.HandleFunc("/prescriptions", createPrescription(prescriptionRepository)).Methods(http.MethodPost)
	r.HandleFunc("/prescriptions/{id}", deletePrescription(prescriptionRepository)).Methods(http.MethodDelete)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
