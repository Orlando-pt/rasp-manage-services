run:
	go run main.go

build:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o bin/server main.go

send-project:
	ssh rasp-server@rasp-server "mkdir -p /home/rasp-server/apps/rasp-manager && rm -rf /home/rasp-server/apps/rasp-manager/*"
	scp -r ./* rasp-server@rasp-server:/home/rasp-server/apps/rasp-manager

install:
	sudo chmod +x install.sh
	./install.sh

