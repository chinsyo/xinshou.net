TARGET=xinshou.net

run: build
	./${TARGET}

format: 
	go fmt ./...

build: proxy
	go mod vendor && go build

clean: proxy
	go mod tidy && rm -rf ./${TARGET}

proxy:
	export GOPROXY=goproxy.cn

.PHONY: run build proxy
