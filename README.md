# Comman-Do

Comman-Do is a simple Todo example Application, built to demonstrate a post I wrote for my blog [devtheweb.io](devtheweb.io).

Common-Do uses [TwitchTv's fantastic Twirp library](https://github.com/twitchtv/twirp) to flesh out a small RPC API for creating and managing ToDos from a CLI.

### Usage:
```
go get -u github.com/williamhgough/comman-do/cmd/comman-doc
go get -u github.com/williamhgough/comman-do/cmd/comman-dos

Window 1:
comman-dos

Window 2:
comman-doc --help

Usage of ./comman-doc:
  -command string
        Command to interact with Todos. create, get, list, delete (default "list")
  -id int
        The ID of the todo you wish to retrieve (default 1)
  -title string
        The todo title
```

protoc --proto_path=$GOPATH/src:. --twirp_out=. --go_out=. ./rpc/todo/service.proto