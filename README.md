# 🎓 Students Management API

![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-336791?style=flat&logo=postgresql)

A robust RESTful API built with **Golang 1.25** to manage student records. This project demonstrates backend engineering best practices, including **clean architecture**, **structured logging**, **input validation**, and **graceful shutdown**.

> **Note:** This is a backend-only application. API interaction is handled via **Postman** or **cURL**.

---

## 🚀 Tech Stack & Features

* **Language:** Golang (1.25.5)
* **Database:** PostgreSQL (Driver: `lib/pq`)
* **Routing:** Standard Library `net/http` (ServeMux)
* **Validation:** `go-playground/validator` for strict payload verification.
* **Logging:** `log/slog` (Structured JSON Logging).
* **Config:** YAML-based configuration management via `ilyakaznacheev/cleanenv`.

---

## 📂 Project Structure

The project follows a modular layout to ensure scalability and maintainability:

```bash
├── cmd/
│   └── student-api/    # Main entry point (main.go)
├── internal/
│   ├── config/         # Configuration loader & Types
│   │   └── types/      # Data structures (Student models)
│   ├── http/
│   │   └── handlers/
│   │       └── student/ # HTTP Handlers (CRUD logic)
│   └── storage/
│       ├── postgresql/  # Database Access Layer (PostgreSQL implementation)
│       └── storage.go   # Storage Interface definition
├── config/
│   └── local.yaml      # Environment configuration
├── go.mod              # Dependency management
└── README.md           # Documentation
🛠️ Getting StartedFollow these steps to set up the project locally.1. PrerequisitesGo installed (version 1.25 or higher).PostgreSQL installed and running.2. Database SetupEnsure your PostgreSQL service is running. The application automatically creates the students table if it doesn't exist.Update the credentials in config/local.yaml to match your local database:YAML# config/local.yaml
storage_path: "host=localhost port=5432 user=postgres password=YOUR_PASSWORD dbname=students-api sslmode=disable"
http_server:
  address: "localhost:8082"
3. Run the ApplicationYou must provide the path to the configuration file using the -config flag.Bash# Clone the repository
git clone [https://github.com/ZeeshanSaleem-official/student-api.git](https://github.com/ZeeshanSaleem-official/student-api.git)
cd student-api

# Install dependencies
go mod tidy

# Run the server
go run cmd/student-api/main.go -config=./config/local.yaml
You should see a log message confirming the server has started at localhost:8082.🔌 API EndpointsBase URL: http://localhost:8082MethodEndpointDescriptionPayload (JSON)POST/api/studentsCreate a new student{"name": "Zeeshan", "email": "test@example.com", "age": 21}GET/api/students/{id}Retrieve student by IDN/AGET/api/studentsList all studentsN/APUT/api/students/{id}Update student details{"name": "Zeeshan Updated", "age": 22}DELETE/api/students/{id}Remove a studentN/A🧪 Testing with PostmanSince this system focuses on backend performance, use Postman to interact with the endpoints:Open Postman.Create a request (e.g., POST http://localhost:8082/api/students).Go to Body → Raw → JSON.Paste the payload:JSON{
  "name": "Candidate One",
  "email": "candidate@cern.ch",
  "age": 22
}
Hit Send and observe the structured JSON response.💡 Design DecisionsStandard Library Routing (net/http):I utilized Go's http.NewServeMux to handle routing logic without relying on heavy external frameworks like Gin. This demonstrates a deep understanding of the language's core capabilities and minimizes runtime overhead.Dependency Injection:The student handler receives a storage.Storage interface rather than a concrete struct. This decouples the business logic from the database, adhering to the Dependency Inversion Principle.Graceful Shutdown:The application listens for OS signals (SIGINT, SIGTERM) to ensure that the server shuts down cleanly, closing active connections and preventing data corruption.📜 LicenseDistributed under the MIT License.