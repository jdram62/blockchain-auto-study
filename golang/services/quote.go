package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	uni "github.com/jdram62/blockchain-auto-study/contracts/uniV2"
	"log"
	"math/big"
)

func Quote(ctx context.Context, clientEth *ethclient.Client, swapAddress *common.Address, txHash *common.Hash) {
	var poolInstance *uni.UniswapV2ERC20
	poolInstance, err := uni.NewUniswapV2ERC20(*swapAddress, clientEth)
	if err != nil {
		log.Fatalln("Quote - NewUniswapV2ERC20", err)
	}
	reserves, err := poolInstance.GetReserves(nil)
	if err != nil {
		log.Fatalln("Quote - GetReserves", err)
	}
	decimals, err := poolInstance.Decimals(nil)
	if err != nil {
		log.Fatalln("Quote - Decimals", err)
	}
	div := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	fmt.Println("div", div)
	reserve0 := new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve0), new(big.Float).SetInt(div))
	//reserve1 := new(big.Float).Quo(new(big.Float).SetInt(reserves.Reserve1), new(big.Float).SetInt(div))
	//mul1 := new(big.Float).Mul(new(big.Float).SetInt(reserves.Reserve1), new(big.Float).SetInt(div))

	//usdc has 6 decimals
	fmt.Println(*swapAddress, reserve0, reserves.Reserve1)

}
