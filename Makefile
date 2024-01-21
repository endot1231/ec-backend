.PHONY: up
up:
	docker-compose up -d

.PHONY: test-app
test-app:
	docker-compose exec app go test -v ./...

.PHONY: ssh-app
ssh-app:
	docker-compose exec -it app /bin/sh