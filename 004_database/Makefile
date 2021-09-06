start:
	docker-compose up --build

inspect:
	docker-compose ps

psql:
  # password is "password"
	docker-compose run database psql --host=database --user=developer --dbname=default_db

test:
	curl http://localhost:9000/

stop:
	docker-compose down