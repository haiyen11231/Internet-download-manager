generate: 
	protoc -I=. \
		--go_out=internal/generated \
		--go-grpc_out=internal/generated \
		--grpc-gateway_out=internal/generated \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out . \
		api/go_load.proto

clean_newsfeed:
	rm internal/grpc/pb/newsfeed/*.go