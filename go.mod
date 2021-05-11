module epik

go 1.16

require (
	github.com/EpiK-Protocol/go-epik v0.4.2-0.20210424170134-d3cdb5101755
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/go-state-types v0.1.0
	github.com/filecoin-project/specs-actors/v2 v2.3.4
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-log/v2 v2.1.2-0.20200626104915-0016c0b4b3e4
	github.com/shopspring/decimal v1.2.0
	github.com/supranational/blst v0.2.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace github.com/filecoin-project/specs-actors/v2 => github.com/EpiK-Protocol/go-epik-actors/v2 v2.0.0-20210321153151-558c615d25f2
