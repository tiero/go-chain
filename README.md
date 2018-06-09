# go-chain
Permissioned Blockchain using [Raft](https://raft.github.io) consensus

Written in Go


# Build

In the root directory

```
go build
```


Install dep

```
go install
```

# Setup

## Node

Run the first node on port **3000**

```
$GOPATH/bin/go-chain --id=@tiero  --host=127.0.0.1 --port=3000 --peers=127.0.0.1:4000
```

Run the second node on port **4000**

```
$GOPATH/bin/go-chain --id=@alice  --host=127.0.0.1 --port=4000 --peers=127.0.0.1:3000
```

## Cluster

Alternatively you can run a cluster launching the `run-cluster` bash script in  `scripts` folder
Change manually from the script the port and the number of nodes. Default tre nodes at ports: 3000, 4000, 5000

```
sh scripts/run-cluster
```

If you close the process with `Ctrl + C`, the script will terminate the opened processes

# Usage

## Block

Get list of blocks

```
curl -X GET http://localhost:3000/block 
```

Mine new Block

```
curl -X POST \
  http://localhost:3000/block \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{ "Value":100000000, "Input":"@tiero", "Output": "@bob" }'
```

## Peers

Get list of peers

```
curl -X GET http://localhost:3000/peer 
```

Add new peers

```
curl -X POST \
  http://localhost:3000/peer \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{ "Peers":["127.0.0.1:4000"] }'
```







