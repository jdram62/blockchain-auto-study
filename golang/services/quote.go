package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	uni "github.com/jdram62/blockchain-auto-study/contracts/uniV2"
	"log"
)

func Quote(ctx context.Context, clientEth *ethclient.Client, swapAddress common.Address, txHash common.Hash) {
	var poolInstance *uni.UniswapV2ERC20
	poolInstance, err := uni.NewUniswapV2ERC20(swapAddress, clientEth)
	if err != nil {
		log.Fatalln("Quote - NewUniswapV2ERC20", err)
	}
	addy0, err := poolInstance.Token0(nil)
	if err != nil {
		log.Fatalln("Quote - Token0", err)
	}
	addy1, err := poolInstance.Token1(nil)
	if err != nil {
		log.Fatalln("Quote - Token1Token1", err)
	}
	name, err := poolInstance.Name(nil)
	if err != nil {
		log.Fatalln("Quote - Name", err)
	}
	fmt.Println(name, " ", addy0, " ", addy1)

	/*reserves, err := poolInstance.GetReserves(nil)
	if err != nil {
		log.Fatalln("Quote - GetReserves", err)
	}
	div := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	reserve0 := new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve0), new(big.Float).SetInt(div))
	//usdc has 10^6 - hard code denoms
	reserve1 := new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve1), big.NewFloat(1000000.0))*/
}
