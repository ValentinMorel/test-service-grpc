# test-service-grpc

## Installation 

Need :
- go 1.19

A makefile is here to simplify everything you need: 

```
make build-deps
make run-server
```

In another terminal, start the client: 

```
make run-whois USER=fra
```

or: 

```
make run-list
```

Feel free to change the port and the expiration duration for the cache in the config.json !
