# go-chain
Simple blockchain in Go

# Build

In the root directory

```
go build
```

# Usage

## Node

Run the first node

```
HTTP_PORT="3000" $GOPATH/bin/go-chain
```

Run the second node on other port

```
HTTP_PORT="3000" $GOPATH/bin/go-chain
```

## Peers

Add second node to the first as a peer

```
curl -X POST \
  http://localhost:3000/peer \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{ "Port":"3001"}'
```


# Endpoints

It exposes an HTTP server with the following endpoints:

* GET   `/ping`
* GET   `/block`
* POST  `/block`  Body: `{ "Value":100000000, "Input":"coinbase", "Output": "@tiero" }`
* GET   `/peer` 
* POST  `/peer` Body: `{ "Port" : "3000" }`








