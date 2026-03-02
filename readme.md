# Go REST API - Customer Management

Aplikasi REST API menggunakan Golang, Gofiber, dan MySQL dengan Clean Architecture.

## Instalasi

1. Install Gofiber
   ```bash
   go get github.com/gofiber/fiber/v2
   ```

2. Install Goqu (Query Builder)
   ```bash
   go get github.com/doug-martin/goqu/v9
   go get github.com/doug-martin/goqu/v9/dialect/mysql
   ```

3. Install Godotenv (Environment Variables)
   ```bash
   go get github.com/joho/godotenv
   ```

4. Install MySQL Driver
   ```bash
   go get github.com/go-sql-driver/mysql
   ```

## Setup Database

1. Buat database dan table:
   ```sql
   CREATE DATABASE go_restAPI;
   USE go_restAPI;
   
   CREATE TABLE customers (
       id VARCHAR(36) PRIMARY KEY,
       code VARCHAR(50) NOT NULL UNIQUE,
       name VARCHAR(255) NOT NULL,
       created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
       deleted_at TIMESTAMP NULL DEFAULT NULL,
       INDEX idx_deleted_at (deleted_at)
   );
   
   -- Insert sample data
   INSERT INTO customers (id, code, name) VALUES 
   ('1', 'CUST001', 'John Doe'),
   ('2', 'CUST002', 'Jane Smith'),
   ('3', 'CUST003', 'Bob Johnson');
   ```

2. Konfigurasi `.env`:
   ```env
   SERVER_HOST=localhost
   SERVER_PORT=8000
   DATABASE_HOST=localhost
   DATABASE_PORT=3306
   DATABASE_NAME=go_restAPI
   DATABASE_USER=root
   DATABASE_PASS=
   TZ=Asia/Jakarta
   ```

## Jalankan Aplikasi

```bash
go run main.go
```

Server berjalan di: `http://localhost:8000`

## API Endpoints

### GET /customers
Mengambil semua data customer yang aktif (belum dihapus)

**Request:**
```bash
curl http://localhost:8000/customers
```

**Response Success:**
```json
{
  "code": "00",
  "message": "Success",
  "data": [
    {
      "id": "1",
      "code": "CUST001",
      "name": "John Doe"
    },
    {
      "id": "2",
      "code": "CUST002",
      "name": "Jane Smith"
    }
  ]
}
```

**Response Error:**
```json
{
  "code": "99",
  "message": "error message",
  "data": ""
}
```

## Arsitektur Project

```
User Request → API → Service → Repository → Database
     ↓         ↓       ↓          ↓            ↓
  HTTP      Handler  Logic    Query SQL    MySQL
```

### Struktur Folder
```
Go-RestAPI/
├── domain/              # Entity & Interface
│   └── customer.go
├── dto/                 # Data Transfer Object
│   ├── customer_data.go
│   └── response.go
├── internal/
│   ├── api/            # HTTP Handlers
│   │   └── customer.go
│   ├── service/        # Business Logic
│   │   └── customer.go
│   ├── repository/     # Database Operations
│   │   └── customer.go
│   ├── config/         # Configuration
│   │   ├── loader.go
│   │   └── model.go
│   └── connection/     # Database Connection
│       └── database.go
├── .env                # Environment Variables
├── go.mod
└── main.go
```

### Layer Explanation

**1. API Layer** (`internal/api`)
- Handle HTTP request/response
- Validasi input
- Set timeout context
- Format response JSON

**2. Service Layer** (`internal/service`)
- Business logic
- Convert entity ke DTO
- Orchestrate repository calls

**3. Repository Layer** (`internal/repository`)
- Database operations (CRUD)
- Query SQL menggunakan Goqu
- Convert database result ke entity

**4. Domain Layer** (`domain`)
- Define interface (contract)
- Define entity struct

**5. DTO Layer** (`dto`)
- Struct untuk response JSON
- Pisahkan data internal dengan data yang dikirim ke client

## Flow Request GET /customers

```
1. User → GET http://localhost:8000/customers

2. API (customer.go Index)
   - Terima request
   - Bikin context dengan timeout 10 detik
   - Panggil service.Index(ctx)

3. Service (customer.go Index)
   - Panggil repository.FindAll(ctx)
   - Convert []Customer ke []CustomerData (DTO)
   - Return data

4. Repository (customer.go FindAll)
   - Query: SELECT * FROM customers WHERE deleted_at IS NULL
   - Scan result ke []Customer
   - Return data

5. Database (MySQL)
   - Execute query
   - Return rows

6. Response ke User
   {
     "code": "00",
     "message": "Success",
     "data": [...]
   }
```

## Fitur

- ✅ Get all customers (soft delete aware)
- ✅ Clean Architecture (separation of concerns)
- ✅ Environment configuration (.env)
- ✅ Context timeout (10 seconds)
- ✅ Soft delete implementation
- ✅ Generic response wrapper
- ✅ Query builder (Goqu)

## Tech Stack

- **Framework**: Gofiber v2
- **Database**: MySQL
- **Query Builder**: Goqu v9
- **Environment**: Godotenv
- **Architecture**: Clean Architecture
