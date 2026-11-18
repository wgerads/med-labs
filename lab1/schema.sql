CREATE TABLE patient
(
    id         SERIAL PRIMARY KEY,
    full_name  VARCHAR(150) NOT NULL,
    birth_date DATE         NOT NULL,
    phone      VARCHAR(20),
    policy     VARCHAR(20)  NOT NULL
);

CREATE TABLE doctor
(
    id        SERIAL PRIMARY KEY,
    full_name VARCHAR(150) NOT NULL,
    specialty VARCHAR(100) NOT NULL
);

CREATE TABLE ward
(
    id     SERIAL PRIMARY KEY,
    number INT NOT NULL UNIQUE,
    beds   INT NOT NULL CHECK (beds > 0)
);

CREATE TABLE diagnosis
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(150) NOT NULL,
    mkb_code VARCHAR(20)  NOT NULL
);

CREATE TABLE medication
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    form VARCHAR(100) NOT NULL
);

CREATE TABLE medical_card
(
    id              SERIAL PRIMARY KEY,
    patient_id      INT  NOT NULL REFERENCES patient (id),
    doctor_id       INT  NOT NULL REFERENCES doctor (id),
    ward_id         INT  NOT NULL REFERENCES ward (id),
    diagnosis_id    INT REFERENCES diagnosis (id),
    opened_at       DATE NOT NULL DEFAULT CURRENT_DATE,
    closed_at       DATE,
    recommendations VARCHAR(500),
    CHECK (closed_at IS NULL OR closed_at >= opened_at)
);

CREATE TABLE prescription
(
    id              SERIAL PRIMARY KEY,
    medical_card_id INT         NOT NULL REFERENCES medical_card (id) ON DELETE CASCADE,
    medication_id   INT         NOT NULL REFERENCES medication (id),
    frequency       VARCHAR(50) NOT NULL,
    duration        VARCHAR(50) NOT NULL
);
