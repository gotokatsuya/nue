LINTIGNOREDEPS='vendor/.+\.go'
NUE_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/misc/" | grep -v "/vendor/")

all:

unit: lint vet test

lint:
	@echo "go lint"
	@lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDEPS}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

vet:
	@echo "go vet"
	@go tool vet -all -structtags -shadow .

test:
	@go test $(NUE_ONLY_PKGS)


.PHONY: all unit lint vet test
