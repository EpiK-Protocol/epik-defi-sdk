package epik

import (
	"github.com/EpiK-Protocol/epik-defi-sdk/wallet"

	"github.com/EpiK-Protocol/go-epik/api"
	"github.com/EpiK-Protocol/go-epik/chain/actors/builtin/miner"
	"github.com/EpiK-Protocol/go-epik/chain/actors/builtin/vesting"
	"github.com/EpiK-Protocol/go-epik/chain/actors/builtin/vote"
	"github.com/filecoin-project/go-state-types/big"
)

type Wallet struct {
	*wallet.LocalWallet
}

type MinerInfo struct {
	miner.MinerInfo
	api.MinerPower
	miner.Funds
	api.RetrievalState
	AvailableBalance big.Int
}

type VoterInfo struct {
	vote.VoterInfo
}

type ExpertInfo struct {
	api.ExpertInfo
}

type CoinbaseInfo struct {
	vesting.CoinbaseInfo
}
