setup:
	docker pull nginx

start:
	docker run --name nginx-demo --publish 8080:80 --detach --rm nginx

inspect:
	docker ps --filter "name=nginx-demo"

test:
	curl http://localhost:8080

stop:
	docker stop nginx-demo 2> /dev/null

teardown:
	docker rmi `docker images --quiet nginx`