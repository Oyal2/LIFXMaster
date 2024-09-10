.PHONY: all proto proto_ts clean package-mac package-linux package-win package build-go-win build-go-mac build-go-linux build-go

# Project structure
ROOT_DIR := $(shell pwd)
ELECTRON_DIR := $(ROOT_DIR)/web
GO_DIR := $(ROOT_DIR)/backend
DIST_DIR := $(ROOT_DIR)/dist

# Proto file locations
BACKEND_PROTO_DIR := $(GO_DIR)/internal
FRONTEND_PROTO_DIR := $(ELECTRON_DIR)/src/proto

# GRPC CMD location
BACKEND_PROTO_DIR := $(GO_DIR)/cmd/grpc

APP_NAME := lifxmaster
APP_VERSION := 1.0.0

# OS-specific settings
ifeq ($(OS),win)
    GO_BINARY := grpc-server.exe
else
    GO_BINARY := grpc-server
endif


proto: clean proto-ts proto-go

proto-go: clean
	protoc --go_out=$(BACKEND_PROTO_DIR) --go_opt=paths=source_relative \
	       --go-grpc_out=$(BACKEND_PROTO_DIR) --go-grpc_opt=paths=source_relative \
	       proto/*.proto

proto-ts: 
	npx protoc \
		--ts_out ./$(FRONTEND_PROTO_DIR) \
		 proto/*.proto

# Build Go gRPC server
build-go-win:
	cd $(BACKEND_PROTO_DIR) && GOOS=windows go build -o $(DIST_DIR)/grpc-server.exe

build-go-mac:
	cd $(BACKEND_PROTO_DIR) && GOOS=darwin go build -o $(DIST_DIR)/grpc-server

build-go-linux:
	cd $(BACKEND_PROTO_DIR) && GOOS=linux go build -o $(DIST_DIR)/grpc-server

build-go: build-go-$(OS)

# Build Electron app
build-electron:
	cd $(ELECTRON_DIR) && npm install && npm run build

# Package the app
package-mac: clean-dist build-go-mac
	cd $(ELECTRON_DIR) && npm run build:mac && mv dist/$(APP_NAME)-* $(DIST_DIR)

package-linux: clean-dist build-go-linux
	cd $(ELECTRON_DIR) && npm run build:linux && mv dist/$(APP_NAME)-* $(DIST_DIR)

package-win: clean-dist build-go-win
	cd $(ELECTRON_DIR) && npm run build:win && mv dist/$(APP_NAME)-* $(DIST_DIR)

package: package-$(OS)

# Clean build artifacts
clean-dist:
	rm -rf $(DIST_DIR)
	mkdir -p $(DIST_DIR)

clean:
	find $(BACKEND_PROTO_DIR)/proto -name '*.pb.go' -delete
	find $(FRONTEND_PROTO_DIR) -type f -delete