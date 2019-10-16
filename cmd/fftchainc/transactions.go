package main

import (
	fftchaintypes "github.com/threefoldtech/FFTChain/pkg/types"
	"github.com/threefoldtech/rivine/extensions/minting"
	mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"
	"github.com/threefoldtech/rivine/types"

	"github.com/threefoldtech/rivine/pkg/client"
)

func RegisterDevnetTransactions(bc client.BaseClient) {
	registerTransactions(bc)
}

func RegisterTestnetTransactions(bc client.BaseClient) {
	registerTransactions(bc)
}

func registerTransactions(bc client.BaseClient) {
	// create minting plugin client...
	mintingCLI := mintingcli.NewPluginConsensusClient(bc)
	// ...and register minting types
	types.RegisterTransactionVersion(fftchaintypes.TransactionVersionMinterDefinition, minting.MinterDefinitionTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  fftchaintypes.TransactionVersionMinterDefinition,
	})
	types.RegisterTransactionVersion(fftchaintypes.TransactionVersionCoinCreation, minting.CoinCreationTransactionController{
		MintConditionGetter: mintingCLI,
		TransactionVersion:  fftchaintypes.TransactionVersionCoinCreation,
	})
	types.RegisterTransactionVersion(fftchaintypes.TransactionVersionCoinDestruction, minting.CoinDestructionTransactionController{
		TransactionVersion: fftchaintypes.TransactionVersionCoinDestruction,
	})

}
