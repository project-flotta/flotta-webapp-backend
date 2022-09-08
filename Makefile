app-up:
	cp .env.example .env
	cp app/config.yaml.example app/config.yaml
	docker-compose up -d

test:
	cd app/ && go test -v ./...