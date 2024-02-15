.PHONY: build

test:
	go test ./...

build:
	go build -o resizecanvas.exe cmd/resizecanvas/main.go
