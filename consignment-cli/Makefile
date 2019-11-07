build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns \
	-e DB_HOST=mongodb:27017 --link mongodb:mongodb microservices-in-golang_consignment-cli
