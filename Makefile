TARGET=xinshou.net

run: build
	./${TARGET}

format:
	go fmt ./...

build:
	go mod vendor && go build

clean:
	rm -rf ./${TARGET}

.PHONY: run build
