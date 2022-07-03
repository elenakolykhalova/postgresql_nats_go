docker:
		docker-compose up --build
main:
		go run ./cmd/main.go
msg:	
		go run ./message_sender/sendMsg.go
