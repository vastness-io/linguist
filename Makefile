VERSION=$(shell cat ./VERSION)
COMMIT=$(shell git rev-parse --short HEAD)
LATEST_TAG=$(shell git tag -l | head -n 1)

export VERSION COMMIT LATEST_TAG
.PHONY: test

setup:
	./hack/update-github-linguist.sh

test:
	@echo "=> Running tests"
	go test ./... -v

build:
	./hack/cross-platform-build.sh

verify:
	./hack/verify-version.sh
