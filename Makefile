BIN_DIR=./cmd

build:
	go build -o ${BIN_DIR}/bin/mq ${BIN_DIR}/main.go

clean:
	rm -rf ${BIN_DIR}/bin

.PHONY: build clean

