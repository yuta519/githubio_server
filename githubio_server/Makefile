.PHONY: install test

install:
	go mod tidy
	go run main.go &
	go test test/*
	kill $(PID)
test:
	go run main.go &
	go test -v test/*
	kill $(PID)


PID=$$(lsof -P -i:8000 | grep "8000" | cut -w -f2)
