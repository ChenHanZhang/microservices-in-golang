init:
	protoc  --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
	 	--proto_path=${GOPATH}/src:.  --go_out=plugins=micro:. ./consignment/consignment.proto

	protoc  --plugin=protoc-gen-go=${GOPATH}/bin/protoc-gen-go \
	 	--proto_path=${GOPATH}/src:. --go_out=plugins=micro:. ./vessel/vessel.proto