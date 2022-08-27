run:
	./scripts/docker_net_conf.sh;
	docker-compose up -d;
	go run .


build:
	go build -ldflags "-s -w" -o ./bin/gmweather *.go


