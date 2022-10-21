.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpdump cmd/httpdump/main.go

.PHONY: image
image:
	buildah build -t quay.io/tosmi/httpdump:latest .
