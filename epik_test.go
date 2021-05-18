package epik

import (
	"context"
	"testing"

	"github.com/EpiK-Protocol/go-epik/chain/types"
	"github.com/filecoin-project/go-address"
	"github.com/shopspring/decimal"
)

const (
	endPoint       string = "ws://122.9.153.165:30536/rpc/v0"
	token          string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.wDA2ZJlAVH7EJFubQXAEw-qI4-TNsj8DrMFwva0qkhw"
	testPrivateKey string = "7b2254797065223a22626c73222c22507269766174654b6579223a22486e322f42703432456167647052323538416348793665484b2b494d4e6c42782b2f466f676261323941303d227d"
)

func TestEpikBalance(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	addr, err := address.NewFromString("f010001")
	if err != nil {
		panic(err)
	}
	bal, err := node.WalletBalance(context.Background(), addr)
	if err != nil {
		panic(err)
	}
	t.Logf("balance is %s", bal.String())
}

func TestAccountKey(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	addr, err := address.NewFromString("f031817")
	if err != nil {
		panic(err)
	}
	addr, err = node.StateAccountKey(context.Background(), addr, types.EmptyTSK)
	if err != nil {
		panic(err)
	}
	t.Logf("account is %s", addr.String())
}

func TestMiningPledgeAdd(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	miner, err := address.NewFromString("f031829")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	c, err := w.MiningPledgeAdd(node, miner, owner, decimal.NewFromInt(1000))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestMiningPledgeWithdraw(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	miner, err := address.NewFromString("f031829")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	c, err := w.MiningPledgeWithdraw(node, miner, owner, decimal.NewFromInt(1000))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestRetrievePledgeAdd(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	target, err := address.NewFromString("f031817")
	if err != nil {
		panic(err)
	}
	miners := []address.Address{}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	c, err := w.RetrievePledgeAdd(node, target, miners, owner, decimal.NewFromInt(1000))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestRetrievePledgeApplyWithdraw(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	daemon, err := address.NewFromString("f031817")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	pledge, err := node.StateRetrievalPledge(context.Background(), daemon, types.EmptyTSK)
	if err != nil {
		panic(err)
	}
	amount := pledge.Balance
	c, err := w.RetrievePledgeApplyWithdraw(node, daemon, owner, decimal.NewFromBigInt(amount.Int, -18))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestRetrievePledgeWithdraw(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	daemon, err := address.NewFromString("f031817")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	pledge, err := node.StateRetrievalPledge(context.Background(), daemon, types.EmptyTSK)
	if err != nil {
		panic(err)
	}
	amount := pledge.Balance
	c, err := w.RetrievePledgeWithdraw(node, daemon, owner, decimal.NewFromBigInt(amount.Int, -18))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestVoteSend(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	expert, err := address.NewFromString("f01000")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	c, err := w.VoteSend(node, expert, owner, decimal.NewFromInt(1000))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestVoteRescind(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	expert, err := address.NewFromString("f01000")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	info, err := GetVoterInfo(node, owner)
	if err != nil {
		panic(err)
	}
	amount := info.Candidates[expert.String()]
	c, err := w.VoteRescind(node, expert, owner, decimal.NewFromBigInt(amount.Int, -18))
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}

func TestVoteWithdraw(t *testing.T) {
	node, closer, err := NewFullNodeAPI(endPoint, token)
	if err != nil {
		panic(err)
	}
	defer closer()
	expert, err := address.NewFromString("f01000")
	if err != nil {
		panic(err)
	}
	w, err := NewEpikWallet()
	if err != nil {
		panic(err)
	}
	owner, err := w.ImportPrivateKey(testPrivateKey)
	if err != nil {
		panic(err)
	}
	c, err := w.VoteWithdraw(node, expert, owner)
	if err != nil {
		panic(err)
	}
	t.Logf("message cid is %s", c.String())
}
