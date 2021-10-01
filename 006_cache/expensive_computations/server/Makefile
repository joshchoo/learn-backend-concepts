setup:
	docker build -t fibonacci .

start:
	docker run --name fibonacci-app --publish 9000:9000 --detach --rm -e "REDIS_URI=localhost:6379" fibonacci

test:
	time curl -X POST http://localhost:9000 -d "44"

stop:
	docker stop fibonacci-app

teardown:
	docker rmi `docker images --quiet fibonacci`