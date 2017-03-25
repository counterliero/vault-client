.PHONY: build clean

BINARY=vlt

build: ${BINARY}

.get-deps: *.go
	go get -t -d -v ./...
	touch .get-deps

clean:
	rm -f .get-deps
	rm -f ${BINARY}

${BINARY}: *.go .get-deps
	go build -o $@ *.go
