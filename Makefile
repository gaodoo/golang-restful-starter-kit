MAIN_VERSION:=$(shell git describe --abbrev=0 --tags || echo "0.1")
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
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

run:
	go run ${LDFLAGS} server.go

build: clean
	go build ${LDFLAGS} -a -o server server.go

clean:
	rm -rf server coverage.out coverage-all.out
