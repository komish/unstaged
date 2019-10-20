
MAIN_PATH=unstaged.go
BUILD_PATH=build/
LINUX_BINARY_NAME=unstaged_linux
DARWIN_BINARY_NAME=unstaged_darwin
WINDOWS_BINARY_NAME=unstaged_windows.exe


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

clean:
	rm -rf $(BUILD_PATH)
