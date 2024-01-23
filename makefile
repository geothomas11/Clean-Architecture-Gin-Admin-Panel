GOCMD:=go
BUILD_DIR:=build 
BINARY_DIR:=$(BUILD_DIR)/bin 


all:test build

run:
	$(GOCMD) run ./cmd/api/main.go

wire:
	cd pkg/di_in && wire