TARGET=xinshou.net

format:
	go fmt ./...

run: build
	./${TARGET}

build:
	go mod vendor && go build
	
.PHONY: run build
