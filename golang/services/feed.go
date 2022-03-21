package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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
	ctx, cancel := context.WithTimeout(ctxMain, 10*time.Second)
	sub, err := clientEth.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalln("Feed - SubscribeNewHead:", err)
	}
	defer sub.Unsubscribe()
	var recentTxs common.Hash
	for {
		select {
		default:
		case swap := <-logs:
			if recentTxs != swap.TxHash {
				recentTxs = swap.TxHash
				go Quote(ctxMain, clientEth, swap.Address, swap.TxHash)
			}
		case <-ctxMain.Done():
			cancel()
			return
		case err = <-sub.Err():
			fmt.Println("Feed - sub.Err", err)
		}
	}
}
