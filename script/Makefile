
VERSION ?= latest

run: build
	@echo "Running..."
	@./bin/my-go-starter

build:
	@echo "Building..."
	@go build -o ./bin/my-go-starter ./cmd/my-go-starter/main.go

up: 
	@echo "Starting..."
	@go run ./cmd/my-go-starter/main.go


migration:
	@echo "Creating migration..."
	@migration create -ext sql -dir ./migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@echo "Creating seed..."
	@go run ./cmd/seed/main.go

build-image:
	@echo "Building docker image with version: $(VERSION)"
	@docker build -t my-go-starter:$(VERSION) .

push-image:
	@echo "Pushing docker image with version: $(VERSION)"
	@docker push my-go-starter:$(VERSION)

deploy:
	@cd deploy
	@echo "Deploying with helm..."
	@helm upgrade --install my-go-starter ./chart/my-go-starter --set image.tag=$(VERSION) --set image.repository=my-go-starter

deploy-kubectl-apply:
	@echo "Deploying with kubectl..."
	@kubectl apply -f ./deploy/k8s/deployment.yaml
	@kubectl apply -f ./deploy/k8s/service.yaml
