
MAIN_PATH=unstaged.go
BUILD_PATH=build/
LINUX_BINARY_NAME=unstaged_linux
DARWIN_BINARY_NAME=unstaged_darwin
WINDOWS_BINARY_NAME=unstaged_windows.exe
# export UNSTAGED_COMMIT_HASH="$(git rev-parse HEAD)"
GIT_COMMIT_HASH=${UNSTAGED_COMMIT_HASH}
# export UNSTAGED_VERSION="x.x.x"
VERSION=${UNSTAGED_VERSION}


build:
	make build-linux
	make build-darwin
	make build-windows

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_PATH)$(LINUX_BINARY_NAME) -v $(MAIN_PATH)

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_PATH)$(DARWIN_BINARY_NAME) -v $(MAIN_PATH)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_PATH)$(WINDOWS_BINARY_NAME) -v $(MAIN_PATH)

releases:
	make build-rel-linux
	make build-rel-darwin
	make build-rel-windows
build-rel-linux:
	GOOS=linux GOARCH=amd64 go build \
			-ldflags '-X github.com/komish/unstaged/version.Version=$(VERSION) -X github.com/komish/unstaged/version.CommitHash=$(GIT_COMMIT_HASH)' \
		 	-o $(BUILD_PATH)$(LINUX_BINARY_NAME) -v $(MAIN_PATH)

build-rel-darwin:
	GOOS=darwin GOARCH=amd64 go build \
			 -ldflags '-X github.com/komish/unstaged/version.Version=$(VERSION) -X github.com/komish/unstaged/version.CommitHash=$(GIT_COMMIT_HASH)' \
			 -o $(BUILD_PATH)$(DARWIN_BINARY_NAME) -v $(MAIN_PATH)

build-rel-windows:
	GOOS=windows GOARCH=amd64 go build \
			 -ldflags '-X github.com/komish/unstaged/version.Version=$(VERSION) -X github.com/komish/unstaged/version.CommitHash=$(GIT_COMMIT_HASH)' \
			 -o $(BUILD_PATH)$(WINDOWS_BINARY_NAME) -v $(MAIN_PATH)

clean:
	rm -rf $(BUILD_PATH)
