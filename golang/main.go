// Author: Jakob Ramirez
package main

import (
	"context"
	"flag"
	"fmt"
	services "github.com/jdram62/blockchain-auto-study/services"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*WIP:
price quote functions for different exchanges
	consider orders smart routing
*/

func main() {
	endPoint := flag.String("ep", "main", "Websocket RPC Endpoint")
	var userInput int
	for loop := true; loop; {
		fmt.Println("\n\tWATCHLIST\n1 - Add Contract Address\n2 - Remove Contract Address\n3 - Continue")
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			log.Println("main - userInput", err)
		}
		switch userInput {
		case 1:
			_, err = services.AddContractWatchlist()
			if err != nil {
				log.Fatalln("main - AddContractWatchlist", err)
			}
		case 2:
			_, err = services.RemoveContractWatchlist()
			if err != nil {
				log.Fatalln("main - RemoveContractWatchlist", err)
			}
		case 3:
			loop = false
		}
	}
	fmt.Println("...")
	ctxMain, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	services.Feed(ctxMain, endPoint)
	fmt.Println("...")
	time.Sleep(2 * time.Second)
}
