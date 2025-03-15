# makefile

# program name
PROGRAM_NAME := zenchat

# executable file
EXECUTABLE := zenchat

# project work dir
SRC_DIR := $(CURRENT_DIR)

# program arguments
CONFIG_FILE := ./config.yaml

# Go command
GO := go

# build
.PHONY: build

# db upgrade
.PHONY: db
db:
	$(GO) run upgrade/main.go -c=./config.yaml

# run app
.PHONY: run
run:
	$(GO) run app/main.go -c=./config.yaml

# build app linux
.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags "-w -s" -o ./bin/app/$(EXECUTABLE) app/main.go
	upx -9 ./bin/app/$(EXECUTABLE)

# dev
.PHONY: build
build:
	$(GO) build -ldflags "-w -s" -o ./bin/dev/$(EXECUTABLE) app/main.go
	upx -9 ./bin/dev/$(EXECUTABLE)

# clean
.PHONY: clean
clean:
	$(GO) clean
	rm -f ./bin/dev/$(EXECUTABLE)
	rm -f ./bin/app/$(EXECUTABLE)



