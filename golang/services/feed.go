package services

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"os"
	"time"
)

func getWatchList() [][]string {
	file, err := os.Open("watchlist.csv")
	if err != nil {
		log.Fatalln("getWatchList - Open:", err)
	}
	csvReader := csv.NewReader(file)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln("getWatchList - ReadAll:", err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln("getWatchList - Close:", err)
	}
	return csvData
}

func Feed(ctxMain context.Context, endPoint *string) {
	// Pulls watchlist
	watchList := getWatchList()
	fmt.Println(watchList[0])
	// Connect client
	rpcClient, err := rpc.DialContext(ctxMain, *endPoint)
	if err != nil {
		log.Fatalln("Feed - DialContext:", err)
	}
	defer rpcClient.Close()
	ethClient := ethclient.NewClient(rpcClient)
	defer ethClient.Close()
	// Check reserves of lp contract
	address := common.HexToAddress("0x5ab390084812E145b619ECAA8671d39174a1a6d1")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	//instance, err := uni.NewUniswapV2ERC20(address, ethClient)
	if err != nil {
		log.Fatalln("Feed - NewUniswapV2ERC20:", err)
	}
	logs := make(chan types.Log)
	//blockHeaders := make(chan *types.Header, 16)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//sub, err := ethClient.SubscribeNewHead(ctx, blockHeaders)
	sub, err := ethClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalln("Feed - SubscribeNewHead:", err)
	}
	defer sub.Unsubscribe()
	for {
		select {
		case <-ctxMain.Done():
			fmt.Println("Grace exit")
			return
		case err = <-sub.Err():
			fmt.Println(err, "new block sub ")
		case swap := <-logs:
			fmt.Println(swap.TxHash, swap.Data)
		}
	}
}
