.PHONY: test


setup:
	./hack/update-github-linguist.sh

test:
	@echo "=> Running tests"
	go test ./test/* -v