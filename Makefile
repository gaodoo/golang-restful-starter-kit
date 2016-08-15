MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "1.0")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
LDFLAGS:=-ldflags "-X github.com/qiangxue/golang-restful-starter-kit/app.Version=${VERSION}"

default: run

test:
	go test -p=1 -cover -covermode=count ${PACKAGES}

cover:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES), \
		echo ${pkg}; \
		go test -coverprofile=coverage.out -covermode=count ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

run:
	go run ${LDFLAGS} server.go

build: clean
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -a -o server server.go

clean:
	rm -f server coverage.out coverage-all.out
