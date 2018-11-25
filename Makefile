IMAGE_NAME := tehwey/paste-my-taste
VERSION_DOCKER := $(shell git describe --abbrev=0 --tags  | sed 's/^v\(.*\)/\1/')

all: install

install:
	go install -v

test:
	go test ./... -v

image-push-staging:
	npm --prefix web/ run build
	docker build -t $(IMAGE_NAME):staging .
	docker push $(IMAGE_NAME):staging

image-push:
	npm --prefix web/ run build
	docker build -t $(IMAGE_NAME):latest .
	docker tag $(IMAGE_NAME):latest $(IMAGE_NAME):$(VERSION_DOCKER)
	docker push $(IMAGE_NAME):latest
	docker push $(IMAGE_NAME):$(VERSION_DOCKER)

release:
	git tag -a $(VERSION) -m "Release $(VERSION)" || true
	git push origin $(VERSION)
	goreleaser --rm-dist

.PHONY: install test