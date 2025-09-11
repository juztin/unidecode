.PHONY: build_all build clean dist


build:
	@mkdir -p ./build
	go build -ldflags "-s -w" -o build/unidecode cmd/decoder/main.go

build_linux:
	@mkdir -p ./build
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o build/unidecode_linux_amd64 cmd/decoder/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o build/unidecode_linux_arm64 cmd/decoder/main.go

build_darwin:
	@mkdir -p ./build
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o build/unidecode_darwin_amd64 cmd/decoder/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o build/unidecode_darwin_arm64 cmd/decoder/main.go

build_windows:
	@mkdir -p ./build
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o build/unidecode_windows_amd64 cmd/decoder/main.go
	GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o build/unidecode_windows_arm64 cmd/decoder/main.go

build_all: build_linux
build_all: build_darwin
build_all: build_windows

clean:
	@rm -rf ./build/*
