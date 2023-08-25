# Go Web Application

Purpose of this repo is to show of standard coding styles of web application or my own coding patterns in `go`.

### Manage DB

Start mongo db container 

>`docker run -d --name my-mongo-container -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin mongo`

Start postgres db container:

>`docker run --name postgres-v-15-container -e POSTGRES_PASSWORD=your-postgres-password -p 5432:5432 -d postgres:15`


### Manage migrations

>`migrate create -ext sql -dir ./internal/users/migrations -seq <migration_name>`

>`migrate -path ./internal/users/migrations -database "postgresql://postgres:your-postgres-password@localhost:5432/postgres?sslmode=disable" up`


### How to run this project on local machine

Using standard way:

>`go run main.go`


Using docker file:

>`docker build -t go-vue .`

>`docker run -p 8080:8080 go-vue`
