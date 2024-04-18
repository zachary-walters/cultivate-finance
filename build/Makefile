BACKEND_BINARY=backendApp
WASM_BINARY=wasmBin

test: fmt
	(cd ../backend; go test ./... -v)

## up: starts all containers in the background without forcing build
up: 
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: fmt build_api
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down: 
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Docker compose stopped!"

scan: 
	cd ../ && sonar-scanner \
  -Dsonar.projectKey=rothvtrad \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://localhost:9999 \
  -Dsonar.login=sqp_2aec7aedfab84ecf9b0946753cb25feedcb86835

## build_api: builds the backend api as a linux executable
build_api: 
	@echo "Building api backend binary..."
	cd ../backend && env GOOS=linux CGO_ENABLED=0 go build -o bin/${BACKEND_BINARY} ./cmd/api
	@echo "Done building backend binary!"

serve:
	cd ../backend/cmd/api && fresh -c runner.conf

build_wasm:
	@echo "Building the wasm binary..."
	cd ../backend && env GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -o bin/${WASM_BINARY} ./cmd/wasm
	@echo "Done build wasm binary"

fmt: 
	gofmt -s -w ../.