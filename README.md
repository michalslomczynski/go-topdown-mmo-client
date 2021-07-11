# go-topdown-mmo-client
## Docker
Using containers might reduce performance. If it does matter to you, use local build.
Reduced performance might be useful to observe how interpolation of players position works (shadowing method without queue).
### Linux

Client docker image uses host X11 server to display GUI which requires authentication in most cases.
The fastest and simplest method to bypass authentication is to run `xhost +local:root` before use and `xhost -local:root` afterwards (it is important to not forget to execute the latter - [reference](http://wiki.ros.org/docker/Tutorials/GUI "details")).

To run client use `docker-compose build` and `docker-compose up`.

## Local build
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

[reference](https://ebiten.org/documents/install.html "reference")

### Run
Make sure go-topdown-mmo-server is running and run with `go run main.go`
