all: build

build:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/pretty-imp main.go
	GOOS=linux GOARCH=amd64 go build -o bin/linux/pretty-imp main.go
	GOOS=windows GOARCH=386 go build -o bin/windows/pretty-imp.exe main.go

clean:
	rm -r bin/

format:
	go fmt
