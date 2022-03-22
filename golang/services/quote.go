package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jdram62/blockchain-auto-study/contracts/univ2/pair"
	"github.com/jdram62/blockchain-auto-study/contracts/univ2/router"
	"log"
)

const DENOM_6 float64 = 1000000.0
const DENOM_18 float64 = 1000000000000000000.0

var USDC common.Address = common.HexToAddress("0xea32a96608495e54156ae48931a7c20f0dcc1a21")

func Quote(ctx context.Context, watchList [][]string, clientEth *ethclient.Client, lpAddress common.Address) {
	poolInstance, err := pair.NewPair(lpAddress, clientEth) // create into functions
	if err != nil {
		log.Fatalln("Quote - NewPair", err)
	}
	uniRouter, err := router.NewRouter(lpAddress, clientEth) // create into function, test for quote function
	if err != nil {
		log.Fatalln("Quote - NewRouter", err)
	}
	fmt.Println(uniRouter)
	reserves, err := poolInstance.GetReserves(nil)
	if err != nil {
		log.Fatalln("Quote - GetReserves", err)
	}
	fmt.Println(reserves)
	//quote, err := uniRouter.Quote(nil, amountIn, reserves.Reserve0, reserves.Reserve1)
	if err != nil {
		log.Fatalln("Quote - Quote", err)
	}

	// name, err := poolInstance.Name(nil)
	if err != nil {
		log.Fatalln("Quote - Name", err)
	}
	//exchange := name[0:1]
	// check for mainctx cancel before doing anything major
	if ctx.Err() != nil {
		log.Println("context done")
		return
	}
	//amount_to_trade := 500.0
	// convert
	/*if token1 == USDC {
		reserve0 = new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve0), big.NewFloat(DENOM_18))
		reserve1 = new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve1), big.NewFloat(DENOM_6))
	} else {
		reserve0 = new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve0), big.NewFloat(DENOM_18))
		reserve1 = new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve0), big.NewFloat(DENOM_18))
	}*/

}
