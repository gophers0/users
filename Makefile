dev:
	docker-compose -f docker-compose-dev.yml up -d
dev-down:
	docker-compose -f docker-compose-dev.yml down

docker:
	docker-compose up -d

docker-down:
	docker-compose down

build:
	env GOOS=linux GOARCH=amd64 go build -o ./_build/users .

deploy_dev: build
	scp ./_build/marketplace s1:/srv/users/users_new
	ssh -t market-dev 'systemctl stop users.service'
	ssh -t market-dev 'rm /srv/users/current/users && rm /srv/users/users && mv /srv/users/users_new /srv/users/users && ln -s /srv/users/users /srv/users/current/users'
	ssh -t market-dev 'systemctl start users.service'
