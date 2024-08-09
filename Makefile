build:
	@go build -o bin/app cmd/app/main.go

run: build
	@bin/app

dev:
	@air

templ:
	@templ generate
	
db_reset:
	@psql -h localhost -p 5432 -U postgres -c "DROP DATABASE IF EXISTS gophatt"
	@psql -h localhost -p 5432 -U postgres -c "CREATE DATABASE gophatt"

generate_ent:
	@go generate ./ent

generate_schema:
	@go run -mod=mod entgo.io/ent/cmd/ent new ${name}
