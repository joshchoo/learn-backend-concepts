setup:
	docker build -t dockerfile-demo .

start:
	docker run --rm dockerfile-demo

teardown:
	docker rmi `docker images --quiet dockerfile-demo`