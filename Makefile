include .env
export

default: hello help
hello:
	@echo "Hello!"
## start: Start server
start: build
	@echo "> Start server..."
	@./server
## build: Building server
build:
	@echo "> Building..."
	@go build -o server cmd/main.go
	@echo "Success!"
## clean: remove server.exe 
clean:
	@echo "> Clean..."
	@rm -rf server
	@echo "Success!"
help: Makefile
	@echo "Choose a command run in "httpCRUD":"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
