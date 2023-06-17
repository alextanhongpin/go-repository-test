include .env
export


all: run


include Makefile.*.mk


up:
	@docker-compose up -d


down:
	@docker-compose down


run:
	@go run main.go


tidy:
	@go mod tidy


test:
	@gotest -v -failfast -cover -coverprofile=cover.out ./...
	@go tool cover -html=cover.out


prune:
	@docker system prune --volumes --force
