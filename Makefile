docker:
	docker-compose up -d && docker logs -f users_app_1

docker-down:
	docker-compose down

deploy:
	rsync -avzhe ssh . hack:/srv/users/
