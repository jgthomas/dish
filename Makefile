
PROJECT_NAME=dish
PROJECT_PATH=./cmd/dockerish/dockerish.go

build:
	go build -o ${PROJECT_NAME} ${PROJECT_PATH}