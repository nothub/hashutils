all: lint check

.PHONY: lint
lint:
	go vet

.PHONY: check
check:
	go test -v -parallel $(shell grep -c -E "^processor.*[0-9]+" "/proc/cpuinfo") $(shell go list -m)/...
