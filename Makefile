ifneq ($(GOPATH),)
  prefix ?= $(GOPATH)
endif
prefix ?= /usr/local
exec_prefix ?= $(prefix)
ifneq ($(GOBIN),)
  bindir ?= $(GOBIN)
endif
bindir ?= $(exec_prefix)/bin

.PHONY: all install uninstall clean test update-vendor

all:
	env GO111MODULE=on go build -mod vendor -v .

install:
	env GO111MODULE=on GOBIN=$(bindir) go install -mod vendor -v .

uninstall:
	rm -f $(bindir)/vhcwallet

clean:
	rm -f vhcwallet

test:
	env GO111MODULE=on ./run_tests.sh

update-vendor:
	rm -rf vendor
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
	env GO111MODULE=on go mod vendor
