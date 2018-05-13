# go-chain
Simple blockchain in Go

# Build

In the root directory

```
go build
```

# Usage

## Node

Run the first node on port **3000**

```
HTTP_PORT="3000" $GOPATH/bin/go-chain
```

Run the second node on port **4000**

```
HTTP_PORT="4000" $GOPATH/bin/go-chain
```

## Peers

Get current peers connected 

```
curl -X GET http://localhost:3000/peer 
```

Add second node to the first node as a peer

```
curl -X POST \
  http://localhost:3000/peer \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{ "Peers":["ws://localhost:4000"]}'
```


# Endpoints

It exposes an HTTP server with the following endpoints:

* GET   `/ping`
* GET   `/block`
* POST  `/block`  Body: `{ "Value":100000000, "Input":"coinbase", "Output": "@tiero" }`
* GET   `/peer` 
* POST  `/peer` Body: `{ "Peers" : ["ws://localhost:4000"] }`








