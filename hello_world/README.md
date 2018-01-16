# hello_world
### DESCRIPTION:	  Run a Golang 'Hello World!' webapp in a container

Compile for linux amd64:
>env GOOS=linux GOARCH=amd64 go build .\hello_world.go

Build go_hello_world image
>docker build -t go_hello_world .

Run go_hello_world container
>docker run -d -p 8080:8080 go_hello_world