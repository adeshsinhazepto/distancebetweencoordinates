run:
	@echo Running DistanceApp . . . . . . 
	go run cmd/lengthofline/main.go

build:
	@echo Building DistanceApp . . . . . .
	go build cmd/lengthofline/main.go
	@echo Executing the build
	./main

test:
	@echo Running all the tests
	go test ./...
	