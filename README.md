# Eneba Product Search Application

Product search with fuzzy matching. React, Go, and PostgreSQL.

**Live:** https://eneba-hw-ten.vercel.app  
*Note: Backend is hosted on free tier services and may take up to 1 minute to start after inactivity.*


## Stack

**Frontend:** React 19, Vite, Tailwind CSS  
**Backend:** Go 1.24, Gin, GORM, PostgreSQL

## Prerequisites

- Go 1.24+
- PostgreSQL 14+
- Node.js 18+ or Bun 1.0+

## Setup

**1. Database**
```bash
psql -U postgres -c "CREATE DATABASE eneba;"
```

**2. Backend**
```bash
cd backend
cp .env.example .env
# Edit .env with your credentials
go mod download
go run cmd/api/main.go
```

**3. Frontend**
```bash
cd frontend
npm install && npm run dev
```

## Environment Variables

**backend/.env**
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=eneba
DB_SSLMODE=disable
PORT=8080
FRONTEND_ORIGIN=http://localhost:5173
```

**frontend/.env**
```env
VITE_API_URL=http://localhost:8080
```
