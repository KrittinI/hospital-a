# Patient Management API

REST API สำหรับจัดการข้อมูลผู้ป่วย พัฒนาด้วย Go และ PostgreSQL โดยรองรับการรันผ่าน Docker Compose สำหรับ Development Environment

## Tech Stack

- Go
- Gin
- PostgreSQL
- Docker
- pgAdmin

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
GET    /patient/search
GET    /patient/search/:id
POST   /patient
POST   /staff/create
POST   /staff/login
```

### Create Patient

```http
POST /patient
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
