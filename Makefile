build:
	@go build -o pw .

install:
	@bash install.sh
format:
	@echo "formatting server code...."
	@gofmt -l -s -w .

run: 
	@go run main.go