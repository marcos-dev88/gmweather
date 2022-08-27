FROM golang:1.19

WORKDIR /weather-app/src/gmweather

COPY . .

ENV GOPATH=/weather-app

RUN apt-get update && apt-get install libgl1-mesa-glx -y

RUN apt install libglfw3-dev libgl1-mesa-dev libglu1-mesa-dev -y

RUN apt install libx11-dev mesa-utils -y

RUN apt install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev -y

RUN apt install libc-dev libltdl-dev libxxf86vm-dev -y

RUN wget --user-agent=Mozilla -O apache-pulsar-client.deb "https://archive.apache.org/dist/pulsar/pulsar-2.4.1/DEB/apache-pulsar-client.deb"
RUN wget --user-agent=Mozilla -O apache-pulsar-client-dev.deb "https://archive.apache.org/dist/pulsar/pulsar-2.4.1/DEB/apache-pulsar-client-dev.deb"

RUN apt install -y ./apache-pulsar-client.deb
RUN apt install -y ./apache-pulsar-client-dev.deb

RUN go build -ldflags "-s" -o ./bin/gmweather *.go

ENTRYPOINT ["./bin/gmweather"]

