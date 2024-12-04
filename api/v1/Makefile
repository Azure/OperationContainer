all: service
MAIN_DIR := $(shell dirname -- $(shell dirname -- $(shell dirname -- $(shell pwd))))
USER_ID := $(shell id -u):$(shell id -g)
BUF_TOKEN := $(shell echo $(BUF_TOKEN))

# api.swagger.json file is automatically generated using npx buf generate command via the openapi plugin
service:
	docker run --mount src=$(MAIN_DIR),target=/app/binded-data,type=bind \
		-e BUF_TOKEN=$(BUF_TOKEN) \
		-d build-operationcontainer \
		/bin/bash -c \
	"cd binded-data/OperationContainer/api/v1; \
	cd proto; buf mod update; cd ..; buf generate; \
	swagger-codegen-cli generate -i operationcontainer.swagger.json -l go -o restsdk -c swagger-config.json; \
	chown -R $(USER_ID) ../..; \
	cd /app; ./maintain_service populateMethodFiles \
		--protoFilePath binded-data/OperationContainer/api/v1/proto/operationcontainer.proto \
		--svcDirName binded-data/OperationContainer --deleteGarbageFiles=false;"

	# go mod related work is done outside the docker container 
	# to avoid repeated download in the docker container
	go install go.uber.org/mock/mockgen@latest
	go mod tidy # Update go.mod after code generation.
	go get go.uber.org/mock/mockgen/model # Required by mockgen.
	${HOME}/go/bin/mockgen -package mock \
		github.com/Azure/OperationContainer/api/v1 \
		OperationContainerClient > mock/mock_client.go
