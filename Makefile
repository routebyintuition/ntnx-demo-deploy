clean:
	rm -fv server-one

server:
	GOOS=linux go build -ldflags="-s -w" -tags 'one' -o server-one main-server.go logging.go handlers.go misc.go

all: clean server
