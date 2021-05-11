### EpiK Protocol DeFi SDK

#### Query Infomation Interface
* Miner Infomation
``` 
func GetMinerInfo(node api.FullNode, miner address.Address) (info MinerInfo, err error) 
```
* Voter Infomation
```
func GetVoterInfo(node api.FullNode, voter address.Address) (info VoterInfo, err error)
```
* Expert Infomation
```
func GetExpertInfo(node api.FullNode, expert address.Address) (info ExpertInfo, err error)
```

#### Mining Pledge Interface
* Mining Pledge Add
```
func (w *Wallet) MiningPledgeAdd(node api.FullNode, minerID address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error)
```
* Mining Pledge Withdraw
```
func (w *Wallet) MiningPledgeWithdraw(node api.FullNode, minerID address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) 
```

#### Retrieve Pledge Interface
* Retrieve Pledge Add
```
func (w *Wallet) RetrievePledgeAdd(node api.FullNode, target address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error)
```
* Retrieve Pledge Apply Withdraw
```
func (w *Wallet) RetrievePledgeApplyWithdraw(node api.FullNode, target address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error) 
```
* Retrieve Pledge Withdraw
```
func (w *Wallet) RetrievePledgeWithdraw(node api.FullNode, target address.Address, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error)
```

#### Vote Interface
* Vote Send
```
func (w *Wallet) VoteSend(node api.FullNode, candidate, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error)
```
* Vote Rescind
```
func (w *Wallet) VoteRescind(node api.FullNode, candidate, owner address.Address, amount decimal.Decimal) (cID cid.Cid, err error)
```
* Vote Withdraw
```
func (w *Wallet) VoteWithdraw(node api.FullNode, candidate, owner address.Address) (cID cid.Cid, err error)
```