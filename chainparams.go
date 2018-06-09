package main

// GenesisBlockHash SHA256 hash of index == 0, term == 1, prevHash = "0", txData, timestamp
const GenesisBlockHash string = "e891fee258b64b4b017c07a193911c96a9499e52c776f64e5d054715a20068d6"

// GenesisTimestamp Wednesday 9th May 2018 10:16:19 PM UTC
const GenesisTimestamp uint64 = 1525904179

//const EndIssuanceTimestamp uint64	= 2788208179

// TotalSupply
// 21 milions
//const TotalSupplyValue int		= 21 * 10**6 * 10**8

// InitialBlockReward 50 coins
// 1 coin = 10 ** 8 tiero
const InitialBlockReward uint64 = (50 * 100000000)

// CoinbaseInput represent the string used in the special case of no previous unspent output, the transaction representing the block reward
const CoinbaseInput string = "COINBASE"
