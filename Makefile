ifeq ($(OS),Windows_NT)
    RM=-del
    APP_BINARY_PATH=.\bin\fsutils.exe
else
    RM=-rm
    APP_BINARY_PATH=./bin/fsutils
endif
# Use linker flags to provide version/build settings to the target
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"


all: build test
 
build:
	go build $(LDFLAGS) -o ./bin/ ./...

test:
	go test -v ./...
 
run:build
	./${APP_BINARY_PATH}

fmt:
	@gofmt -l -w ./...

 
.PHONY: clean
clean:
	go clean
	$(RM) ${APP_BINARY_PATH}