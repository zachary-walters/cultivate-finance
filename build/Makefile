401K_CALCULATOR_BINARY=401kCalculatorApp
WASM_BINARY=wasmBin
TINYGO_BINARY=tinygoBin.wasm
WEBAPP_BINARY=webapp
API_GATEWAY_BINARY=apiGateway

GOROOT = $(shell go env GOROOT)
TINYGOROOT = $(shell tinygo env TINYGOROOT)

test: fmt
	(cd ../401k_calculator; go test ./... -v -coverpkg=./...)

coverage:
	(cd ../401k_calculator; go test -v -coverprofile bin/cover.out ./... -coverpkg=./...; go tool cover -html bin/cover.out -o bin/cover.html; open bin/cover.html)

## up: starts all containers in the background without forcing build
up: 
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: fmt build_api build_wasm build_webapp build_api_gateway
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down: 
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Docker compose stopped!"

serve_api:
	cd ../401k_calculator/cmd/api && fresh -c runner.conf

webapp: fmt build_tinygo build_webapp 
	@echo "rebuilding and starting webapp"
	cp ${TINYGOROOT}/targets/wasm_exec.js ../webapp/cmd/assets
	cp ../401k_calculator/bin/${TINYGO_BINARY} ../webapp/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' ../webapp/cmd/assets/wasm_exec.js
	docker-compose up -d --build webapp

api: fmt build_api
	@echo "rebuilding and starting api"
	docker-compose up -d --build  401k-calculator-api 

build_api: 
	@echo "Building api 401k_calculator binary..."
	cd ../401k_calculator && env GOOS=linux CGO_ENABLED=0 go build -o bin/${401K_CALCULATOR_BINARY} ./cmd/api
	du -sh ../401k_calculator/bin/${401K_CALCULATOR_BINARY}
	@echo "Done building 401k_calculator binary!"

build_wasm:
	@echo "Building the wasm binary..."
	cd ../401k_calculator && env GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -o bin/${WASM_BINARY} ./cmd/wasm
	du -sh ../401k_calculator/bin/${WASM_BINARY}
	@echo "Done building wasm binary"

build_tinygo:
	@echo "Building tinygo wasm binary..."
	cd ../401k_calculator && tinygo build -o bin/${TINYGO_BINARY} -target wasm -opt=z -no-debug -size full -scheduler=asyncify -panic=trap ./cmd/wasm 
	wasm-opt -s -Oz ../401k_calculator/bin/${TINYGO_BINARY} -o ../401k_calculator/bin/${TINYGO_BINARY} && chmod +x ../401k_calculator/bin/${TINYGO_BINARY}
	du -sh ../401k_calculator/bin/${TINYGO_BINARY}
	@echo "Done building tinygo binary"

build_webapp: build_tinygo
	@echo "Building the webapp..."
	cp ${TINYGOROOT}/targets/wasm_exec.js ../webapp/cmd/assets
	cp ../401k_calculator/bin/${TINYGO_BINARY} ../webapp/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' ../webapp/cmd/assets/wasm_exec.js
	cd ../webapp && env GOOS=linux CGO_ENABLED=0 go build -o bin/${WEBAPP_BINARY} ./cmd/main.go
	du -sh ../webapp/bin/${WEBAPP_BINARY}
	@echo "Done building webapp binary"

build_api_gateway: 
	@echo "Building the api gateway..."
	cd ../api_gateway && env GOOS=linux CGO_ENABLED=0 go build -o bin/${API_GATEWAY_BINARY} ./cmd
	du -sh ../api_gateway/bin/${API_GATEWAY_BINARY}
	@echo "Done building api gateway binary!"

fmt: 
	gofmt -s -w ../.