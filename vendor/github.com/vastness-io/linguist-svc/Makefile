generate:
	@echo "=> generating stubs"
	protoc *.proto --go_out=plugins=grpc:.
	@echo "=> generating mocks"
	mockgen github.com/vastness-io/linguist-svc LinguistClient,LinguistServer > mock/linguist/linguist_mock.go