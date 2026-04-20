run: go run ./cmd/api

update update commend: swag init -g cmd/api/main.go --output ./docs
Link: http://localhost:4040/swagger/index.html

api file:
=========
carousel/
├── handler.go
├── service.go
├── repository.go
└── model.go

🟦 handler.go → HTTP Layer

    .Only deals with HTTP
    .Read request (JSON, params)
    .Call service
    .Send response

    ❌ No DB
    ❌ No business logic


🟩 service.go → Business Logic

   The brain of your app
    .Validation
    .Rules
    .Time setting
    .Transactions
    .Orchestration

🟨 repository.go → Data Layer

    .Only DB work
    .Mongo queries
    .SQL queries
    .Redis access


🟦 model.go
    .For table row
