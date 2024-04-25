BACKEND_BINARY=backendApp
WASM_BINARY=wasmBin
TINYGO_BINARY=tinygoBin.wasm
WEBAPP_BINARY=webapp

GOROOT = $(shell go env GOROOT)
TINYGOROOT = $(shell tinygo env TINYGOROOT)

test: fmt
	(cd ../backend; go test ./... -v)

## up: starts all containers in the background without forcing build
up: 
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: fmt build_api build_wasm build_webapp
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down: 
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Docker compose stopped!"

serve:
	cd ../backend/cmd/api && fresh -c runner.conf

webapp: fmt build_tinygo build_webapp 
	@echo "rebuilding and starting webapp"
	cp ${TINYGOROOT}/targets/wasm_exec.js ../webapp/cmd/assets
	cp ../backend/bin/${TINYGO_BINARY} ../webapp/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' ../webapp/cmd/assets/wasm_exec.js
	docker-compose up -d --build webapp

api: fmt build_api
	@echo "rebuilding and starting api"
	docker-compose up -d --build  api 

build_api: 
	@echo "Building api backend binary..."
	cd ../backend && env GOOS=linux CGO_ENABLED=0 go build -o bin/${BACKEND_BINARY} ./cmd/api
	du -sh ../backend/bin/${BACKEND_BINARY}
	@echo "Done building backend binary!"

build_wasm:
	@echo "Building the wasm binary..."
	cd ../backend && env GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -o bin/${WASM_BINARY} ./cmd/wasm
	du -sh ../backend/bin/${WASM_BINARY}
	@echo "Done building wasm binary"

build_tinygo:
	@echo "Building tinygo wasm binary..."
	cd ../backend && tinygo build -o bin/${TINYGO_BINARY} -target wasm -opt=z -no-debug -size full -scheduler=asyncify ./cmd/wasm
	du -sh ../backend/bin/${TINYGO_BINARY}
	@echo "Done building tinygo binary"

build_webapp: build_tinygo
	@echo "Building the HTMX webapp..."
	cd ../webapp && env GOOS=linux CGO_ENABLED=0 go build -o bin/${WEBAPP_BINARY} ./cmd/main.go
	du -sh ../webapp/bin/${WEBAPP_BINARY}
	@echo "Done building webapp binary"

fmt: 
	gofmt -s -w ../.