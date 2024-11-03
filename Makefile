default: tidy docker-stop docker-up run

docker-stop:
	docker-compose stop

docker-up:
	docker-compose up -d

tidy:
	go mod tidy

build:
	cd cmd/web && go build -o ../../hi-zone-frontline.go main.go

run:
	air

run-internal:
	cd cmd/internal && air

stop: docker-stop
