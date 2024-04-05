BIN_FOLDER=bin
BIN_NAME=gli

.DEFAULT_GOAL := build

build:
	go build -o ${BIN_FOLDER}/${BIN_NAME} .

run: build
	./${BIN_FOLDER}/${BIN_NAME}

test:
	go test -v ./... -count=1

clean:
	rm -rf ${BIN_FOLDER}

