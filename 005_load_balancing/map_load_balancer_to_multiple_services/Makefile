start:
	cp load_balancer/default.conf load_balancer/nginx.conf && \
	docker-compose up --build --remove-orphans --detach

inspect:
	docker-compose ps

test:
	curl http://localhost:8080/

scale-up:
	make scale-up-server && \
	cp load_balancer/scale-up.conf load_balancer/nginx.conf && \
	make reload-load-balancer-config

scale-down:
	cp load_balancer/default.conf load_balancer/nginx.conf && \
	make reload-load-balancer-config && \
	docker-compose up --detach --remove-orphans

scale-up-server:
	docker-compose -f docker-compose.yaml -f docker-compose.scale-up.yaml up --detach

reload-load-balancer-config:
	docker-compose exec load_balancer nginx -t \
	&& docker-compose exec load_balancer nginx -s reload 

stop:
	docker-compose down --remove-orphans