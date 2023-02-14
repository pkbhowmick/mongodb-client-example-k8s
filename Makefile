REGISTRY ?= "pkbhowmick"
REPOSITORY ?= "mongo-ping"
VERSION ?= "v0.0.1"

build-client:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/client cmd/mongodb.go

client-docker:
	docker build . -t ${REGISTRY}/${REPOSITORY}:${VERSION}
