# go-chain
Simple blockchain in Go
 
It exposes an HTTP server on port `3000` with the following endpoints:

* GET `/ping`
* GET `/block`
* POST `/block`  Body: `{ "Value":100000000, "Input":"coinbase", "Output": "@tiero" }`








