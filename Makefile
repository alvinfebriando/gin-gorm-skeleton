SRC_REST := ./cmd/rest/rest.go
BIN_REST := ./bin/rest
BUILD_REST_CMD := go build -o $(BIN_REST) $(SRC_REST)

BIN_MIGRATE := ./bin/migrate
BIN_SEED := ./bin/seed

SKELETON := github.com/alvinfebriando/gin-gorm-skeleton

reload: air
	@GIN_MODE=release air --log.main_only=true --build.cmd "$(BUILD_REST_CMD)" --build.bin "$(BIN_REST)"

run: build
	@GIN_MODE=release $(BIN_REST)

build:
	@$(BUILD_REST_CMD)

migrate:
	@go run ./cmd/migrate/migrate.go

migration:
	@$(BIN_MIGRATE)
	@$(BIN_SEED)

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
	@rm $(BIN_REST)

rename:
	@find . -type f -name "*.mod" -exec sed -i'' -e 's,$(SKELETON),$(MODULE),g' {} +
	@find . -type f -name "*.go" -exec sed -i'' -e 's,$(SKELETON),$(MODULE),g' {} +

gen:
	@protoc -I=./proto --go_out=. --go-grpc_out=. ./proto/*.proto

up:
	@docker compose up

down:
	@docker compose down

dev:
	@docker compose up rest_dev

refresh:
	@docker compose up migration