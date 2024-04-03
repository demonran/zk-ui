export PATH := $(PATH):`go env GOPATH`/bin
export GO111MODULE=on
LDFLAGS := -s -w
PLATFORM := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	env  $(PLATFORM) go build -trimpath -ldflags "$(LDFLAGS)"  -o bin/zkui ./cmd/

