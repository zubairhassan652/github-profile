# My Github Profile For Code Review

Purpose of this repo is to show of standard coding styles or my own coding patterns in `go`.

### How to run this project on local machine

Using standard way:

>`go run main.go`


Using docker file:

>`docker build -t go-gorilla-mux .`

>`docker run -p 8080:8080 go-gorilla-mux`

Start postgres-15 container:

>`docker run --name postgres-v-15-container -e POSTGRES_PASSWORD=your-postgres-password -p 5432:5432 -d postgres:15`

>`docker exec -it your-postgres-container psql -U postgres`

### How to run migrations

>`migrate -path ./users/migrations -database "postgresql://postgres:your-postgres-password@localhost:5432/postgres?sslmode=disable" up`


### How to kill port

`sudo lsof -i :8080`

`sudo kill <PID>`