GO=go
BIN_DIR = $(CURDIR)/bin

tidy:
	$(GO) mod tidy -v

build:
	$(GO) build -v -o $(BIN_DIR)/service

test:
	$(GO) test -cover -race -v ./...

dev:
	docker-compose -f support/docker-compose.dev.yml down -v
	#cd support && docker-compose rm -vsf
	docker-compose -f support/docker-compose.dev.yml up --build