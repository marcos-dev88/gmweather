FROM golang:1.19 

WORKDIR /weather-app/src/gmweather

COPY . .

ENV GOPATH=/weather-app

ENV TZ="American/Sao_Paulo"

ENV XAUTHORITY=/.Xauthority

ENV DISPLAY :0

RUN apt-get update; apt-get install -y libaio1 libaio-dev; apt-get install -y tzdata
RUN apt-get update && apt-get install libgl1-mesa-glx -y
RUN apt install libglfw3-dev libgl1-mesa-dev libglu1-mesa-dev -y
RUN apt install libx11-dev mesa-utils -y
RUN apt install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev -y
RUN apt install libc-dev libltdl-dev libxxf86vm-dev -y
RUN wget --user-agent=Mozilla -O apache-pulsar-client.deb "https://archive.apache.org/dist/pulsar/pulsar-2.4.1/DEB/apache-pulsar-client.deb"
RUN wget --user-agent=Mozilla -O apache-pulsar-client-dev.deb "https://archive.apache.org/dist/pulsar/pulsar-2.4.1/DEB/apache-pulsar-client-dev.deb"
RUN apt install -y ./apache-pulsar-client.deb
RUN apt install -y ./apache-pulsar-client-dev.deb
RUN apt-get install libglu1-mesa:i386 -y; apt-get install libglu1 -y
RUN apt-get install r-base -y; apt-get install r-recommended -y
RUN apt-get update -y; apt-get install build-essential -y
RUN apt-get install -y xterm
RUN apt-get install -qqy x11-apps

RUN chmod -R 755 /weather-app/src/gmweather

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s" -o ./bin/gmweather *.go

CMD ["./bin/gmweather"]
