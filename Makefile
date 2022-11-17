docker-build:
	docker build --no-cache -t relistan/nameifier:`git rev-parse --short HEAD` .

docker-push:
	docker push relistan/nameifier:`git rev-parse --short HEAD` 
