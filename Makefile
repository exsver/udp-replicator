OUT_BIN = udp-replicator

export PATH := $(PATH):/usr/local/go/bin

build:
	go mod tidy
	env CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w -extldflags "-static"' -o $(OUT_BIN) -v

update:
	go get -u
	go mod tidy

download:
	go get
	go mod tidy
