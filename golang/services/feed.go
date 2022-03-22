package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"runtime"
	"time"
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
	clientEth, err := ethclient.DialContext(ctxMain, *endPoint)
	if err != nil {
		log.Fatalln("Feed - DialContext:", err)
	}
	defer clientEth.Close()
	// get addresses from watchlist for subscription filter
	//var addressesWatch []common.Address
	//for _, row := range watchList {
	//addressesWatch = append(addressesWatch, common.HexToAddress(row[0]))
	//}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xDd7dF3522a49e6e1127bf1A1d3bAEa3bc100583B")},
	}
	if err != nil {
		log.Fatalln("Feed - NewUniswapV2ERC20:", err)
	}
	// transaction stream
	logs := make(chan types.Log)
	var recentBlock uint64 = 1
	ctx, cancel := context.WithCancel(ctxMain)
	sub, err := clientEth.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalln("Feed - SubscribeNewHead:", err)
	}
	defer sub.Unsubscribe()
	for {
		ctxQ, stop := context.WithTimeout(ctxMain, time.Second*3)
		select {
		default:
		case swap := <-logs:
			if swap.BlockNumber > recentBlock {
				recentBlock = swap.BlockNumber
				go Quote(ctxQ, watchList, clientEth, swap.Address)
				fmt.Println("gos: ", runtime.NumGoroutine(), " index: ", swap.BlockNumber, " tx: ", swap.TxHash)
			}
		case <-ctxMain.Done():
			stop()
			cancel()
			return
		case err = <-sub.Err():
			fmt.Println("Feed - sub.Err", err)
		}
	}
}
