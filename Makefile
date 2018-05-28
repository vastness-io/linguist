VERSION=$(shell cat ./VERSION)
COMMIT=$(shell git rev-parse --short HEAD)
LATEST_TAG=$(shell git tag -l | head -n 1)

export VERSION COMMIT LATEST_TAG
.PHONY: test

setup:
	./hack/update-github-linguist.sh

test:
	@echo "=> Running tests"
	./hack/run-tests.sh

build:
	./hack/cross-platform-build.sh

verify:
	./hack/verify-version.sh

container: build
	docker build -t quay.io/vastness/linguist:${COMMIT} .

push: container
	docker push quay.io/vastness/linguist:${COMMIT}
	docker tag quay.io/vastness/linguist:${COMMIT} quay.io/vastness/linguist:${VERSION}
	docker push quay.io/vastness/linguist:${VERSION}
	docker tag quay.io/vastness/linguist:${COMMIT} quay.io/vastness/linguist:latest
	docker push quay.io/vastness/linguist:latest
