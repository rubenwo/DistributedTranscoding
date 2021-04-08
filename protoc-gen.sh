#protoc --proto_path=api/proto/v1 --go-grpc_out:. distributor-service.proto
protoc --proto_path=api/proto/v1 --go_out=. --go-grpc_out=. distributor-service.proto