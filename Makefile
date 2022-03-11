all: deps test

deps: tidy vend

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -coverpkg=github.com/sgmccullough/PokeAssist ./...

test-client:
	go test -v ./.

tidy:
	go mod tidy -v

vend:
	go mod vendor -v

build:
	go build