build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e DB_HOST=mongodb:27017 --link mongodb:mongodb \
		-e MICRO_REGISTRY=mdns microservices-in-golang_consignment-service