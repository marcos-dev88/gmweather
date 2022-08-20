run:
	./scripts/docker_net_conf.sh;


build:
	go build -ldflags "-s -w" -o ./bin/gmweather *.go


