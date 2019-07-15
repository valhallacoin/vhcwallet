all:
	env GO111MODULE=on go install -mod vendor -v .

.PHONY: clean test update-vendor
clean:
	rm -f vhcwallet

test:
	env GO111MODULE=on ./run_tests.sh

update-vendor:
	rm -rf vendor
	env GO111MODULE=on go get -u
	env GO111MODULE=on go mod tidy -v
	env GO111MODULE=on go mod vendor
