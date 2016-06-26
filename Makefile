VERSION = $(shell git tag | tail -n 1)
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
RUNTIME_GOPATH := $(GOPATH):$(shell pwd)
SRC := $(wildcard *.go) $(wildcard src/agg/*.go)
TEST_SRC := $(wildcard src/agg/*_test.go)

all: agg

agg: $(SRC)
	GOPATH=$(RUNTIME_GOPATH) go build

test: $(TEST_SRC)
	GOPATH=$(RUNTIME_GOPATH) go test -v $(TEST_SRC)

stringer:
	cd src/agg && stringer -type AggregateType agg.go

dev_dep:
	go get github.com/stretchr/testify
	go get golang.org/x/tools/cmd/stringer

clean:
	rm -f agg agg.exe *.gz *.zip

package: clean agg
ifeq ($(GOOS),windows)
	zip agg-$(VERSION)-$(GOOS)-$(GOARCH).zip agg.exe
else
	gzip -c agg > agg-$(VERSION)-$(GOOS)-$(GOARCH).gz
endif
