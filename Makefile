build:
	docker build -f internal/service.addition/Dockerfile -t samstarling/s-addition .
	docker build -f internal/service.hello-world/Dockerfile -t samstarling/s-hello-world .

run:
	docker run -d -p 127.0.0.1:9000:8080/tcp samstarling/s-addition:latest
	docker run -d -p 127.0.0.1:9001:8080/tcp samstarling/s-hello-world:latest

test:
	go run bin/client.go
