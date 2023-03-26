run:
	go run main.go

build:
	env GOOS=linux GOARCH=arm GOARM=5 go build
