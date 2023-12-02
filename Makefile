ENV := release
NAME := api
SRC_APP := ./cmd/$(NAME)/$(NAME).go
BUILD_BIN := ./bin/$(NAME)
BUILD_CMD := go build -o $(BUILD_BIN) $(SRC_APP)

reload: air
	@GIN_MODE=$(ENV) air --log.main_only=true --build.cmd "$(BUILD_CMD)" --build.bin "$(BUILD_BIN)"

run: build
	@GIN_MODE=$(ENV) $(BUILD_BIN)

build:
	@$(BUILD_CMD)

migrate:
	@go run ./cmd/migrate/migrate.go

seed:
	@go run ./cmd/seed/seed.go

reset: migrate seed

test:
	@go test ./... -v

testfail:
	@go test ./... -v | fgrep FAIL || echo "No test failed"

cover:
	@go test ./... --cover --coverprofile=cover.out
	@go tool cover -html=cover.out
	@rm cover.out

coverall:
	@go test ./... --cover --coverprofile=cover.out >> /dev/null
	@go tool cover --func cover.out | grep total
	@rm cover.out

air:
	@command -v air > /dev/null || go install github.com/cosmtrek/air@latest

clean:
	@rm $(BUILD_BIN)
