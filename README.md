# 📚 Book API — Go + Connect RPC + GORM + PostgreSQL

This project implements a simple Book Management API built with Go, powered by Connect RPC (from connectrpc.com
), and backed by GORM with a PostgreSQL database.

It provides full CRUD operations — Create, Read (List/Get), Update, and Delete — for managing book data.

# 🚀 Features

✅ Built with Go (Golang)
✅ Uses Connect RPC — a modern, high-performance RPC framework
✅ GORM ORM for database abstraction
✅ PostgreSQL as the relational database
✅ Clean architecture and modular folder structure
✅ Auto migration on startup
✅ Simple curl-based testing interface

# 🧩 Tech Stack
Component	Technology Used
Language	Go 1.24+
RPC Framework	Connect RPC
ORM	GORM
Database	PostgreSQL
Code Generation	Buf
Dependency Management	Go Modules

# 📁 Project Structure
Book_api_using_Golang/
├── api/                     # Proto & generated ConnectRPC files
│   └── gen/book/v1          # Generated gRPC/Connect code
├── internal/
│   ├── db/                  # Database connection and migration
│   ├── models/              # GORM models
│   └── service/             # ConnectRPC service logic (CRUD)
├── main.go                  # Application entrypoint
├── go.mod / go.sum          # Dependencies
└── README.md                # You're here!

# ⚙️ Setup Instructions
1️⃣ Clone the repository
git clone https://github.com/SajibKC/bookapi.git
cd bookapi

2️⃣ Set up PostgreSQL

Create a database named bookapi:

psql -U postgres
CREATE DATABASE bookapi;


Update your database credentials in internal/db/database.go:

dsn := "host=localhost user=postgres password=yourpassword dbname=bookapi port=5432 sslmode=disable"

3️⃣ Install dependencies
go mod tidy

4️⃣ Run the server
go run main.go


The API will start at:

http://localhost:8080

# 🧠 API Endpoints (Connect RPC style)
Operation	RPC Method	Description
Create	book.v1.BookService/CreateBook	Create a new book
List	book.v1.BookService/ListBooks	List all books
Get	book.v1.BookService/GetBook	Get details of a specific book
Update	book.v1.BookService/UpdateBook	Update an existing book
Delete	book.v1.BookService/DeleteBook	Delete a book by ID

# 📬 Example Requests (using curl)
# ➕ Create a book
curl -X POST -H "Content-Type: application/json" \\
  -d '{"title":"AOT","author":"Hajime Isayama","price":500}' \\
  http://localhost:8080/book.v1.BookService/CreateBook

# 📖 List all books
curl -X POST \
  -H "Content-Type: application/json" \\
  -d '{}' \\
  http://localhost:8080/book.v1.BookService/ListBooks

# 🔍 Get a book by ID
curl -X POST http://localhost:8080/book.v1.BookService/GetBook \\
  -H "Content-Type: application/json" \\
  -d '{"id":"1"}'

# 🖊️ Update a book
curl -X POST http://localhost:8080/book.v1.BookService/UpdateBook \\
  -H "Content-Type: application/json" \\
  -d '{"id":"1","title":"AOT Updated","author":"Hajime Isayama","price":450}'

# ❌ Delete a book
curl -X POST http://localhost:8080/book.v1.BookService/DeleteBook \\
  -H "Content-Type: application/json" \\
  -d '{"id":"1"}'

# 🧰 Developer Notes

The project uses Connect RPC instead of REST or gRPC — it supports both gRPC and HTTP/1.1 seamlessly.

GORM AutoMigrate is called on startup to ensure tables exist.

The API is fully typed and compatible with modern frontends (e.g., using Connect Web clients).

You can later containerize this with Docker and deploy to Kubernetes or Helm.
