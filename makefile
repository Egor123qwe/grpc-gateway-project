grpc:
	protoc -I./proto --proto_path=proto/api \
	--go_out=proto/api \
	--go-grpc_out=proto/api \
	--grpc-gateway_out=proto/api \
	proto/api/api.proto
