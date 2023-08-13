server:
	go run main.go

build:
	go build -o main.go

d.up:
	docker-compose up 

d.down:
	docker-compose down

d.up.build:
	docker-compose --build up