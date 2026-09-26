# Patient Management API

REST API สำหรับจัดการข้อมูลผู้ป่วย พัฒนาด้วย Go และ PostgreSQL โดยรองรับการรันผ่าน Docker Compose สำหรับ Development Environment

## Tech Stack

- Go
- Gin
- PostgreSQL
- Docker
- pgAdmin

## Tech Stack Incomplete

- Unit test
- NginX

## Project Structure

```text
.
├── api
│   └── server/
├── configs/
│   └── dev.env
├── constants/
├── database/
├── entities/
├── handlers/
├── repositories/
├── services/
├── dto/
├── helper/
├── setup.sql
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── main.go
```

> โครงสร้างอาจแตกต่างจากนี้เล็กน้อยตามการจัด package ของโปรเจกต์

## Features

- Patient CRUD
- Search / Filter patients
- PostgreSQL database
- Request / Response DTO
- Middleware สำหรับ authentication / request context
- Dockerized development environment
- pgAdmin สำหรับจัดการ PostgreSQL

## Environment

สร้างไฟล์:

```text
configs/dev.env
```

ตัวอย่าง:

```env
ENV=development

SERVER_ADDRESS=:8080
CORS_ALLOWED_ORIGIN=http://localhost:3000

DB_DRIVER=postgres
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_NAME=patient_db

JWT_SECRET=jwt_secret
```

> ค่าจริงควรปรับตาม environment ของแต่ละเครื่อง

## Running with Docker

Start ทุก service:

```bash
docker compose --env-file ./configs/dev.env up -d --build
```

Services:

| Service    | Address               |
| ---------- | --------------------- |
| Go API     | http://localhost:8080 |
| PostgreSQL | localhost:5432        |
| pgAdmin    | http://localhost:5050 |

### Start เฉพาะ Go Server

หลังจากแก้ไข Go code สามารถ build เฉพาะ API ได้:

```bash
docker compose --env-file ./configs/dev.env up -d --build go-server
```

ดู log:

```bash
docker compose --env-file ./configs/dev.env logs -f go-server
```

## PostgreSQL

PostgreSQL ใช้ Docker volume เพื่อเก็บข้อมูล:

```yaml
volumes:
  - postgres-data:/var/lib/postgresql/data
```

ดังนั้นการ recreate `go-server` จะไม่ทำให้ข้อมูล PostgreSQL หาย

Database initialization สามารถกำหนดผ่าน:

```text
setup.sql
```

ซึ่งถูก mount เข้า:

```text
/docker-entrypoint-initdb.d/setup.sql
```

## API

Base URL:

```text
http://localhost:8080
```

ตัวอย่าง endpoint:

```text
GET    /api/v1/patient/search
GET    /api/v1/patient/search/:id
POST   /api/v1/patient
POST   /api/v1/staff/create
POST   /api/v1/staff/login
```

### Create Staff

```http
POST /api/v1/staff/create
Content-Type: application/json
```

ตัวอย่าง request:

```json
{
  "username": "staff_01",
  "password": "12345678",
  "confirm_password": "12345678",
  "hospital_name": "HN000001"
}
```

error response:
Username duplicate

```json
{
  "code": 400,
  "message": "Username already in use"
}
```

success response:
Username duplicate

```json
{
  "username": "misterkk-1",
  "message": "Staff created successfully."
}
```

### Login Staff

```http
POST /api/v1/staff/login
Content-Type: application/json
```

ตัวอย่าง request:

```json
{
  "username": "staff_01",
  "password": "12345678",
  "hospital_name": "HN000001"
}
```

error response:
Username duplicate

```json
{
  "code": 400,
  "message": "Username or Password incorrect"
}
```

success response:
Username duplicate

```json
{
  "token": "token",
  "message": "Login success"
}
```

### Create Patient

```http
POST /api/v1/patient
Content-Type: application/json
```

ตัวอย่าง request:

```json
{
  "first_name_th": "สมชาย",
  "middle_name_th": null,
  "last_name_th": "ใจดี",
  "first_name_en": "Somchai",
  "middle_name_en": null,
  "last_name_en": "Jaidee",
  "date_of_birth": "1995-01-01",
  "email": "somchai@example.com",
  "gender": "MALE",
  "national_id": "1234567890123",
  "passport_id": null,
  "patient_hn": "General Hospital",
  "phone_number": "0812345678",
  "gender": "MALE"
}
```

error response:

```json
{
  "code": 400,
  "message": "Paitent already exists"
}
```

success response:

```json
{
  "message": "Patient created successfully."
}
```

### Get Patient by Id (Protected)

```http
GET /api/v1/patient/search/:id
```

id - can be both of national_id or passport_id
error response:

```json
{
  "code": 401,
  "error": "Invalid or expired token"
}
```

error response:

```json
{
  "code": 404,
  "message": "Patient Not Found"
}
```

success response:

```json
{
  "first_name_th": "สมคริส",
  "middle_name_th": "",
  "last_name_th": "ทันใจ",
  "first_name_en": "somchris",
  "middle_name_en": "",
  "last_name_en": "tunjai",
  "date_of_birth": "1992-08-12T00:00:00Z",
  "patient_hn": "HN000001",
  "national_id": "11111111111",
  "passport_id": "",
  "phone_number": "0990000000",
  "email": "somchris@mail.com",
  "gender": ""
}
```

### Get Patient (Protected)

```http
GET /api/v1/patient/search
```

Query parameters:

| Parameter     | Type         | Description    |
| ------------- | ------------ | -------------- |
| `nationalId`  | `string`     | เลขบัตรประชาชน |
| `passportId`  | `string`     | เลข Passport   |
| `firstName`   | `string`     | ชื่อ           |
| `middleName`  | `string`     | ชื่อกลาง       |
| `lastName`    | `string`     | นามสกุล        |
| `email`       | `string`     | Email          |
| `phoneNumber` | `string`     | เบอร์โทรศัพท์  |
| `dateOfBirth` | `YYYY-MM-DD` | วันเกิด        |

error response:

```json
{
  "code": 401,
  "error": "Invalid or expired token"
}
```

success response:

```json
[
  {
    "first_name_th": "สมคริส",
    "middle_name_th": "",
    "last_name_th": "ทันใจ",
    "first_name_en": "somchris",
    "middle_name_en": "",
    "last_name_en": "tunjai",
    "date_of_birth": "1992-08-12T00:00:00Z",
    "patient_hn": "HN000001",
    "national_id": "11111111111",
    "passport_id": "",
    "phone_number": "0990000000",
    "email": "somchris@mail.com",
    "gender": ""
  },
  {
    "first_name_th": "สมคริส",
    "middle_name_th": "",
    "last_name_th": "ทันใจ",
    "first_name_en": "somchris",
    "middle_name_en": "",
    "last_name_en": "tunjai",
    "date_of_birth": "1992-08-12T00:00:00Z",
    "patient_hn": "HN000001",
    "national_id": "11111111111",
    "passport_id": "",
    "phone_number": "0990000000",
    "email": "somchris@mail.com",
    "gender": ""
  },
  ...
]
```
