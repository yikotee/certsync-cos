PROJECT:=Kavi

.PHONY: build-all clean

APP_NAME := certsync-cos
LDFLAGS := -s -w

build-all:
	@mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)-linux-amd64
	GOOS=linux GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)-linux-arm64
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o dist/$(APP_NAME)-windows-amd64.exe
	@ls -lh dist/

clean:
	rm -rf dist/
