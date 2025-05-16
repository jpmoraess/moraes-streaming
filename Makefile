#DEV
build-dev:
	docker build -t moraes-streaming -f docker/image/Dockerfile . && docker build -t turn -f docker/image/Dockerfile.turn .

run-dev:
	docker-compose -f docker/compose/dev.yml up

.PHONY: build-dev run-dev