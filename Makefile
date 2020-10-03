include .env
export

default: help

## build: Building server.
build:
	@echo "***** Building... *****"
	@go build -o server cmd/main.go
	@echo "Success!"
	
## start: Start server.
start: build
	@echo "***** Start server... *****"
	@echo "LINK -> http://localhost:$(PORT) <-"
	@./server



## clean: Remove server.exe 
clean:
	@echo "***** Clean... *****"
	rm -rf server
	@echo "Cleared!"

# DOCKER----------------------------------------------------------------------------------
## docker: call 2 commands - "docker-build", "docker-up"
docker: docker-up

## dr-build: Build docker image with name "server"
docker-build:
	@echo "***** Build image *****"
	@docker build -t server .
	@echo "***** Building SUCCESS! *****"

## dr-up: Docker run image - server
docker-up: dr-build
	@echo "***** Docker start *****"
	@echo "LINK -> http://localhost:$(PORT) <-"
	@docker run --rm --env-file ./.env --name httpServer -p $(PORT):$(PORT) server

## docker-clean: Delete server image
docker-clean:
	@echo "***** Clean docker image *****"
	@echo ">>> 1) docker system prune"
	docker system prune
	@echo "***** SUCCESS *****"
	@echo ">>> 2) docker rmi -f server"
	docker rmi -f server
	@echo "***** SUCCESS *****"
	@echo "Cleared!"
# DOCKER ---------------------------------------------------------------------------------	

help: Makefile
	@echo "Choose a command run in "httpCRUD". Example: make [command]"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

