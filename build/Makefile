include $(PWD)/.env

test: fmt
	(cd $(401K_CALCULATOR_DIR); go test ./... -v -coverpkg=./...)

coverage:
	(cd $(401K_CALCULATOR_DIR); go test -v -coverprofile bin/cover.out ./... -coverpkg=./...; go tool cover -html bin/cover.out -o bin/cover.html; open bin/cover.html)

#### up: starts all containers in the background without forcing build
up: 
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker images started!"

#### up: starts all containers in after building
up_build: fmt build_401k_api build_webapp build_api_gateway wiki
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
	cd $(401K_CALCULATOR_DIR)/cmd/api && fresh -c runner.conf

webapp: fmt build_401k_tinygo build_webapp 
	@echo "rebuilding and starting webapp"
	cp ${TINYGOROOT}/targets/wasm_exec.js $(WEBAPP_DIR)/cmd/assets
	cp $(401K_CALCULATOR_DIR)/bin/${TINYGO_BINARY} $(WEBAPP_DIR)/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' $(WEBAPP_DIR)/cmd/assets/wasm_exec.js
	docker-compose up -d --build webapp

api: fmt build_401k_api
	@echo "rebuilding and starting api"
	docker-compose up -d --build  401k-calculator-api 

wiki: 
	@echo "building and starting wiki"
	docker-compose up -d --build wiki
	cd ${WIKI_DIR}/scripts && ./backup.sh restore

build_401k_api: 
	@echo "Building api 401k_calculator binary..."
	cd $(401K_CALCULATOR_DIR) && env GOOS=linux CGO_ENABLED=0 go build -o bin/${401K_CALCULATOR_BINARY} ./cmd/api
	du -sh $(401K_CALCULATOR_DIR)/bin/${401K_CALCULATOR_BINARY}
	@echo "Done building 401k_calculator binary!"

build_401k_tinygo:
	@echo "Building tinygo wasm binary..."
	cd $(401K_CALCULATOR_DIR) && tinygo build -o bin/${TINYGO_BINARY} -target wasm -opt=z -no-debug -size full -scheduler=asyncify -panic=trap ./cmd/wasm 
	wasm-opt -s -Oz $(401K_CALCULATOR_DIR)/bin/${TINYGO_BINARY} -o $(401K_CALCULATOR_DIR)/bin/${TINYGO_BINARY} && chmod +x $(401K_CALCULATOR_DIR)/bin/${TINYGO_BINARY}
	du -sh $(401K_CALCULATOR_DIR)/bin/${TINYGO_BINARY}
	@echo "Done building tinygo binary"

build_webapp: build_401k_tinygo
	@echo "Building the webapp..."
	cp ${TINYGOROOT}/targets/wasm_exec.js $(WEBAPP_DIR)/cmd/assets
	cp $(401K_CALCULATOR_DIR)/bin/${TINYGO_BINARY} $(WEBAPP_DIR)/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' $(WEBAPP_DIR)/cmd/assets/wasm_exec.js
	cd $(WEBAPP_DIR) && env GOOS=linux CGO_ENABLED=0 go build -o bin/${WEBAPP_BINARY} ./cmd/main.go
	du -sh $(WEBAPP_DIR)/bin/${WEBAPP_BINARY}
	@echo "Done building webapp binary"

build_api_gateway: 
	@echo "Building the api gateway..."
	cd ../api_gateway && env GOOS=linux CGO_ENABLED=0 go build -o bin/${API_GATEWAY_BINARY} ./cmd
	du -sh ../api_gateway/bin/${API_GATEWAY_BINARY}
	@echo "Done building api gateway binary!"

fmt: 
	gofmt -s -w ../.

### DOCKER BUILD/PUSH
docker_build: docker_webapp docker_api_gateway docker_401k_api

docker_webapp: build_webapp
	@echo building docker webapp container...
	cd $(WEBAPP_DIR) && docker build -t cultivate-finance-webapp-service -f webapp.dockerfile .
	docker tag cultivate-finance-webapp-service $(DOCKERHUB_REPO)/cultivate-finance-webapp-service

docker_api_gateway: build_api_gateway
	@echo building docker api gateway container...
	cd $(API_GATEWAY_DIR) && docker build -t cultivate-finance-api-gateway-service -f api_gateway.dockerfile .
	docker tag cultivate-finance-api-gateway-service $(DOCKERHUB_REPO)/cultivate-finance-api-gateway-service

docker_401k_api: build_401k_api
	@echo building docker 401k calculator container...
	cd $(401K_CALCULATOR_DIR) && docker build -t cultivate-finance-401k-calculator-service -f 401k_calculator.dockerfile .
	docker tag cultivate-finance-401k-calculator-service $(DOCKERHUB_REPO)/cultivate-finance-401k-calculator-service

docker_push: docker_build
	docker push $(DOCKERHUB_REPO)/cultivate-finance-webapp-service
	docker push $(DOCKERHUB_REPO)/cultivate-finance-api-gateway-service
	docker push $(DOCKERHUB_REPO)/cultivate-finance-401k-calculator-service
	
### K8S
k_deploy: k_delete k_deploy_nats k_deploy_api k_deploy_webapp k_deploy_401k

k_delete: k_delete_kubectl

k_delete_kubectl:
	kubectl delete pods,services,deployment --all

k_deploy_nats:
	kubectl apply -f k8s/nats.yml

k_deploy_webapp:
	kubectl apply -f k8s/webapp.yml

k_deploy_401k:
	kubectl apply -f k8s/401k_calculator.yml

k_deploy_api:
		kubectl apply -f k8s/api_gateway.yml

### Versioning
backup_wiki:
	@echo "backing up the wiki..."
	cd ${WIKI_DIR}/scripts && ./backup.sh backup

restore_wiki:
	@echo "restoring the wiki"
	cd ${WIKI_DIR}/scripts && ./backup.sh restore