docker-build:
	docker build . -t jbshaw/auth-service
proto-generate:
	protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative protob/auth_service.proto
