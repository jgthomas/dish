
PROJECT_NAME=dish
PROJECT_PATH=./cmd/dockerish/dockerish.go

RUN_COMMAND=run
TEST_COMMAND=echo


run:
	go run ${PROJECT_PATH} ${RUN_COMMAND} ${TEST_COMMAND} yes

build:
	go build -o ${PROJECT_NAME} ${PROJECT_PATH}