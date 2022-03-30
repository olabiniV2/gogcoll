GO_FILES := *.go
SOURCE_FILES := $(GO_FILES)

GO := go
GOTEST := $(GO) test
GOINSTALL := $(GO) install

COVERPROFILE := coverprofile

.PHONY := default clean test ci-test ci-deps ci-upload-coverage

default: test

test:
	$(GOTEST) ./...

$(COVERPROFILE): $(SOURCE_FILES)
	$(GOTEST) -cover -coverprofile $@ ./...

coverage: $(COVERPROFILE)
	$(GO) tool cover -html $(COVERPROFILE)

coverage-tails: $(COVERPROFILE)
	$(GO) tool cover -html $(COVERPROFILE) -o ~/Tor\ Browser/coverage.html
	xdg-open ~/Tor\ Browser/coverage.html

# ------------------------------------- Tasks for CI -----------------------------------------

ci-test: $(COVERPROFILE)

ci-deps:
	$(GOINSTALL) github.com/mattn/goveralls@latest

ci-upload-coverage: $(COVERPROFILE) ci-deps
	goveralls -coverprofile=$(COVERPROFILE)
