.PHONY: build test run

PROJECT?=github.com/eddiefisher/pinger/src
APP?=pinger
PORT?=8000

GOOS?=linux
GOARCH?=amd64

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS?=-ldflags "-w -s \
	-X ${PROJECT}/version.Release=${RELEASE} \
	-X ${PROJECT}/version.BuildTime=${BUILD_TIME} \
	-X ${PROJECT}/version.Commit=${COMMIT}"

clean:
	docker rmi $(APP):$(RELEASE) || true

buildc: clean
	cd ./src ; CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build ${LDFLAGS} -o ../bin/${APP}

container: buildc
	docker build -t $(APP):$(RELEASE) .

runc: container
	docker stop $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(APP):$(RELEASE)

# Build the project
build:
	cd ./src ; go build ${LDFLAGS} -o ../bin/${APP}

bin1:
	sh -c "./bin/${APP} -name=m1"

bin2:
	sh -c "./bin/${APP} -name=m2"

# Run the project
run:
	cd ./src ; go run -v ./*.go

# Test the project
test:
	cd ./src ; go test -v -race ./...
