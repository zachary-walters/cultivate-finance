include $(PWD)/.env

#### Testing
test_401k: fmt
	(cd $(401K_CALCULATOR_DIR); go test ./... -v -coverpkg=./... -race)

test_debt_snowball: fmt
	(cd $(DEBT_SNOWBALL_CALCULATOR_DIR); go test ./... -v -coverpkg=./... -race)

test: fmt test_401k test_debt_snowball
	
coverage_debt_snowball: fmt 
		(cd $(DEBT_SNOWBALL_CALCULATOR_DIR); go test -v -coverprofile bin/cover.out ./... -coverpkg=./...; go tool cover -html bin/cover.out -o bin/cover.html; open bin/cover.html)

coverage:
	(cd $(401K_CALCULATOR_DIR); go test -v -coverprofile bin/cover.out ./... -coverpkg=./...; go tool cover -html bin/cover.out -o bin/cover.html; open bin/cover.html)

#### up: starts all containers in the background without forcing build
up: 
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker images started!"

#### up: starts all containers in after building
up_build: fmt build_401k_api build_webapp build_api_gateway build_debt_snowball_api build_link_service
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

nats:
	@echo "recreating nats service"
	docker-compose up -d --build nats

webapp: fmt build_webapp 
	@echo "rebuilding and starting webapp"
	cp ${TINYGOROOT}/targets/wasm_exec.js $(WEBAPP_DIR)/cmd/assets
	cp $(401K_CALCULATOR_DIR)/bin/${TINYGO_401K_BINARY} $(WEBAPP_DIR)/cmd/assets
	cp $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${TINYGO_DEBT_SNOWBALL_BINARY} $(WEBAPP_DIR)/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' $(WEBAPP_DIR)/cmd/assets/wasm_exec.js
	docker-compose up -d --build webapp

401k_api: fmt build_401k_api
	@echo "rebuilding and starting 401k api"
	docker-compose up -d --build  401k-calculator-api 

debt_snowball_api: fmt build_debt_snowball_api
	@echo "rebuilding and starting debt_snowball api"
	docker-compose up -d --build debt-snowball-api

api_gateway: build_api_gateway
	@echo "rebuilding and starting api_gateway"
	docker-compose up -d --build api-gateway	

link_service: fmt build_link_service
	@echo "rebuilding and starting link_service"
	docker-compose up -d --build link-service-api

build_401k_api: 
	@echo "Building api 401k_calculator binary..."
	cd $(401K_CALCULATOR_DIR) && env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o bin/${401K_CALCULATOR_BINARY} ./cmd/api
	du -sh $(401K_CALCULATOR_DIR)/bin/${401K_CALCULATOR_BINARY}
	@echo "Done building 401k_calculator binary!"

build_401k_tinygo:
	@echo "Building 401k tinygo wasm binary..."
	mkdir -p $(401K_CALCULATOR_DIR)/bin
	cd $(401K_CALCULATOR_DIR) && tinygo build -o bin/${TINYGO_401K_BINARY} -target wasm -opt=z -no-debug -size full -scheduler=asyncify -panic=trap ./cmd/wasm 
	wasm-opt -s -Oz $(401K_CALCULATOR_DIR)/bin/${TINYGO_401K_BINARY} -o $(401K_CALCULATOR_DIR)/bin/${TINYGO_401K_BINARY} && chmod +x $(401K_CALCULATOR_DIR)/bin/${TINYGO_401K_BINARY}
	du -sh $(401K_CALCULATOR_DIR)/bin/${TINYGO_401K_BINARY}
	@echo "Done building tinygo binary"

build_webapp: build_401k_tinygo build_debt_snowball_tinygo
	@echo "Building the webapp..."
	cp ${TINYGOROOT}/targets/wasm_exec.js $(WEBAPP_DIR)/cmd/assets
	cp $(401K_CALCULATOR_DIR)/bin/${TINYGO_401K_BINARY} $(WEBAPP_DIR)/cmd/assets
	cp $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${TINYGO_DEBT_SNOWBALL_BINARY} $(WEBAPP_DIR)/cmd/assets
	sed -e '/finalizeRef not implemented/ s/^/\/\//' -i '' $(WEBAPP_DIR)/cmd/assets/wasm_exec.js
	cd $(WEBAPP_DIR) && env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o bin/${WEBAPP_BINARY} ./cmd/main.go
	du -sh $(WEBAPP_DIR)/bin/${WEBAPP_BINARY}
	@echo "Done building webapp binary"

build_debt_snowball_api: 
	@echo "Building api debt_snowball binary..."
	cd ${DEBT_SNOWBALL_CALCULATOR_DIR} && env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o bin/${DEBT_SNOWBALL_CALCULATOR_BINARY} ./cmd/api
	du -sh $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${DEBT_SNOWBALL_CALCULATOR_BINARY}
	@echo "Done building debt_snowball binary!"

build_debt_snowball_tinygo:
	@echo "Building Debt Snowball tinygo wasm binary..."
	mkdir -p $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin
	cd $(DEBT_SNOWBALL_CALCULATOR_DIR) && tinygo build -o bin/${TINYGO_DEBT_SNOWBALL_BINARY} -target wasm -opt=z -size full -scheduler=asyncify -panic=trap -no-debug ./cmd/wasm 
	wasm-opt -s -Oz $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${TINYGO_DEBT_SNOWBALL_BINARY} -o $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${TINYGO_DEBT_SNOWBALL_BINARY} && chmod +x $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${TINYGO_DEBT_SNOWBALL_BINARY}
	du -sh $(DEBT_SNOWBALL_CALCULATOR_DIR)/bin/${TINYGO_DEBT_SNOWBALL_BINARY}
	@echo "Done building tinygo binary"

build_api_gateway: 
	@echo "Building the api gateway..."
	mkdir -p ../api_gateway/bin
	cd ../api_gateway && env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o bin/${API_GATEWAY_BINARY} ./cmd
	du -sh ../api_gateway/bin/${API_GATEWAY_BINARY}
	@echo "Done building api gateway binary!"

build_link_service:
	@echo "Building link service ..."
	mkdir -p ../link_service/bin
	cd ../link_service && env GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o bin/${LINK_SERVICE_BINARY} ./cmd/api
	du -sh ../link_service/bin/${LINK_SERVICE_BINARY}	
	@echo "Done building link service binary!"

fmt: 
	gofmt -s -w ../.

### DOCKER BUILD/PUSH
docker_build: docker_webapp docker_api_gateway docker_401k_api docker_debt_snowball

docker_wiki_prod:
	@echo "Building the wiki for production..."
	cd ${WIKI_DIR}/scripts && ./backup.sh build_prd
	cd ${WIKI_DIR} && docker build -t cultivate-finance-wiki-service -f wiki.dockerfile .
	docker tag cultivate-finance-wiki-service $(DOCKERHUB_REPO)/cultivate-finance-wiki-service

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

docker_debt_snowball: build_debt_snowball_api
	@echo building docker debt snowball calculator container...
	cd $(DEBT_SNOWBALL_CALCULATOR_DIR) && docker build -t cultivate-finance-debt-snowball-calculator-service -f debt_snowball_calculator.dockerfile .
	docker tag cultivate-finance-debt-snowball-calculator-service $(DOCKERHUB_REPO)/cultivate-finance-debt-snowball-calculator-service

docker_push: docker_build
	docker push $(DOCKERHUB_REPO)/cultivate-finance-webapp-service
	docker push $(DOCKERHUB_REPO)/cultivate-finance-api-gateway-service
	docker push $(DOCKERHUB_REPO)/cultivate-finance-401k-calculator-service
	docker push $(DOCKERHUB_REPO)/cultivate-finance-debt-snowball-calculator-service

docker_push_webapp: docker_webapp
	docker push $(DOCKERHUB_REPO)/cultivate-finance-webapp-service

docker_push_wiki_prod: docker_wiki_prod
	docker push $(DOCKERHUB_REPO)/cultivate-finance-wiki-service		
	
##### K8S
### Development (Minikube)
k_deploy: k_delete k_deploy_nats k_deploy_api k_deploy_webapp k_deploy_401k k_deploy_ingress_staging

k_delete: 
	kubectl delete pods,services,deployment,ingress --all

k_deploy_nats:
	kubectl apply -f k8s/nats.yml

k_deploy_webapp:
	kubectl apply -f k8s/webapp.yml

k_deploy_401k:
	kubectl apply -f k8s/401k_calculator.yml

k_deploy_api:
	kubectl apply -f k8s/api_gateway.yml

k_deploy_ingress_staging:
	until kubectl apply -f k8s/ingress-traefik-staging.yml; do echo "Retrying ingress deploy..." && sleep 5; done

k_deploy_ingress_prod:
	until kubectl apply -f k8s/ingress-traefik-prod.yml; do echo "Retrying ingress deploy..." && sleep 5; done

m_deploy: m_down m_up k_deploy
	@echo "NOTE: if on macOS, you will need to tunnel your traffic by running the following: "
	@echo "sudo minikube tunnel"

m_up: 
	@echo "starting minikube..."
	minikube start 
	minikube addons enable ingress

m_down:
	@echo "deleting minikube"
	minikube delete

###### Wiki 
### Versioning
backup_wiki:
	@echo "backing up the wiki..."
	cd ${WIKI_DIR}/scripts && ./backup.sh backup

restore_wiki:
	@echo "restoring the wiki"
	cd ${WIKI_DIR}/scripts && ./backup.sh restore

### Deploys
wiki: 
	@echo "building and starting wiki"
	docker-compose up -d --build wiki

wiki_arm:
	@echo "building and starting wiki"
	docker-compose up -d --build wiki_arm 

wiki_prd: 
	@echo "building and starting wiki for a prod environment"
	docker-compose up -d --build wiki 
	cd ${WIKI_DIR}/scripts && ./backup.sh restore_prd