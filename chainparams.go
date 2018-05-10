package main

// GenesisBlockHash
// SHA256 hash of index, prevHash, txData, timestamp
const GenesisBlockHash string	= "3cd45a480c2601ed55245eac8b233c680f111eaad30c568a318e5213f7f0f522"

// GenesisTimestamp
// Wednesday 9th May 2018 10:16:19 PM UTC
const GenesisTimestamp int64	= 1525904179 
//const EndIssuanceTimestamp int64	= 2788208179

// TotalSupply
// 21 milions
//const TotalSupplyValue int		= 21 * 10**6 * 10**8 

// InitialBlockReward 
// 50 coins
// 1 coin = 10 ** 8 tiero
const InitialBlockReward int		= (50 * 100000000) 

// CoinbaseInput
// represent the string used in the special case of no previous unspent output, the transaction representing the block reward
const CoinbaseInput string		= "COINBASE"

		
