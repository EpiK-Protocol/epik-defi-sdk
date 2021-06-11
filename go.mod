module github.com/EpiK-Protocol/epik-defi-sdk

go 1.16

require (
	github.com/EpiK-Protocol/go-epik v0.4.2-0.20210517071105-121cd49dd879
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

replace github.com/filecoin-project/specs-actors/v2 => github.com/EpiK-Protocol/go-epik-actors/v2 v2.4.0-alpha.0.20210517033919-7cac385c0096
