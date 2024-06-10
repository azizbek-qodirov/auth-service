CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

run:
	go run cmd/main.go
init:
	go mod init
	go mod tidy 
	go mod vendor
	
migrate_up:
	migrate -path migrations -database postgres://ahrorbek:ahrorbek@localhost:5432/ideal_cleaning_db -verbose up

migrate_down:
	migrate -path migrations -database postgres://ahrorbek:ahrorbek@localhost:5432/ideal_cleaning_db -verbose down

migrate_force:
	migrate -path migrations -database postgres://ahrorbek:ahrorbek@localhost:5432/ideal_cleaning_db -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq ideal_cleaning_db_table

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

lint: ## Run golangci-lint with printing to stdout
	golangci-lint -c .golangci.yaml run --build-tags "musl" ./...

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o api/docs force 1