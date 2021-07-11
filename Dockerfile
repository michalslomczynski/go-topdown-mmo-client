FROM ubuntu:latest

ENV TZ=Europe/Warsaw

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && \
    apt update && apt install -y \
    golang libc6-dev libglu1-mesa-dev \
    libgl1-mesa-dev libxcursor-dev libxi-dev \
    libxinerama-dev libxrandr-dev libxxf86vm-dev \
    libasound2-dev pkg-config ca-certificates xauth

WORKDIR /home/client/

COPY . .

RUN go get github.com/hajimehoshi/ebiten/v2

CMD ["go", "run", "main.go"]