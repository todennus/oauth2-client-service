start-rest:
	go run ./cmd/main.go rest

docker-build:
	docker build -t todennus/oauth2-client-service -f ./build/package/Dockerfile .
