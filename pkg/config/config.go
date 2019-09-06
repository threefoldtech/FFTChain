package config

import (
	"math/big"

	"github.com/threefoldtech/rivine/build"
	"github.com/threefoldtech/rivine/modules"
	"github.com/threefoldtech/rivine/types"
)

var (
	rawVersion = "v0.1"
	// Version of the chain binaries.
	//
	// Value is defined by a private build flag,
	// or hardcoded to the latest released tag as fallback.
	Version build.ProtocolVersion
)

const (
	// TokenUnit defines the unit of one Token.
	TokenUnit = "FFT"
	// TokenChainName defines the name of the chain.
	TokenChainName = "fftchain"
)

// chain network names
const (
	NetworkNameDevnet = "devnet"

	NetworkNameTestnet = "testnet"
)

func GetDefaultGenesis() types.ChainConstants {
	return GetTestnetGenesis()
}

// GetBlockchainInfo returns the naming and versioning of tfchain.
func GetBlockchainInfo() types.BlockchainInfo {
	return types.BlockchainInfo{
		Name:            TokenChainName,
		NetworkName:     NetworkNameTestnet,
		CoinUnit:        TokenUnit,
		ChainVersion:    Version,       // use our own blockChain/build version
		ProtocolVersion: build.Version, // use latest available rivine protocol version
	}
}

func GetDevnetGenesis() types.ChainConstants {
	cfg := types.DevnetChainConstants()

	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersion(1)
	cfg.GenesisTransactionVersion = types.TransactionVersion(1)

	// size limits
	cfg.BlockSizeLimit = 2000000
	cfg.ArbitraryDataSizeLimit = 83

	// block time
	cfg.BlockFrequency = 12

	// Time to MaturityDelay
	cfg.MaturityDelay = 10

	// The genesis timestamp
	cfg.GenesisTimestamp = types.Timestamp(1567691594)

	cfg.MedianTimestampWindow = 11

	// block window for difficulty
	cfg.TargetWindow = 20

	cfg.MaxAdjustmentUp = big.NewRat(120, 100)
	cfg.MaxAdjustmentDown = big.NewRat(100, 120)

	cfg.FutureThreshold = 120000000000
	cfg.ExtremeFutureThreshold = 180000000000

	cfg.StakeModifierDelay = 2000000000000

	// Time it takes before transferred blockstakes can be used
	cfg.BlockStakeAging = 1024000000000

	// Coins you receive when you create a block
	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(0) // Minimum transaction fee
	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(10)
	cfg.TransactionFeeCondition = types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f")))

	// Set Transaction Pool config
	cfg.TransactionPool = types.TransactionPoolConstants{
		TransactionSizeLimit:    16000,
		TransactionSetSizeLimit: 250000,
		PoolSizeLimit:           19750000,
	}

	// allocate initial coin outputs
	cfg.GenesisCoinDistribution = []types.CoinOutput{
		{
			Value:     cfg.CurrencyUnits.OneCoin.Mul64(200000000),
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f"))),
		},
	}

	// allocate initial block stake outputs
	cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
		{
			Value:     types.NewCurrency64(2000),
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f"))),
		},
	}

	return cfg
}

func GetDevnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
		"localhost:23112",
	}
}

func GetDevnetGenesisMintCondition() types.UnlockConditionProxy {
	return types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f")))
}

func GetTestnetGenesis() types.ChainConstants {
	cfg := types.TestnetChainConstants()

	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersion(1)
	cfg.GenesisTransactionVersion = types.TransactionVersion(1)

	// size limits
	cfg.BlockSizeLimit = 2000000
	cfg.ArbitraryDataSizeLimit = 83

	// block time
	cfg.BlockFrequency = 120

	// Time to MaturityDelay
	cfg.MaturityDelay = 720

	// The genesis timestamp
	cfg.GenesisTimestamp = types.Timestamp(1567691594)

	cfg.MedianTimestampWindow = 11

	// block window for difficulty
	cfg.TargetWindow = 1000

	cfg.MaxAdjustmentUp = big.NewRat(25, 10)
	cfg.MaxAdjustmentDown = big.NewRat(10, 25)

	cfg.FutureThreshold = 3000000000
	cfg.ExtremeFutureThreshold = 6000000000

	cfg.StakeModifierDelay = 20000000000

	// Time it takes before transferred blockstakes can be used
	cfg.BlockStakeAging = 1024000000000

	// Coins you receive when you create a block
	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(0) // Minimum transaction fee
	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(10)
	cfg.TransactionFeeCondition = types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("012baf6a2019b8184328f9ddc13e8aa6484c3272d98411444bdd743e9bb62e3572568ce6dc63bf")))

	// Set Transaction Pool config
	cfg.TransactionPool = types.TransactionPoolConstants{
		TransactionSizeLimit:    16000,
		TransactionSetSizeLimit: 250000,
		PoolSizeLimit:           19750000,
	}

	// allocate initial coin outputs
	cfg.GenesisCoinDistribution = []types.CoinOutput{
		{
			Value:     cfg.CurrencyUnits.OneCoin.Mul64(200000000),
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("012baf6a2019b8184328f9ddc13e8aa6484c3272d98411444bdd743e9bb62e3572568ce6dc63bf"))),
		},
	}

	// allocate initial block stake outputs
	cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
		{
			Value:     types.NewCurrency64(2000),
			Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("012baf6a2019b8184328f9ddc13e8aa6484c3272d98411444bdd743e9bb62e3572568ce6dc63bf"))),
		},
	}

	return cfg
}

func GetTestnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
		"bootstrap1.testnet.fft.threefold.tech:22312",
		"bootstrap2.testnet.fft.threefold.tech:22312",
		"bootstrap3.testnet.fft.threefold.tech:22312",
		"bootstrap4.testnet.fft.threefold.tech:22312",
	}
}

func GetTestnetGenesisMintCondition() types.UnlockConditionProxy {
	return types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("012baf6a2019b8184328f9ddc13e8aa6484c3272d98411444bdd743e9bb62e3572568ce6dc63bf")))
}

func init() {
	Version = build.MustParse(rawVersion)
}
