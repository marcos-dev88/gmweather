run:
	bash ./scripts/docker_net_conf.sh;
	docker-compose up;


build:
	go build -ldflags "-s -w" -o ./bin/gmweather *.go


