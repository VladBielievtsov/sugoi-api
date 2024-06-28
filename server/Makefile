run: build
	@./bin/sugoiapi

build:
	@go build -o bin/sugoiapi cmd/api/main.go

seed:
	@go run cmd/seed/main.go

drop:
	@go run cmd/drop/main.go