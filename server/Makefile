run: build
	@./bin/sugoiapi

build:
	@go build -o bin/sugoiapi cmd/api/main.go

seed:
	@go run cmd/seed/main.go

drop:
	@go run cmd/drop/main.go

docker-build:
	docker-compose build sugoiapi

docker-restart:
	docker-compose up -d

docker-run:
	docker-compose up -d --build