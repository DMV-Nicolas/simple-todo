run:
	go run main.go
build:
	go build -o simple-todo main.go
templ:
	templ generate
.PHONY: run build templ