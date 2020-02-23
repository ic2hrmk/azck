# Lit all back service on the project
SERVICES=link

install-dependency-manager:
	go get -u github.com/golang/dep/cmd/dep

run-dependency-manager:
	dep ensure -v

build:
	go build -o azck entry/entry.go

install-linter:
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.12.3

run-linter:
	golangci-lint run -v

run-tests:
	cd test-pack && $(MAKE) run-all-tests
