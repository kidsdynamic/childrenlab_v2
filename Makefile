.PHONY: deps vet test dev cp build clean buildFront

PACKAGES = $(shell glide novendor)
DOCKER_REPO_URL = jack08300/childrenlab_v2

deps:
	dep ensure

vet:
	go vet $(PACKAGES)

build: clean
	GOOS=linux go build -o ./app/build/main ./app/*.go

clean:
	rm -rf app/build/*
	find . -name '*.test' -delete

push-image: clean build build-image
	docker tag childrenlab_v2 $(DOCKER_REPO_URL):latest
	docker push $(DOCKER_REPO_URL):latest

build-image:
	docker build --rm -t childrenlab_v2:latest .