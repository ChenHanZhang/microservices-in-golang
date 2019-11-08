build:
	# this method has already been deprecated.
	# now using docker-compose to build container
	docker build -t user-service .

run:
	docker run -p 50053:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns user-service