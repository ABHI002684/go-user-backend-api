# Go Backend Development Task – User API

This project is a RESTful backend API built using **Go** to manage users with their **name** and **date of birth (DOB)**.  
The API dynamically calculates and returns a user’s **age** when fetching user details.

This project was developed as part of the **Go Backend Development Task** for Ainyx Solutions.

---

## 🚀 Tech Stack

- **Go**
- **Fiber** – Web framework
- **PostgreSQL** – Database
- **SQLC** – Type-safe SQL code generation
- **Uber Zap** – Structured logging
- **go-playground/validator** – Input validation
- **godotenv** – Environment variable management

---

## 📁 Project Structure

/cmd/server/main.go 

/config/ 

/db/migrations/ 

/db/sqlc/<generated> 

/internal/ 

├── handler/ 

├── repository/ 

├── service/ 

├── routes/

├── middleware/ 

├── models/ 

└── logger/ 

.env 

go.mod 

go.sum 

README.md 

Reasoning.md 

---

## ✨ Features

- CRUD APIs for users
- Stores DOB in database
- Dynamically calculates age using Go’s `time` package
- Clean layered architecture (handler → service → repository)
- Type-safe database access using SQLC
- Environment-based configuration

---

## ⚙️ Setup & Run Instructions

### 1️⃣ Prerequisites

- Go (>= 1.21)
- PostgreSQL
- sqlc installed
- Postman (optional)

---

### 2️⃣ Clone the Repository

```bash
git clone https://github.com/ABHI002684/go-user-backend-api 
cd go-user-backend-api 

```

### 3️⃣ Environment Configuration

Create a .env file in the project root: 

DB_URL=postgres://postgres:<password>@localhost:5432/userdb?sslmode=disable 

PORT=3000 

### 4️⃣ Database Setup 
CREATE DATABASE userdb;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

### 5️⃣ Generate SQLC Code
sqlc generate

### 6️⃣ Install Dependencies
go mod tidy

### 7️⃣ Run the Server
go run cmd/server/main.go 

Server runs at:
http://127.0.0.1:3000

### 🔗 API Endpoints

### Create User  
POST /users 

{
  "name": "Alice",
  "dob": "1990-05-10"
}

### Get User
GET /users/:id

### List Users
GET /users 

### Update User
PUT /users/:id

### Delete User
DELETE /users/:id

### 🧪 Testing
All APIs were tested using Postman to verify correct CRUD behavior and dynamic age calculation.

### 📄 Documentation
Reasoning.md – explains design decisions and learning approach



