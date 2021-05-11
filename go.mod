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
	github.com/sirupsen/logrus v1.7.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

// replace github.com/EpiK-Protocol/go-epik => ../go-epik

// replace github.com/supranational/blst => github.com/supranational/blst v0.2.0

replace github.com/filecoin-project/specs-storage => github.com/EpiK-Protocol/go-epik-storage v0.1.1-0.20210109141728-73c1715728b4

replace github.com/filecoin-project/specs-actors/v2 => github.com/EpiK-Protocol/go-epik-actors/v2 v2.0.0-20210321153151-558c615d25f2

replace github.com/filecoin-project/go-fil-markets => github.com/EpiK-Protocol/go-epik-markets v0.5.3-0.20210418131717-54b84f7175db

replace github.com/filecoin-project/go-data-transfer => github.com/EpiK-Protocol/go-data-transfer v1.1.1-0.20210218064751-b53c28fecbf2

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
