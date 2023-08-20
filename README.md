# My Github Profile For Code Review

Purpose of this repo is to show of standard coding styles or my own coding patterns in `go`.

### Manage DB

Start postgres-15 container:

>`docker run --name postgres-v-15-container -e POSTGRES_PASSWORD=your-postgres-password -p 5432:5432 -d postgres:15`


### Manage migrations

>`migrate create -ext sql -dir ./<users>/migrations -seq <migration_name>`

>`migrate -path ./<users>/migrations -database "postgresql://postgres:your-postgres-password@localhost:5432/postgres?sslmode=disable" up`


### How to run this project on local machine

Using standard way:

>`go run main.go`


Using docker file:

>`docker build -t go-vue .`

>`docker run -p 8080:8080 go-vue`
