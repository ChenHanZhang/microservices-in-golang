build:
	# this method has already been deprecated.
	# now using docker-compose to build container
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run:
	docker run -p 50052:50052 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns \
		-e DB_HOST=mongodb:27017 --link mongodb:mongodb microservices-in-golang_vessel-service