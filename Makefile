
PROJECT_NAME=dish
PROJECT_PATH=./cmd/dockerish/dockerish.go

RUN_COMMAND=run
SHELL_COMMAND=/bin/bash

launch:
	go run ${PROJECT_PATH} ${RUN_COMMAND} ${SHELL_COMMAND}

build:
	go build -o ${PROJECT_NAME} ${PROJECT_PATH}