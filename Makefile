PKGS=$(shell go list ./...)

.PHONY: test
test:
	@go test -v $(PKGS)
