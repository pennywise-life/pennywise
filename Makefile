build:
	@go build -o pw .

install:
	@bash install.sh
format:
	@echo "formatting server code...."
	@gofmt -l -s -w .

run: 
	@go run main.go

test:
	@go test -covermode=count -coverpkg=./... -coverprofile cover.out -v ./... 
	@go tool cover -html cover.out -o cover.html

total_test:
	@go tool cover -func cover.out