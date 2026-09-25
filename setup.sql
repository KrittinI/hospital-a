CREATE DATABASE hospital;

\c gorest;

CREATE TABLE staff (
   id SERIAL PRIMARY KEY,
   username VARCHAR(100) UNIQUE NOT NULL,
   password VARCHAR(255) NOT NULL,
   hospital_name VARCHAR(20)
);

CREATE TABLE patients (
   id SERIAL PRIMARY KEY,
   first_name_th VARCHAR(100) NOT NULL,
   middle_name_th VARCHAR(100) NOT NULL,
   last_name_th VARCHAR(100) NOT NULL,
   first_name_en VARCHAR(100) NOT NULL,
   middle_name_en VARCHAR(100) NOT NULL,
   last_name_en VARCHAR(100) NOT NULL,
   date_of_birth DATE NOT NULL,
   patient_hn VARCHAR(100) NOT NULL
   email VARCHAR(255) NOT NULL,
   phone_number VARCHAR(20),
   gender VARCHAR(20),
   national_id VARCHAR(100),
   passport_id VARCHAR(100),
   CHECK ( national_id IS NOT NULL OR passport_id IS NOT NULL),
   UNIQUE (national_id, patient_hn),
   UNIQUE (passport_id, patient_hn)
);

INSERT INTO patients (
    first_name_th,
    middle_name_th,
    last_name_th,
    first_name_en,
    middle_name_en,
    last_name_en,
    date_of_birth,
    patient_hn,
    email,
    phone_number,
    gender,
    national_id,
    passport_id
) VALUES
(
    'สมชาย',
    NULL,
    'ใจดี',
    'Somchai',
    NULL,
    'Jaidee',
    '1990-05-15',
    'HN000001',
    'somchai@example.com',
    '0812345678',
    'MALE',
    '1101700201234',
    NULL
),
(
    'สมหญิง',
    NULL,
    'รักสุขภาพ',
    'Somying',
    NULL,
    'Raksukap',
    '1992-08-20',
    'HN000002',
    'somying@example.com',
    '0823456789',
    'FEMALE',
    '1101700202345',
    NULL
),
(
    'John',
    'Michael',
    'Smith',
    'John',
    'Michael',
    'Smith',
    '1988-03-10',
    'HN000001',
    'john.smith@example.com',
    '0834567890',
    'MALE',
    NULL,
    'P12345678'
),
(
    'Jane',
    NULL,
    'Doe',
    'Jane',
    NULL,
    'Doe',
    '1995-11-25',
    'HN000003',
    'jane.doe@example.com',
    '0845678901',
    'FEMALE',
    NULL,
    'P87654321'
),
(
    'กิตติ',
    NULL,
    'สุขใจ',
    'Kitti',
    NULL,
    'Sukjai',
    '1985-01-30',
    'HN000004',
    'kitti@example.com',
    '0856789012',
    'MALE',
    '1101700203456',
    NULL
),
(
    'สมชาย',
    NULL,
    'ใจดี',
    'Somchai',
    NULL,
    'Jaidee',
    '1990-05-15',
    'HN000005',
    'somchai@example.com',
    '0812345678',
    'MALE',
    '1101700201234',
    NULL
);