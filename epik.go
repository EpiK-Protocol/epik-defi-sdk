package epik

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"epik/wallet"
	"fmt"
	"net/http"

	"epik/client"

	"github.com/EpiK-Protocol/go-epik/api"
	"github.com/EpiK-Protocol/go-epik/chain/actors"
	"github.com/EpiK-Protocol/go-epik/chain/actors/builtin/miner"
	"github.com/EpiK-Protocol/go-epik/chain/actors/builtin/retrieval"
	"github.com/EpiK-Protocol/go-epik/chain/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin"
	fminer "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/vote"
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/shopspring/decimal"
)

func NewFullNodeAPI(endPoint, token string) (node api.FullNode, closer jsonrpc.ClientCloser, err error) {
	fmt.Println("new fullnode api")
	httpHeader := http.Header{}
	httpHeader.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	node, closer, err = client.NewFullNodeRPC(context.Background(), endPoint, httpHeader)
	return
}

func NewEpikWallet() (epikWallet *Wallet, err error) {
	ks := wallet.NewMemKeyStore()
	w, err := wallet.NewWallet(ks)
	epikWallet = &Wallet{w}
	return
}

func (w *Wallet) ImportPrivateKey(privateKey string) (addr address.Address, err error) {
	keyInfo := &types.KeyInfo{}
	data, err := hex.DecodeString(privateKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, keyInfo)
	if err != nil {
		return
	}
	return w.WalletImport(context.Background(), keyInfo)
}

func (w *Wallet) sendMessage(node api.FullNode, msg *types.Message) (cID cid.Cid, err error) {
	msg.Nonce, err = node.MpoolGetNonce(context.Background(), msg.From)
	if err != nil {
		return
	}
	msg.GasFeeCap, err = node.GasEstimateFeeCap(context.Background(), msg, 0, types.EmptyTSK)
	if err != nil {
		return
	}
	msg.GasLimit, err = node.GasEstimateGasLimit(context.Background(), msg, types.EmptyTSK)
	if err != nil {
		return
	}
	signature, err := w.WalletSign(context.Background(), msg.From, msg.Cid().Bytes(), api.MsgMeta{})
	if err != nil {
		return cid.Undef, err
	}
	signedMsg := &types.SignedMessage{
		Message:   *msg,
		Signature: *signature,
	}
	fmt.Println(json.Marshal(signedMsg))
	return node.MpoolPush(context.Background(), signedMsg)
}

func (w *Wallet) MiningPledgeAdd(node api.FullNode, minerID address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {

	bal, err := node.WalletBalance(context.Background(), owner)
	if err != nil {
		return
	}
	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}
	if bal.LessThan(big.Int(epk)) {
		return cid.Undef, fmt.Errorf("not enough balance")
	}
	msg := &types.Message{
		To:     minerID,
		From:   owner,
		Value:  abi.TokenAmount(epk),
		Method: miner.Methods.AddPledge,
		Params: nil,
	}
	return w.sendMessage(node, msg)
}

func (w *Wallet) MiningPledgeWithdraw(node api.FullNode, minerID address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {
	ownerID, err := node.StateLookupID(context.Background(), owner, types.EmptyTSK)
	if err != nil {
		return
	}
	funds, err := node.StateMinerFunds(context.Background(), minerID, types.EmptyTSK)
	if err != nil {
		return
	}
	bal := funds.MiningPledgors[ownerID.String()]
	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}
	if bal.LessThan(big.Int(epk)) {
		return cid.Undef, fmt.Errorf("not enough balance")
	}
	params, err := actors.SerializeParams(&fminer.WithdrawPledgeParams{
		AmountRequested: abi.TokenAmount(epk), // Default to attempting to withdraw all the extra funds in the miner actor
	})
	if err != nil {
		return
	}
	msg := &types.Message{
		To:     minerID,
		From:   owner,
		Value:  types.NewInt(0),
		Method: miner.Methods.WithdrawPledge,
		Params: params,
	}
	return w.sendMessage(node, msg)
}

func (w *Wallet) RetrievePledgeAdd(node api.FullNode, target address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {

	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}
	params, err := actors.SerializeParams(&target)
	if err != nil {
		return cid.Undef, fmt.Errorf("serializing params failed: %w", err)
	}

	bal, err := node.WalletBalance(context.Background(), owner)
	if err != nil {
		return
	}
	if bal.LessThan(big.Int(epk)) {
		return cid.Undef, fmt.Errorf("not enough balance")
	}
	msg := &types.Message{
		To:     retrieval.Address,
		From:   owner,
		Value:  abi.TokenAmount(epk),
		Method: retrieval.Methods.AddBalance,
		Params: params,
	}
	return w.sendMessage(node, msg)
}

func (w *Wallet) RetrievePledgeApplyWithdraw(node api.FullNode, target address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {
	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}
	params, err := actors.SerializeParams(&retrieval.WithdrawBalanceParams{
		ProviderOrClientAddress: owner,
		Amount:                  big.Int(epk),
	})
	if err != nil {
		return cid.Undef, xerrors.Errorf("serializing params failed: %w", err)
	}
	retrieve, err := node.StateRetrievalPledge(context.Background(), target, types.EmptyTSK)
	if err != nil {
		return
	}
	if retrieve.Balance.LessThan(big.Int(epk)) {
		return cid.Undef, fmt.Errorf("not enough balance")
	}
	msg := &types.Message{
		To:     retrieval.Address,
		From:   owner,
		Value:  abi.NewTokenAmount(0),
		Method: retrieval.Methods.ApplyForWithdraw,
		Params: params,
	}
	return w.sendMessage(node, msg)
}
func (w *Wallet) RetrievePledgeWithdraw(node api.FullNode, target address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {
	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}
	params, err := actors.SerializeParams(&retrieval.WithdrawBalanceParams{
		ProviderOrClientAddress: owner,
		Amount:                  big.Int(epk),
	})
	if err != nil {
		return cid.Undef, xerrors.Errorf("serializing params failed: %w", err)
	}

	retrieve, err := node.StateRetrievalPledge(context.Background(), target, types.EmptyTSK)
	if err != nil {
		return
	}
	if retrieve.Balance.LessThan(big.Int(epk)) {
		return cid.Undef, fmt.Errorf("not enough balance")
	}
	msg := &types.Message{
		To:     retrieval.Address,
		From:   owner,
		Value:  abi.NewTokenAmount(0),
		Method: retrieval.Methods.WithdrawBalance,
		Params: params,
	}
	return w.sendMessage(node, msg)
}

func (w *Wallet) VoteSend(node api.FullNode, candidate, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {
	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}
	sp, err := actors.SerializeParams(&candidate)
	if err != nil {
		return
	}

	msg := &types.Message{
		From:   owner,
		To:     builtin.VoteFundActorAddr,
		Value:  types.BigInt(epk),
		Method: builtin.MethodsVote.Vote,
		Params: sp,
	}
	return w.sendMessage(node, msg)
}
func (w *Wallet) VoteRescind(node api.FullNode, candidate, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) {
	epk, err := types.ParseEPK(fmt.Sprintf("%sEPK", amount.String()))
	if err != nil {
		return
	}

	sp, err := actors.SerializeParams(&vote.RescindParams{
		Candidate: candidate,
		Votes:     types.BigInt(epk),
	})
	if err != nil {
		return
	}

	msg := &types.Message{
		From:   owner,
		To:     builtin.VoteFundActorAddr,
		Value:  types.BigInt(epk),
		Method: builtin.MethodsVote.Rescind,
		Params: sp,
	}
	return w.sendMessage(node, msg)
}
func (w *Wallet) VoteWithdraw(node api.FullNode, candidate, owner address.Address) (cID cid.Cid, err error) {

	sp, err := actors.SerializeParams(&candidate)
	if err != nil {
		return
	}

	msg := &types.Message{
		From:   owner,
		To:     builtin.VoteFundActorAddr,
		Value:  big.Zero(),
		Method: builtin.MethodsVote.Withdraw,
		Params: sp,
	}
	return w.sendMessage(node, msg)
}

func GetVoterInfo(node api.FullNode, voter address.Address) (info VoterInfo, err error) {
	i, err := node.StateVoterInfo(context.Background(), voter, types.EmptyTSK)
	if err != nil {
		return
	}
	info = VoterInfo{
		VoterInfo: *i,
	}
	return
}

func GetExpertInfo(node api.FullNode, expert address.Address) (info ExpertInfo, err error) {
	i, err := node.StateExpertInfo(context.Background(), expert, types.EmptyTSK)
	if err != nil {
		return
	}
	info = ExpertInfo{
		ExpertInfo: *i,
	}
	return
}

func GetMinerInfo(node api.FullNode, miner address.Address) (info MinerInfo, err error) {
	ctx := context.Background()
	i, err := node.StateMinerInfo(ctx, miner, types.EmptyTSK)
	if err != nil {
		return
	}
	power, err := node.StateMinerPower(ctx, miner, types.EmptyTSK)
	if err != nil {
		return
	}
	funds, err := node.StateMinerFunds(ctx, miner, types.EmptyTSK)
	if err != nil {
		return
	}
	retrieve, err := node.StateRetrievalPledge(ctx, info.Owner, types.EmptyTSK)
	if err != nil {
		return
	}
	balance, err := node.StateMinerAvailableBalance(ctx, info.Owner, types.EmptyTSK)
	if err != nil {
		return
	}
	info = MinerInfo{
		MinerInfo:        i,
		MinerPower:       *power,
		Funds:            funds,
		AvailableBalance: balance,
		RetrievalState:   *retrieve,
	}
	return
}
