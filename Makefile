MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
VERSION:=${MAIN_VERSION}\#$(shell git log -n 1 --pretty=format:"%h")
PACKAGES:=$(shell go list ./... | sed -n '1!p' | grep -v /vendor/)
RELEASE_FILE:=server-${MAIN_VERSION}.tar.gz
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
	mkdir -p build
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -a -o build/server server.go
	cp -r config build
	cd build && tar zcf ${RELEASE_FILE} *

clean:
	rm -rf build server coverage.out coverage-all.out
