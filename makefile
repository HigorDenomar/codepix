# Variables
GRPC_PATH=application/grpc

# Tasks
default: protoc	

protoc:
	@protoc --go_out=$(GRPC_PATH)/pb --go_opt=paths=source_relative --go-grpc_out=$(GRPC_PATH)/pb --go-grpc_opt=paths=source_relative --proto_path=$(GRPC_PATH)/protofiles $(GRPC_PATH)/protofiles/*.proto
