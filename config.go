package main

import (
	"flag"
	"strings"
	"time"
)

//NodeIDType  string
type NodeIDType string

//AddressType string
type AddressType string

//Config object type
type Config struct {
	// ID identifies a uniqe string id to recognize the server in the
	// cluster forever
	ID NodeIDType
	// Address is the transport address composed by host:port format
	// identifies the current peer in the cluster
	Address AddressType
	// Peers includes the local address and all peers to be included in
	// the cluster. Each peer is unique, no double entries
	Peers []AddressType
	// HeartbeatTimeout specifies the time in follower state without
	// a leader before we attempt an election.
	HeartbeatTimeout time.Duration
	// ElectionTimeout specifies the time in candidate state without
	// a leader before we attempt an election.
	ElectionTimeout time.Duration
}

func (config Config) isValid() bool {

	if config.ID != "" && config.Address != "" && len(config.Peers) > 0 &&
		config.HeartbeatTimeout > 0 && config.ElectionTimeout > 0 {
		return true
	}

	return false
}

func generateConfigFromFlags() Config {
	//Flag Commannds
	//REFACTOR better cli or conf file
	host := flag.String("host", "", "remote host")
	port := flag.String("port", "", "port")
	id := flag.String("id", "", "hort")
	peers := flag.String("peers", "", "initial peers")
	flag.Parse()

	initialPeers := strings.Split(*peers, ",")
	serverAddress := *host + ":" + *port
	initialPeers = append(initialPeers, serverAddress)

	var initialPeersAddresses []AddressType
	for _, ip := range initialPeers {
		//TODO Add url scheme checks
		initialPeersAddresses = append(initialPeersAddresses, AddressType(ip))
	}

	return Config{NodeIDType(*id), AddressType(serverAddress), initialPeersAddresses, (5 * time.Millisecond), (5000 * time.Millisecond)}
}
