GO := go
APP_NAME := converter
SRC_DIR := .
BUILD_DIR := ./build
API_KEY ?= $(CONVERTER_COIN_MARKET_CAP_API_KEY)
AMOUNT ?= 10
CURRENCY_FROM ?= USD
CURRENCY_TO ?= BTC

.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)/cmd/$(APP_NAME)/main.go

.PHONY: test
test:
	$(GO) test -race -v ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: lint
lint:
	golangci-lint run -v

.PHONY: run
run:
	CONVERTER_COIN_MARKET_CAP_API_KEY=$(API_KEY) $(GO) run $(SRC_DIR)/cmd/$(APP_NAME)/main.go $(AMOUNT) $(CURRENCY_FROM) $(CURRENCY_TO)

.PHONY: all
all: build run

.PHONY: test-all
test-all: test

.PHONY: help
help:
	@echo 'Available commands:'
	@echo '  make                - build the project'
	@echo '  make test           - run tests'
	@echo '  make clean          - clean up built files'
	@echo '  make lint           - lint the code'
	@echo '  make run            - build and run the project'
	@echo '  make all            - build and run the project'
	@echo '  make test-all       - build and run tests'
	@echo '  make ci             - build, run tests, and lint the code'
	@echo '  make help           - display this help message'
