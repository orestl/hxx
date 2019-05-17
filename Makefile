PWD          ?= $(MAKEFILE)

VERSION      ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(CURDIR)/VERSION 2> /dev/null || echo v0)
TIMESTAMP    ?= $(shell date +%FT%T%z)
BUILDID      ?= $(shell xxd -l 16 -ps -c 16 /dev/random)

LDFLAGS      ?= "-X main.Version=$(VERSION) -X main.BuildTimestamp=$(TIMESTAMP) -B=0x$(BUILDID)"

PACKAGES     := $(shell go list -f '{{ .Dir }}' ./... )
IMPORT_PATHS := $(shell go list -f '{{ .ImportPath }}' ./... )
SOURCES      := $(shell go list -f '{{ $$outer := . }}{{range .GoFiles}}{{ $$outer.Dir }}/{{.}} {{end}}' ./... )
TEST_SOURCES := $(shell go list -f '{{ $$outer := . }}{{range .TestGoFiles}}{{ $$outer.Dir }}/{{.}} {{end}}' ./... )

.PHONY: build install cover bench test clean objdump godist

build: vendor $(SOURCES)
	go build -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) .

install: vendor $(SOURCES)
	go install -ldflags=$(LDFLAGS) -gcflags=$(GCFLAGS) .

#
# developer targets
#

cover: vendor $(TEST_SOURCES)
	COVERFILE=$(shell mktemp) && for i in $(PACKAGES); do \
		go test $(TEST_FLAGS) -p=1 -coverprofile=$$COVERFILE $$i; \
		go tool cover -func=$$COVERFILE; \
	done

bench: vendor $(TEST_SOURCES) $(SOURCES)
	go test $(BENCH_FLAGS) -bench=. ./...

test: vendor $(TEST_SOURCES) $(SOURCES)
	go test $(TEST_FLAGS) ./...

clean:
	rm -f $(BIN)

objdump:
	for i in $(IMPORT_PATHS); do \
		go tool objdump -S -s "$$i[.].*" $(BIN); \
	done

godist:
	go tool dist banner
	go tool dist env


