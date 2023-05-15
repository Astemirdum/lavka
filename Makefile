GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.20","$(shell printf "$(GO_VERSION_SHORT)\n1.20" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.20. Found: $(GO_VERSION_SHORT))
endif

SERVICE_NAME=lavka
SERVICE_PATH=.
ENV = .env

.PHONY: run
run:
	docker compose -f ./docker-compose.yaml --env-file $(ENV) up -d

.PHONY: run-svc
run-svc: #  make run-svc svc=redis
	docker compose -f ./docker-compose.yaml --env-file $(ENV) up -d $(svc)

.PHONY: stop
stop:
	docker compose -f ./docker-compose.yaml --env-file $(ENV) down

.PHONY: lint
lint:
	go vet ./...
	golangci-lint run --fix

.PHONY: test
test:
	go test -v -race -timeout 90s -count=1 -shuffle=on  -coverprofile cover.out ./...
	@go tool cover -func cover.out | grep total | awk '{print $3}'
	go tool cover -html="cover.out" -o coverage.html

.PHONY: gen-sqlboiler
gen-sqlboiler:
	sqlboiler psql -c sqlboiler.toml \
    --add-soft-deletes \
    --wipe \
    --always-wrap-errors \
    --add-enum-types \
    --no-tests


.PHONY: image-build
image-build:
	docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .

.PHONY: image-push
image-push:
	docker push ${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: docker-login
docker-login:
	docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASS}


.PHONY: .deps
deps: .deps
.deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2
	go mod download

.PHONY: build
build:
	CGO_ENABLED=0  go build \
			-o ./bin/${SERVICE_NAME} ./cmd/main.go

.PHONY: clean
clean:
	rm bin/${SERVICE_NAME}

.PHONY: upgrade
upgrade:
	go get -u -t ./... && go mod tidy -v

.PHONY: mocks
mocks:
	cd internal/handler; go generate;


.PHONY: helm-init
helm-init:
	helm install ${MY_RELEASE} helm/lavka-app -f helm/lavka-app/values.yaml --set setter=kek,setter=kek1 -n ${NAMESPACE} --create-namespace
#	@if [ ${MY_RELEASE} -z ]; then \
#  		helm install --generate-name  helm/lavka-app -f helm/lavka-app/values.yaml --set setter=kek,setter=kek1; \
#  	else helm install ${MY_RELEASE} helm/lavka-app -f helm/lavka-app/values.yaml --set setter=kek,setter=kek1;\
#  	fi

.PHONY: helm-upgrade
helm-upgrade:
	helm upgrade ${MY_RELEASE} helm/lavka-app --namespace ${NAMESPACE}

.PHONY: k-clean
k-clean:
	kubectl delete sc,pvc,pv,cm,ing,secret,svc,all --all -n ${NAMESPACE}

.PHONY: helm-rollout
helm-rollout:
	helm rollback ${MY_RELEASE} --namespace ${NAMESPACE}