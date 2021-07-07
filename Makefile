.PHONY: docker-push
docker-push: docker-build
	docker push pxsalehi/simple-sink:0.0.1

.PHONY: docker-build
docker-build: build
	docker build . -f Dockerfile -t pxsalehi/simple-sink:0.0.1

.PHONY: build
build:
	GOOS=linux go build -o simple-sink main.go

.PHONY: simple-sink
simple-sink:
	go build -o bin/simple-sink main.go