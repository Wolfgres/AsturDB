EXECUTABLE=asturdb
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

windows:
	env GOOS=windows GOARCH=amd64 go build -o bin/$(WINDOWS)/$(EXECUTABLE) -ldflags="-s -w -X main.version=$(VERSION)" && cp -r configs bin/$(WINDOWS)/ && cp -r scripts bin/$(WINDOWS)/ && mkdir bin/$(WINDOWS)/log && cd bin/ && zip ./$(WINDOWS)/ ./$(WINDOWS).zip

linux:
	sh build/build_linux.sh $(LINUX) $(EXECUTABLE) $(VERSION)

clean: 
	cd bin/ && rm -rf ./*