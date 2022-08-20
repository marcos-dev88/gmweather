FROM golang:1.19

WORKDIR /weather-app/src/gmweather

COPY . .

ENV GOPATH=/weather-app

RUN go build -ldflags "-s -w" -o bin/gmweather main.go

ENTRYPOINT ["./bin/gmweather"]

EXPOSE ${API_PORT}
