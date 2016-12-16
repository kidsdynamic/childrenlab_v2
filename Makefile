.PHONY: deps vet test dev clean build build-image push-image ecr-login deploy-prod

PACKAGES = $(shell glide novendor)
VERSION = $(shell git describe --tags)
BUILD_LDFLAGS = "\
          -X cacoo-site/version.GITCOMMIT=`git rev-parse --short HEAD` \
          -X cacoo-site/version.VERSION=$(VERSION)"

ECS_REGION = us-west-1
ECR_REPO_URL = 631054961367.dkr.ecr.us-west-1.amazonaws.com/cacoo-site
DOCKER_TAG ?= latest

default: test

all: deps fmt test build build_image push_image

deps:
	glide install

vet:
	go vet $(PACKAGES)

test: vet
	go test -v -cover $(PACKAGES)

dev:
	CACOO_DEBUG=1 go run -ldflags=$(BUILD_LDFLAGS) main.go

clean:
	rm -f build/cacoo-site
	find . -name '*.test' -delete

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -ldflags=$(BUILD_LDFLAGS) -o build/cacoo-site *.go

build-image:
	docker build --rm -t cacoo-site:$(DOCKER_TAG) .

ecr-login:
	aws ecr get-login --region $(ECS_REGION) | sh -

push-image:
	docker tag cacoo-site:$(DOCKER_TAG) $(ECR_REPO_URL):$(DOCKER_TAG)
	docker push $(ECR_REPO_URL):$(DOCKER_TAG)

deploy-prod:
	docker run --rm \
      -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
      -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
      silintl/ecs-deploy \
      -r $(ECS_REGION) \
      -c cacoo-cluster-1 \
      -n site \
      -i $(ECR_REPO_URL):$(DOCKER_TAG)
