# go-topdown-mmo-client

### Dependencies

Debian / Ubuntu
`sudo apt install libc6-dev libglu1-mesa-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config`

Fedora
`sudo dnf install mesa-libGLU-devel mesa-libGLES-devel libXrandr-devel libXcursor-devel libXinerama-devel libXi-devel libXxf86vm-devel alsa-lib-devel pkg-config`

Solus
`sudo eopkg install libglu-devel libx11-devel libxrandr-devel libxinerama-devel libxcursor-devel libxi-devel libxxf86vm-devel alsa-lib-devel pkg-config`

Arch
`sudo pacman -S mesa libxrandr libxcursor libxinerama libxi pkg-config`

Alpine
`sudo apk add alsa-lib-dev libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev mesa-dev pkgconf`


Finally resolve ebiten dependency with `go get github.com/hajimehoshi/ebiten/v2`

### Run

Make sure go-topdown-mmo-server is running and run with `go run main.go`

#### TODO
  - Docker and detailed project description
