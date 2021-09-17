APP_EXECUTABLE="out/myshop"

test:
	go test ./...

build:
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE) ./cmd/

run: build
	./$(APP_EXECUTABLE)