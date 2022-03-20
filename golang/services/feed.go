package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	uni "github.com/jdram62/blockchain-auto-study/contracts/uniV2"
	"log"
)

func Feed(ctxMain context.Context, endPoint *string) {
	// Pulls watchlist
	watchList := GetWatchList()
	if *endPoint == "test" {
		*endPoint = GetEndpoints()[0][1]
	} else {
		*endPoint = GetEndpoints()[0][0]
	}
	// Connect client
	clientRpc, err := rpc.DialContext(ctxMain, *endPoint)
	if err != nil {
		log.Fatalln("Feed - DialContext:", err)
	}
	defer clientRpc.Close()
	clientEth := ethclient.NewClient(clientRpc)
	defer clientEth.Close()
	// get addresses from watchlist for subscription filter
	var addressesWatch []common.Address
	for _, row := range watchList {
		addressesWatch = append(addressesWatch, common.HexToAddress(row[1]))
	}
	query := ethereum.FilterQuery{
		Addresses: addressesWatch,
	}
	if err != nil {
		log.Fatalln("Feed - NewUniswapV2ERC20:", err)
	}
	logs := make(chan types.Log)
	// possibly creat a new ctx
	sub, err := clientEth.SubscribeFilterLogs(ctxMain, query, logs)
	if err != nil {
		log.Fatalln("Feed - SubscribeNewHead:", err)
	}
	defer sub.Unsubscribe()
	var poolInstance *uni.UniswapV2ERC20
	for {
		select {
		default:
		case <-ctxMain.Done():
			return
		case err = <-sub.Err():
			fmt.Println(err, "new sub ")
		case swap := <-logs:
			// Check lp reserves of the swap
			poolInstance, err = uni.NewUniswapV2ERC20(swap.Address, clientEth)
			if err != nil {
				log.Fatalln("Feed - NewUniswapV2ERC20")
			}
			reserves, _ := poolInstance.GetReserves(nil)
			if err != nil {
				log.Fatalln("Feed - GetReserves")
			}
			fmt.Println("addy: ", swap.Address, "lp reserves: ", reserves)
		}

	}
}
